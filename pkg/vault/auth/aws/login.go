/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package aws

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"kubevault.dev/apimachinery/apis"
	vsapi "kubevault.dev/apimachinery/apis/kubevault/v1alpha1"
	authtype "kubevault.dev/operator/pkg/vault/auth/types"
	vaultuitl "kubevault.dev/operator/pkg/vault/util"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	vaultapi "github.com/hashicorp/vault/api"
	"github.com/hashicorp/vault/sdk/helper/awsutil"
	"github.com/pkg/errors"
)

const (
	iamServerIdHeader = "X-Vault-AWS-IAM-Server-ID"
)

type auth struct {
	vClient     *vaultapi.Client
	creds       *credentials.Credentials
	headerValue string
	role        string
	path        string
}

// links : https://www.vaultproject.io/docs/auth/aws.html

func New(authInfo *authtype.AuthInfo) (*auth, error) {
	if authInfo == nil {
		return nil, errors.New("authentication information is empty")
	}
	if authInfo.VaultApp == nil {
		return nil, errors.New("AppBinding is empty")
	}

	vApp := authInfo.VaultApp
	cfg, err := vaultuitl.VaultConfigFromAppBinding(vApp)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create vault config from AppBinding")
	}

	vc, err := vaultapi.NewClient(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create vault client")
	}

	if authInfo.Secret == nil {
		return nil, errors.New("authentication secret is missing")
	}

	secret := authInfo.Secret
	accessKeyID, ok := secret.Data[apis.AWSAuthAccessKeyIDKey]
	if !ok {
		return nil, errors.New("access_key_id is missing")
	}
	secretAccessKey, ok := secret.Data[apis.AWSAuthAccessSecretKey]
	if !ok {
		return nil, errors.New("secret_access_key is missing")
	}
	securityToken := secret.Data[apis.AWSAuthSecurityTokenKey]

	authPath := string(vsapi.AuthTypeAws)
	if authInfo.Path != "" {
		authPath = authInfo.Path
	}

	var headerValue string
	if authInfo.ExtraInfo != nil && authInfo.ExtraInfo.AWS != nil {
		headerValue = authInfo.ExtraInfo.AWS.HeaderValue
	}

	creds, err := retrieveCreds(string(accessKeyID), string(secretAccessKey), string(securityToken))
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve credentials")
	}

	if authInfo.VaultRole == "" {
		return nil, errors.Wrap(err, "Vault role is empty")
	}

	return &auth{
		vClient:     vc,
		creds:       creds,
		role:        authInfo.VaultRole,
		headerValue: headerValue,
		path:        authPath,
	}, nil
}

// Login will log into vault and return client token
func (a *auth) Login() (string, error) {
	path := fmt.Sprintf("/v1/auth/%s/login", a.path)
	req := a.vClient.NewRequest("POST", path)

	loginData, err := a.generateLoginData()
	if err != nil {
		return "", errors.Wrap(err, "failed to generate login data")
	}

	payload := loginData
	payload["role"] = a.role
	if err := req.SetJSONBody(payload); err != nil {
		return "", err
	}

	resp, err := a.vClient.RawRequest(req)
	if err != nil {
		return "", err
	}

	var loginResp authtype.AuthLoginResponse
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&loginResp)
	if err != nil {
		return "", err
	}
	return loginResp.Auth.ClientToken, nil
}

// Generates the necessary data to send to the Vault server for generating a token
// This is useful for other API clients to use
func (a *auth) generateLoginData() (map[string]interface{}, error) {
	loginData := make(map[string]interface{})

	// Use the credentials we've found to construct an STS session
	stsSession, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{Credentials: a.creds},
	})
	if err != nil {
		return nil, err
	}

	var params *sts.GetCallerIdentityInput
	svc := sts.New(stsSession)
	stsRequest, _ := svc.GetCallerIdentityRequest(params)

	// Inject the required auth header value, if supplied, and then sign the request including that header
	if a.headerValue != "" {
		stsRequest.HTTPRequest.Header.Add(iamServerIdHeader, a.headerValue)
	}
	err = stsRequest.Sign()
	if err != nil {
		return nil, err
	}

	// Now extract out the relevant parts of the request
	headersJson, err := json.Marshal(stsRequest.HTTPRequest.Header)
	if err != nil {
		return nil, err
	}
	requestBody, err := ioutil.ReadAll(stsRequest.HTTPRequest.Body)
	if err != nil {
		return nil, err
	}
	loginData["iam_http_request_method"] = stsRequest.HTTPRequest.Method
	loginData["iam_request_url"] = base64.StdEncoding.EncodeToString([]byte(stsRequest.HTTPRequest.URL.String()))
	loginData["iam_request_headers"] = base64.StdEncoding.EncodeToString(headersJson)
	loginData["iam_request_body"] = base64.StdEncoding.EncodeToString(requestBody)

	return loginData, nil
}

func retrieveCreds(accessKey, secretKey, sessionToken string) (*credentials.Credentials, error) {
	credConfig := &awsutil.CredentialsConfig{
		AccessKey:    accessKey,
		SecretKey:    secretKey,
		SessionToken: sessionToken,
	}
	creds, err := credConfig.GenerateCredentialChain()
	if err != nil {
		return nil, err
	}
	if creds == nil {
		return nil, errors.New("could not compile valid credential providers from static config, environment, shared, or instance metadata")
	}

	_, err = creds.Get()
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve credentials from credential chain")
	}
	return creds, nil
}
