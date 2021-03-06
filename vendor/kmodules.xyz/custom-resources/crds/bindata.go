// Package crds Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// appcatalog.appscode.com_appbindings.v1.yaml
// appcatalog.appscode.com_appbindings.yaml
package crds

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _appcatalogAppscodeCom_appbindingsV1Yaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x5a\xdd\x6f\xdb\x46\x12\x7f\xf7\x5f\x31\x50\x1f\x9c\x00\xfa\x40\xaf\x2f\x07\xdd\x43\xcf\x71\x5c\x20\x57\xc7\x09\x2c\x37\x87\xe2\x72\x38\x2f\xb9\x23\x71\x6b\x72\x97\xdd\x5d\x4a\xd6\x15\xf9\xdf\x0f\x33\xbb\x14\x29\x89\xa4\xe4\x06\x01\x8e\x0f\x6d\xc4\xdd\x9d\x9d\xcf\xdf\x7c\xd0\xa2\x54\x9f\xd0\x3a\x65\xf4\x1c\x44\xa9\xf0\xd9\xa3\xa6\x5f\x6e\xfa\xf4\x57\x37\x55\x66\xb6\xfe\xfe\xe2\x49\x69\x39\x87\xeb\xca\x79\x53\xdc\xa3\x33\x95\x4d\xf1\x2d\x2e\x95\x56\x5e\x19\x7d\x51\xa0\x17\x52\x78\x31\xbf\x00\x48\x2d\x0a\x7a\xf9\xa0\x0a\x74\x5e\x14\xe5\x1c\x74\x95\xe7\x17\x00\xb9\x48\x30\x77\xb4\x07\x40\x94\xe5\xf4\xa9\x4a\xd0\x6a\xf4\xc8\xb7\x68\x51\xe0\x1c\x52\xe1\x45\x6e\x56\x17\x00\xe1\xb7\x28\xcb\x44\x69\xa9\xf4\xca\x4d\x45\x59\xc6\x65\xfa\xa7\x4b\x8d\xc4\x69\x6a\x8a\x0b\x57\x62\x4a\x54\x57\xd6\x54\x25\x1f\xe9\xdc\x16\x48\xc6\xfb\x53\xe1\x71\x65\xac\xaa\x7f\x4f\x5a\x37\xd3\xaf\xfa\x64\xfd\x93\x05\x00\x08\x7a\xb8\x2a\xcb\x37\x81\x29\x7e\x99\x2b\xe7\x7f\x3e\x58\xb8\x55\xce\xf3\x62\x99\x57\x56\xe4\x7b\x82\xf0\x7b\xa7\xf4\xaa\xca\x85\x6d\xaf\x5c\x00\xb8\xd4\x94\x38\x87\x3b\xe2\xb4\x14\x29\xca\x0b\x80\x75\xb0\x0e\x73\x3a\x01\x21\x25\x2b\x5d\xe4\x1f\xad\xd2\x1e\xed\xb5\xc9\xab\x42\xef\xe4\xf8\xcd\x19\xfd\x51\xf8\x6c\x0e\x53\x52\xcc\xd4\x6f\xcb\x20\x45\xad\xd2\x87\xe6\x05\xad\xcd\xc1\x79\x5b\x8b\x72\x7c\x3c\x5e\xbe\x47\xe1\xd3\xde\xbb\x61\x22\xb5\x6b\x4c\x8f\xfc\x62\x8f\xe4\xd5\x6a\x9f\x27\x29\x7c\x78\x11\x96\xd7\xdf\x8b\xbc\xcc\xc4\xf7\x41\x75\x69\x86\x85\x98\xc7\xfd\xa6\x44\x7d\xf5\xf1\xdd\xa7\x1f\x16\x7b\xaf\x01\x4a\x6b\x4a\xb4\x7e\x67\xe2\xf0\xb4\xbc\xbd\xf5\x16\x40\xa2\x4b\xad\x2a\x3d\x87\xc1\x25\x11\x0c\xbb\x40\x92\x9b\xa3\x03\x9f\x61\x6d\x09\x94\x91\x07\x30\x4b\xf0\x99\x72\x60\xb1\xb4\xe8\x50\x7b\x16\x71\x8f\x30\xd0\x26\xa1\xc1\x24\xbf\x61\xea\xa7\xb0\x40\x4b\x64\xc0\x65\xa6\xca\x25\xa4\x46\xaf\xd1\x7a\xb0\x98\x9a\x95\x56\xff\xdd\xd1\x76\xe0\x0d\x5f\x9a\x0b\x8f\xd1\x99\x9a\x87\x2d\xaf\x45\x0e\x6b\x91\x57\x38\x06\xa1\x25\x14\x62\x0b\x16\xe9\x16\xa8\x74\x8b\x1e\x6f\x71\x53\x78\x6f\x2c\x82\xd2\x4b\x33\x87\xcc\xfb\xd2\xcd\x67\xb3\x95\xf2\x75\x94\xa7\xa6\x28\x2a\xad\xfc\x76\x96\x1a\xed\xad\x4a\x2a\x6f\xac\x9b\x49\x5c\x63\x3e\x73\x6a\x35\x11\x36\xcd\x94\xc7\xd4\x57\x16\x67\xa2\x54\x13\x66\x5d\x7b\x86\x8a\x42\x7e\x67\x23\x2e\xb8\xcb\x3d\x5e\x8f\xdc\x23\x3c\x1c\x49\x03\x16\xa0\x80\x02\xe5\x40\xc4\xa3\x41\x8a\x46\xd1\xf4\x8a\xb4\x73\x7f\xb3\x78\x80\xfa\x6a\x36\xc6\xa1\xf6\x59\xef\xcd\x41\xd7\x98\x80\x14\xa6\xf4\x12\x6d\x30\xe2\xd2\x9a\x82\x69\xa2\x96\xa5\x51\xda\xf3\x8f\x34\x57\xa8\x0f\xd5\xef\xaa\xa4\x50\x9e\xec\xfe\x7b\x85\xce\x93\xad\xa6\x70\x2d\xb4\x36\x1e\x12\x84\xaa\x24\xff\x95\x53\x78\xa7\xe1\x5a\x14\x98\x5f\x0b\x87\xdf\xdc\x00\xa4\x69\x37\x21\xc5\x9e\x67\x82\x36\x6a\x1f\x6e\x0e\x5a\x6b\x2d\xd4\x20\xdb\x63\xaf\x06\xf9\x16\x25\xa6\x64\x38\xd2\x1d\x1d\x82\xa5\xb1\x84\x71\x7b\x67\xbb\x63\x93\x9e\xa0\xee\x6b\xa3\x97\x6a\x75\xb8\x76\x70\xe7\x75\x6b\xeb\x2e\x4c\x33\xb3\xa1\xc0\x89\xca\x24\x98\x87\x8d\xf2\x19\xb3\x43\x49\xe7\x88\x24\xc0\x3d\xfe\x5e\x29\xcb\x50\xbb\xff\xf4\x73\xc9\x9c\x8a\x37\x95\x96\x39\x76\xad\x1d\x72\x7a\x15\xb6\x06\x87\xfe\x78\xf3\x1e\x50\x53\x76\x91\x70\x7d\x05\x49\x58\xda\x64\x2a\xcd\x60\xa3\xf2\x1c\x12\xec\x24\x09\x50\x39\x94\x24\xdd\x5a\xe4\x8a\x3c\x2c\x28\x19\xed\x9a\xa2\x21\x25\x56\x97\x41\xe4\x1a\x97\x7a\x24\x06\x32\x4a\x21\xfc\x1c\x92\xad\xef\xbe\xac\xc7\x67\xea\x47\x69\x87\x69\x65\x71\xf1\xa4\xca\x87\xdb\xc5\x27\xb4\x6a\xb9\x3d\x43\x13\xef\xba\xce\x81\x54\x4e\x24\x39\x3a\x78\xb8\x5d\xec\xc9\xb1\xa6\x75\xfa\xe7\x31\xaa\xd6\xcf\x26\x43\xdd\x32\x37\x69\x22\x1a\x3c\xca\x0f\x0f\xf4\x2f\xe5\x48\x18\xa3\x57\x39\x5f\x97\x9a\xca\x8a\x15\x85\x28\xfc\x6a\xaa\x1e\xd2\x11\xa2\x2b\x17\x14\xdd\x58\x51\x3b\x8f\x42\x76\x6b\x36\x28\x2e\x31\x26\x47\xd1\xc5\x33\x9b\x2b\x3d\xc7\x6b\x46\x8f\x71\xef\x63\xf0\x1b\x8b\x4b\xb4\xa8\x09\xe6\x4c\x63\xf9\x14\x39\xc2\x3a\x90\xaf\x7e\x58\x09\x37\xca\x67\x68\xa1\x21\x69\x2c\x3c\x56\x36\x7f\x84\xa2\x72\x0c\x5a\x14\xac\x6a\xa9\x48\x27\x9f\x35\xbc\x23\x0f\xea\xf3\xc3\x0d\x26\x99\x31\x4f\xc4\x96\xad\xb4\xae\x75\xae\x74\x44\xcc\xca\x79\xb4\x63\xfa\xa1\x61\x6b\xaa\xb6\x22\x77\x0c\x4c\x47\x9d\xc4\x87\x63\x0e\xea\x8a\xa0\x67\xed\x30\x8b\x3c\xd2\xe6\xc7\x1a\x8e\xe8\x47\x08\x8d\x9d\xee\xa6\xbb\xe8\xbf\xec\x25\x79\x22\x14\x98\x6b\x2a\x76\xce\xe5\x89\x36\x07\x93\x6a\x30\x65\xa8\xe5\xe0\x97\xfb\x5b\xa6\x72\x16\x0e\x00\xfb\x91\xf6\xa0\x34\x08\xbd\xad\xd3\x50\xf0\x0b\xf2\xf4\x28\xdc\xd7\xc9\x64\xac\x3f\x53\xa6\x87\x0c\x79\x3b\xf8\x4c\xf8\x9a\x77\xc0\xe7\xd2\x10\x60\x25\xdb\x13\x60\x04\x2d\x40\x52\xda\xff\xf0\x97\x13\x6c\x53\xf1\xb3\x42\xdb\xb3\xeb\xf7\x0a\x6d\x0f\x14\x1d\x31\x7e\xf9\xc8\xbb\xd9\x1a\x3b\x53\xd4\xd8\xcc\x4b\x51\x47\x63\x76\x70\x53\x1d\x16\x02\xed\xe7\xf2\xf2\xc7\xcb\xcb\x7d\xfb\x7d\x7b\x2b\x71\xb1\x78\x76\x3c\x2c\x62\x8c\xbb\xc8\x66\x38\x4d\x1c\x55\x0e\xc7\x0c\x24\xf8\x2c\x8a\xb2\x2f\xab\x85\x87\x8a\x97\x71\x28\x61\x08\x27\x76\xc0\x11\x23\x5e\x45\x17\x10\x65\x99\x2b\x94\x20\x1c\x94\x16\x97\xea\x79\x80\x24\x43\x07\xd5\x60\xd1\x0d\xa2\x58\xb3\x19\x5d\x40\x55\xd5\xe1\x25\xda\x10\xde\xf4\x69\x85\x9e\xda\x04\xe1\xee\xaf\x0a\x70\x1b\x31\xa2\x5b\x29\x13\x06\x96\x9e\x25\x0a\x8b\x9e\xa5\x20\xe3\x40\x12\x39\x2a\xc2\xea\xa7\xb2\xf9\x59\xf9\x83\xf1\x7d\xa5\xd6\xb1\x7d\xc9\x4d\xc8\xa4\x35\x06\x8a\xb2\x1c\x93\xe6\x9d\x17\x5a\x0a\x7b\x5c\x00\x85\x87\xa0\x89\xec\x02\xaf\x1e\xff\xb5\xb3\xcb\xbf\x33\xe3\xfc\x9c\xa4\x9b\x31\x9e\xbd\x9e\xc2\xcd\xb3\x48\x7d\xbe\x05\xa3\x19\x65\xf9\xf6\x1e\x92\xa6\x9d\x89\xba\x13\x10\x61\xca\x23\x5d\xf2\x58\xa7\x0f\x72\x03\xce\x81\x3d\x44\xbd\xa1\x6e\x21\xe6\xc4\x3a\x2f\xed\xe7\xa4\xbf\xed\x92\x79\x73\xfd\x52\x61\xde\x27\x7a\x9d\xe9\x99\x1b\x62\x06\x0a\xb5\xca\x98\x5b\xea\x39\xf2\x35\xb5\x57\x4a\x00\x3e\xc7\x76\xec\xed\xdd\x82\x35\x6a\x7a\x0c\xcb\x0d\xa8\x8b\xfd\xc7\x2b\x9c\xae\xa6\x63\x78\x7c\xaa\x12\x9c\xec\xde\x3f\x42\x1a\x1a\x89\x78\x03\x28\x3d\x89\xec\xf7\x90\xa4\x4b\xa9\x5f\x64\xf0\x65\x55\x25\x08\x02\x72\xb1\xc5\xd0\x3a\x29\x93\xb3\xe1\x5f\x4f\x6b\x95\x52\xeb\x23\x72\x67\x7a\x28\xd2\x79\x0d\xef\x3e\x82\x90\xd2\xa2\x73\x6c\x91\xab\x90\xa0\x5a\x50\x19\xfa\x4e\xb5\x84\xd8\x5b\x11\xd9\x21\x8a\x35\x9a\x42\x89\xb6\x50\xce\xa9\x84\xab\x29\x10\xe4\x63\x53\xaa\xc4\x98\xb1\xda\x46\x7c\x9d\xef\xe3\xb1\x14\x8e\x53\xa8\xb0\x89\xf2\x56\xec\xa0\xba\xae\x8e\xd8\xbb\x5b\x88\x36\x06\x01\xc3\x7a\x54\x92\xba\xa9\xa5\x42\x1b\xe4\xf5\x1e\x8b\xd2\x47\x92\xc4\x94\xa0\xff\x5a\xf2\xde\x44\x38\x95\x82\xa8\x7c\x06\x64\x44\xf8\x3c\xa2\x95\x39\xf1\xb4\x31\x56\xfe\xfd\x73\x77\x75\x03\xa4\x3d\xb2\xad\xc8\x73\xb3\x21\x4f\xff\xc9\x8a\x55\x41\x6d\x29\xbc\xfa\x3c\xfa\x6e\x3a\x9d\x7e\x1e\xbd\x66\xad\x86\xec\x53\x0a\x2b\x0a\xf4\xec\x2d\x9f\x47\x3f\x86\xf5\x3e\xcf\xb2\xd8\xa6\x3d\x06\xe4\x9a\xaf\xa7\xd0\x1a\x04\xbd\x01\xfc\x69\x38\x3a\xd1\x9e\x8d\x3e\x36\xbc\x87\x46\x1e\x7d\x8d\x3c\x2d\xb1\xbc\xe1\x8e\x39\x74\x36\x5d\x6d\x96\xd1\x9a\x1a\xf8\xc6\xaa\x21\x1a\x95\xce\x95\x46\xf8\xf5\xea\xfd\xed\xec\x1f\x8b\x0f\x77\x50\x8a\x6d\x6e\x84\x8c\x04\xbd\x15\xda\xe5\xd4\x85\x77\x76\x2f\xde\x00\x61\xfa\x5a\xe4\xe4\xb6\x7c\xbe\x1e\xd0\x44\xec\x69\x71\xcf\x08\x41\x32\xdc\x7d\x78\x00\x87\xa9\xc5\x2e\x50\x36\x16\x42\x6f\x23\xeb\x84\xbf\xa1\x20\xd3\xb2\xc6\xaf\xbb\x9b\x4f\x37\xf7\x2d\x61\x21\x33\xb9\xa4\x0a\xc1\x29\xaf\xd6\x5d\x78\xa1\x74\xc8\x87\xca\xe8\x29\x3c\x18\xd6\x60\x5b\x75\x14\xf0\xa9\xd1\x5e\x10\xe4\x30\x5f\xed\x23\xe3\x0e\x8a\xad\x6a\xfc\xea\xf6\x9f\x57\xbf\x2e\xc0\x79\x63\x31\x90\x6a\x9d\x0d\x51\xb9\x60\x9a\x1d\x0e\x34\x98\x9f\x9e\x27\xcd\x64\x77\x82\x45\x82\x52\xa2\x9c\xd4\x33\x9a\x39\x78\x5b\x1d\x0b\xbb\x77\x88\xe1\xc4\xae\x71\x52\xe9\x27\x6d\x36\x7a\xc2\x16\x70\x9d\x47\x83\xdc\x27\x7c\x71\x11\x95\xd3\xd5\x07\xf0\x8a\x37\x61\x70\x8d\x75\xc2\x68\x06\x1a\x97\x5d\x7d\x95\xae\x07\xb4\xad\x92\x97\xcd\xc9\xc9\xc6\x22\x23\x89\xc8\x1d\x08\xe7\x4c\xaa\xc8\x0f\x9b\x39\x44\x43\xfb\xb8\x1e\x1e\xee\x7f\xfa\x7b\x9f\xfd\x3a\xef\xae\x25\x61\x6c\x1b\x7d\xe7\xfc\x69\x7f\x06\x2f\x4d\xea\x66\xa9\xd1\x29\x96\xde\xcd\xcc\x9a\x52\x24\x6e\x66\x1b\x63\x9f\x94\x5e\x4d\x48\x80\x49\x30\xba\xe3\x79\xbd\x9b\x7d\xc7\xff\xeb\x01\xa4\x87\x0f\x6f\x3f\xcc\xe1\x4a\x4a\x30\xdc\x7c\x56\x0e\x97\x55\x1e\xc2\xc9\x4d\x5b\xa3\xd8\x31\x8f\x03\xc7\x50\x29\xf9\x63\x77\x99\xf6\x67\xd1\x2a\x98\xf7\x81\xc0\x80\x7c\xfb\x14\x66\xdd\x2a\x17\x30\xaa\x3e\xc0\xc1\x10\x23\x2d\xc6\x4d\x82\xbb\xca\x36\x60\x52\x17\x68\x9d\xf0\x80\x45\x28\x3e\xa2\x17\x40\x82\xcb\x10\x84\xb8\x65\x14\x57\xda\xa1\x1d\x80\xae\x40\x82\x63\xf3\x68\x8b\xf2\xd8\x25\xe6\x51\x27\xb0\xaf\x98\x88\xd0\x4a\xaf\x72\x3c\x90\x3e\x62\x43\xb7\x91\xf7\x35\xb1\x27\xb7\x45\x5f\x59\x8d\xb2\x99\xab\x26\xd6\x3c\xa1\x6d\x4b\xdb\x4d\xb3\xa5\x81\x43\x79\xcf\xd0\x66\x77\x8f\xf9\x06\x53\x41\x29\x5c\xaa\x65\x08\x87\xc8\x0d\xf5\x26\x66\xad\x64\x3d\x4f\x76\x14\x39\xe4\x50\xe4\x06\x75\x31\xd9\x57\xd6\xa0\x48\xb3\x28\x27\x88\x16\xe9\xb6\x1a\x9c\xb7\x15\x8f\x6c\xc7\x5c\x3c\x38\xaa\xee\x62\xa9\xdb\x4d\x94\xb8\x78\x91\xff\xd5\xaa\xa1\xfa\x57\x8a\xb2\xbb\xdd\x50\xde\x01\x6a\x6f\xa9\xf7\xf3\x06\x36\x99\xf0\xb8\xe6\xc9\x77\x33\x47\x4a\x8d\x76\x55\x81\x54\x31\x95\x14\xe3\x53\xf8\xa9\x55\x3e\xf5\x32\xdb\x69\x74\x6e\xfa\x77\x26\x0f\x93\xf6\x34\xaf\x64\xa8\xec\x9e\x70\x0b\xa3\x5f\x16\x37\xf7\x77\x57\xef\x6f\x46\xdd\xa4\x93\x2a\x0e\xe0\x6b\xae\x62\x17\x16\x30\x9c\x74\xc9\x38\x1e\xd2\x7d\x3d\x6b\xa8\xb4\x0c\x42\x75\x92\xe4\x6b\xdf\xbe\xf9\x0f\xdd\x3c\x6a\x15\xf7\x06\x32\xb1\xc6\xb6\x2f\xc1\x75\xf8\x1c\xd8\x58\xa2\x97\x68\xd0\x3e\xb7\xa5\xb0\x34\x54\x7b\x91\x2f\x1d\xc6\xd7\x51\x93\x43\x89\xe6\xc0\x71\xf9\x83\xdb\x01\x62\xf5\xb5\x9c\x7f\x8c\x2c\x92\xfc\x3f\xe3\x76\x34\x87\x3f\x46\x14\x64\xa3\x79\x5b\xa9\x30\xf2\x86\xde\xd4\xf2\x7e\xf9\x02\x1f\x74\x68\xcf\x3a\x69\xc6\x74\x71\xc0\xf8\xe5\xa5\x83\x82\x92\x78\xfc\x5e\xb2\xd7\xa7\x75\x61\xf5\xa9\x01\x9e\x90\xf2\x67\xec\x9d\xcf\xec\x7f\x54\xe0\xad\xad\x4f\x37\x20\xf6\xec\x21\x3c\x51\x0b\x4d\xc0\xee\xab\x68\x6f\x97\x4f\xc6\xef\x80\xa9\x45\x5f\x3d\x77\x8e\x30\x91\xee\xd0\xb0\xe4\x68\x50\xd6\xae\x3e\x22\x4f\x42\xf6\x35\xa0\x70\xde\x14\x08\xe2\xf2\x27\x91\x57\x83\xa3\x9b\x23\x6e\x62\xcf\xf4\x4a\x1b\x3d\x49\x94\x16\x76\xfb\x3a\x7e\x6a\x0b\x7c\xf5\xe7\xb8\xe6\x89\xf8\xb3\x8b\xbd\x96\x93\x3f\xe1\xb6\x7f\xe6\x77\xa6\x68\xeb\x17\x0b\x15\x04\x89\x72\xbc\x2a\x0d\x77\x9a\x5b\x20\x19\xc3\x5d\xaf\x4f\x6b\x1d\x0e\xd0\xb5\x4f\x3a\x78\xb7\x84\xc4\xf8\x2c\xde\x26\xf4\x30\xd1\x96\x9d\x38\xd1\x1d\xce\xb5\x02\x15\xe5\x40\xad\xb4\xa1\x5e\x82\x1b\x88\xe6\xd0\x20\x71\xfe\xc8\x41\xa7\x86\x74\x7e\xf2\xd3\x4f\x94\xfe\xb4\x69\x86\xc7\x62\x84\x52\x4f\x3d\x29\xfe\xe4\x84\xab\xc6\x09\xf7\x93\x35\x3d\x18\xd8\x09\x16\xbc\x7f\x10\x31\x0a\xb4\xab\xde\x8a\x15\xa8\x79\x8e\xdf\x80\x43\xae\x0c\x1f\xef\xf1\x59\x39\xdf\xc0\x7a\x53\x95\xb4\x90\xa4\x97\xe4\x57\x23\x4c\x48\x09\xf7\xb8\x7c\x51\x10\x1c\x7d\x2e\xaa\xcb\x84\xbd\x5a\x62\xd0\x9f\x58\x57\xb2\x53\xda\xde\xfa\xf3\x65\xa2\xc1\xc9\x0f\x3a\x1d\xd2\x75\x76\x37\x27\x08\x9c\x85\x34\xd0\xee\xec\x5e\xcc\x52\xe8\x07\xbf\x0d\x5f\x27\xc3\xe5\x8c\x2d\x16\x0b\xb3\xc6\x73\x93\xef\x7d\xbd\x7b\x30\x9a\x02\x4d\x07\xa2\x97\xf7\x63\x9f\xe1\xd8\xea\xc3\x85\x6f\x91\x71\x63\x96\x0d\xbc\x36\xed\xc8\x60\x5c\xc2\xff\x01\x04\xee\x2a\xbc\x33\x0d\x16\x77\x9f\x30\x18\x3b\xf8\x9f\x30\xd8\xa5\x1b\x90\xe5\x1c\xb3\x2d\x07\xa0\xfc\x48\x9a\x9e\x4a\x29\xb0\xff\xb5\x15\x85\x37\x2f\xe3\x03\x37\x81\x97\xf0\xa1\x1d\x07\xf4\x70\x26\x0b\xa7\xdd\x86\x94\xd5\xbb\xd8\x3b\x13\x3f\xe1\x52\x83\xcb\x61\x51\x58\x7b\xd4\xec\xf2\xca\xf0\xe0\xe4\x81\x3a\xe6\x7a\xae\xb9\x14\xa9\xca\x95\x17\x1e\xc9\x2f\x56\x56\x14\xd4\xc7\xa6\x90\x09\x2d\x73\xca\xa2\x94\x54\xa9\x85\x0d\x9f\xa3\x8e\x21\x72\x40\x83\xeb\xae\x3f\xde\x3b\x62\xa7\xfe\xe3\xbd\x6f\xcf\x51\xb7\x25\x27\x7b\x7f\xc7\x74\x71\xd2\x06\x47\x2f\x79\xfc\x29\x5b\x03\x4f\xaa\xeb\xc4\xaa\x3d\x3d\x75\x55\xb2\xfb\x93\xbb\x39\xfc\xf1\xe5\xe2\x7f\x01\x00\x00\xff\xff\x03\x3a\xaa\xc6\xc9\x2b\x00\x00")

func appcatalogAppscodeCom_appbindingsV1YamlBytes() ([]byte, error) {
	return bindataRead(
		_appcatalogAppscodeCom_appbindingsV1Yaml,
		"appcatalog.appscode.com_appbindings.v1.yaml",
	)
}

func appcatalogAppscodeCom_appbindingsV1Yaml() (*asset, error) {
	bytes, err := appcatalogAppscodeCom_appbindingsV1YamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "appcatalog.appscode.com_appbindings.v1.yaml", size: 11209, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _appcatalogAppscodeCom_appbindingsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x5a\xdd\x6f\xdb\x38\xb6\x7f\xcf\x5f\x71\xe0\x79\x48\x0b\xf8\x03\xbd\xf3\x72\xe1\xfb\x30\x37\x4d\x53\xa0\x33\x69\x5a\xc4\x99\x5e\x0c\x6e\x17\x1b\x4a\x3c\xb2\xb8\xa1\x48\x0d\x49\xd9\xf1\x0e\xe6\x7f\x5f\x1c\x92\xb2\x64\x9b\x92\xd3\xd9\x19\xac\x1f\xda\x58\x24\x0f\xcf\xe7\xef\x7c\xc8\xac\x16\x5f\xd0\x58\xa1\xd5\x12\x58\x2d\xf0\xd9\xa1\xa2\x6f\x76\xfe\xf4\xdf\x76\x2e\xf4\x62\xf3\x26\x43\xc7\xde\x5c\x3c\x09\xc5\x97\x70\xdd\x58\xa7\xab\x7b\xb4\xba\x31\x39\xbe\xc3\x42\x28\xe1\x84\x56\x17\x15\x3a\xc6\x99\x63\xcb\x0b\x80\xdc\x20\xa3\x87\x0f\xa2\x42\xeb\x58\x55\x2f\x41\x35\x52\x5e\x00\x48\x96\xa1\xb4\xb4\x07\x80\xd5\xf5\xfc\xa9\xc9\xd0\x28\x74\xe8\xaf\x52\xac\xc2\x25\xe4\xcc\x31\xa9\xd7\x17\x00\xe1\x3b\xab\xeb\x4c\x28\x2e\xd4\xda\xce\x59\x5d\xc7\x65\xfa\xd3\xe6\x9a\xe3\x3c\xd7\xd5\x85\xad\x31\x27\xaa\x8c\x73\xcf\x0e\x93\x9f\x8d\x50\x0e\xcd\xb5\x96\x4d\xa5\xfc\x8d\x33\xf8\x71\xf5\xe9\xee\x33\x73\xe5\x12\xe6\x74\x60\xee\x76\x35\x7a\x56\xc2\x45\x0f\xed\x57\x7a\xbe\x04\xeb\x8c\x50\xeb\xe4\xc1\x4d\xd0\x58\xef\xec\x97\xde\x93\xb1\xe3\xad\x9a\xe6\x27\x3a\xea\x11\xbb\x5a\xf7\xf9\xe0\xcc\xd1\xd7\xb5\xd1\x4d\xed\xb5\x91\xd4\x40\x38\x1b\x55\x9b\x33\x87\x6b\x6d\x44\xfb\x7d\xd6\x53\x2a\x7d\x6b\x4f\xb6\x5f\xbd\x6d\x00\x82\x89\xaf\xea\xfa\x6d\xd0\xb7\x7f\x28\x85\x75\x3f\x1d\x2d\xdc\x0a\xeb\xfc\x62\x2d\x1b\xc3\xe4\x81\x8d\xfc\x73\x2b\xd4\xba\x91\xcc\xf4\x57\x2e\x00\x6a\x83\x16\xcd\x06\x7f\x56\x4f\x4a\x6f\xd5\x7b\x81\x92\xdb\x25\x14\x4c\x5a\xe2\xc5\xe6\x9a\x04\xbe\x23\x41\x6a\x96\x23\xa7\x67\x4d\x66\xa2\xb7\xd9\x25\xfc\xf6\xfb\x05\xc0\x86\x49\xc1\xbd\xf2\x82\x74\xba\x46\x75\xf5\xf9\xc3\x97\xef\x57\x79\x89\x15\x0b\x0f\xe9\x32\x5d\xa3\x71\x7b\x25\x04\x9f\xdb\x7b\xfb\xfe\x19\x00\x47\x9b\x1b\x51\x7b\x8a\x70\x49\xa4\xc2\x1e\xe0\xe4\xdf\x68\xc1\x95\x08\xd1\xe6\xc8\xc1\xfa\x6b\x40\x17\xe0\x4a\x61\xc1\xa0\x17\x4b\x39\xcf\x52\x8f\x2c\xd0\x16\xa6\x40\x67\xff\xc0\xdc\xcd\x61\x45\xa2\x1b\x0b\xb6\xd4\x8d\xe4\x90\x6b\xb5\x41\xe3\xc0\x60\xae\xd7\x4a\xfc\x73\x4f\xd9\x82\xd3\xfe\x4a\xc9\x1c\x46\x45\xb7\x1f\xef\xd4\x8a\x49\x52\x42\x83\x53\x60\x8a\x43\xc5\x76\x60\x90\xee\x80\x46\xf5\xa8\xf9\x2d\x76\x0e\x1f\xb5\x41\x10\xaa\xd0\x4b\x28\x9d\xab\xed\x72\xb1\x58\x0b\xd7\xc6\x77\xae\xab\xaa\x51\xc2\xed\x16\xb9\x56\xce\x88\xac\x71\xda\xd8\x05\xc7\x0d\xca\x85\x15\xeb\x19\x33\x79\x29\x1c\xe6\xae\x31\xb8\x60\xb5\x98\x79\xc6\x95\xf3\x20\x51\xf1\xef\xf6\xe6\xb9\xec\x71\x7a\x14\x03\xe1\xe3\xfd\x6b\x50\xef\xe4\x64\x20\x2c\xb0\x78\x2c\xf0\xdf\xa9\x97\x1e\x91\x56\xee\x6f\x56\x0f\xd0\x5e\xea\x4d\x70\xa8\x73\xaf\xed\xee\x98\xed\x14\x4f\x8a\x12\xaa\x40\x13\x0c\x57\x18\x5d\x79\x8a\xa8\x78\xad\x85\x72\xfe\x4b\x2e\x05\xaa\x43\xa5\xdb\x26\xab\x84\x23\x4b\xff\xda\xa0\x75\x64\x9f\x39\x5c\x33\xa5\xb4\x83\x0c\xa1\xa9\x29\x44\xf9\x1c\x3e\x28\xb8\x66\x15\xca\x6b\x66\xf1\x2f\x57\x3b\x69\xd8\xce\x48\xa5\xe7\x15\xdf\x07\xe7\xc3\x8d\x41\x5b\xfb\xc7\x2d\x8e\x26\x2d\xd4\xc5\xff\xaa\xc6\x9c\x4c\x45\xfa\xa2\x23\x50\x68\x43\x91\xde\x3b\x99\x8a\x3e\x0f\x4d\x5e\xbd\xd7\x5a\x15\x62\x7d\xb8\x72\x74\xdb\x75\x6f\xe3\x3e\x10\x4b\xbd\xa5\xe0\x88\xca\x23\x98\x83\xad\x70\xa5\x67\x84\xf2\x09\xdc\xe3\xaf\x8d\x30\x1e\x39\xfa\x9f\x21\x6e\x3c\x47\xec\x6d\xa3\xb8\xc4\xd3\x95\x63\x8e\xae\xc2\xc6\xe0\xa4\x9f\x6f\x3e\x02\x2a\x42\x51\x0e\xd7\x57\x90\x85\xa5\x6d\x29\xf2\x12\xb6\x42\x4a\xef\x19\xf6\x84\x93\xa8\x7c\xdd\xa2\x18\x06\x25\xa2\xd9\x90\x7f\xe7\xc4\x64\x11\x04\x6b\xf1\x85\xe4\x4a\x10\x29\xb4\xa9\x98\x5b\x42\xb6\x73\x98\x58\x4e\xfa\x41\xfb\x11\xca\x62\xde\x18\x5c\x3d\x89\xfa\xe1\x76\xf5\x05\x8d\x28\x76\x67\xe5\xff\x90\x3a\x05\x5c\x58\x96\x49\xb4\xf0\x70\xbb\x3a\xe0\x7f\x43\xeb\xf4\xe7\x31\x2a\xb6\x9f\x6d\x89\xaa\x67\x4a\x92\x3f\x1a\x33\x4a\x0d\x0f\xf4\x97\xb0\x24\x86\x56\x6b\xe9\x2f\xcb\x75\x63\xd8\x9a\xc2\x0d\x7e\xd1\x4d\x92\x70\x04\xd8\xc6\x06\xe5\x76\x76\x53\xd6\x21\xe3\x29\x6d\x06\x75\x65\x5a\x4b\x64\xa7\xdc\x7a\xf3\xe4\xe7\x3d\x64\xf2\x18\x77\x3e\x06\x1f\x31\x58\xa0\x41\x45\x30\xa5\x3b\x3b\xe7\xe8\xe3\x65\xcc\xb8\x00\x37\xc2\x95\x68\xa0\x23\xa8\x0d\x3c\x36\x46\x3e\x42\xd5\x58\x0f\x3b\x14\x78\xa2\x10\xa4\x89\xaf\x0a\x3e\x14\xfe\x82\x2d\x66\xa5\xd6\x4f\x49\x92\x94\xab\x1a\xa5\x5a\x3d\x0b\x15\xf1\xae\xb1\x0e\xcd\x94\xbe\x28\xd8\xe9\xa6\xaf\xbe\xfd\xf5\xf3\x49\x82\xe4\x58\x54\x41\x5b\xcd\x24\x57\x8e\xb1\xff\x91\xb6\x3e\xb6\x90\x42\x5f\x82\xfb\xef\x35\xd6\x45\xf6\xe5\x00\xc1\x51\x87\xf7\xdc\x52\x05\xf6\x32\x6e\x68\x6b\x30\xa1\x02\x5d\x87\x82\x12\x7e\xbe\xbf\xf5\x34\x8e\x62\xdc\x1e\x67\x8b\x03\x95\x2b\x60\x6a\xd7\x26\x8e\xe0\x05\xe4\xcf\x51\xa8\x3f\x2e\x8b\x36\xee\x45\xb2\x3c\x94\xe8\x37\x83\x2b\x99\xdb\xf3\x8c\xcf\xb5\xb6\xc8\x21\xdb\x9d\xf1\xc2\x0e\x66\x84\x72\xdf\xff\xd7\x28\xbb\x54\x9a\xac\xd1\x24\xf7\xfc\xda\xa0\x49\x02\xcc\x09\xc3\x97\x8f\x7e\xaf\xd7\xfe\x5e\xf5\x2d\xce\xfa\xa5\xa8\x97\xa9\x77\x62\xdd\x0c\x2b\xff\xf2\xf2\x87\xcb\xcb\x84\xb5\xfe\x32\xab\xf8\xf2\xed\x85\x1e\xbf\x8a\xd1\x6b\x23\x83\xe1\x2c\xf1\xd2\x58\x9c\x7a\x80\xc0\x67\x56\xd5\x12\x43\xf9\x30\x1d\x14\xd3\x17\x17\x14\xff\x7b\x40\x88\xb1\x2c\xa2\xc1\x59\x5d\x4b\x81\x1c\x98\xa5\x02\xbc\x10\xcf\xe0\x43\xff\xa8\x6e\xea\x7f\x5a\xa3\x47\x81\x16\x0b\x22\x4f\xd5\xce\xf1\x15\x4a\x13\x8e\xac\xf7\xea\x0d\xf4\xff\x70\x90\x9a\x18\xe3\x29\x15\xce\x3c\x2c\x24\x17\xc8\xc1\x93\x0b\x81\xff\x41\xb8\x3f\x2a\x7e\xda\x4f\x63\xe4\x0b\x90\xde\x63\xf1\x5a\x6c\x62\x7b\x20\x75\xc8\x74\x2d\x6e\xb1\xba\x9e\x92\x9e\xad\x63\x8a\x33\xc3\x09\x3e\x92\x4a\x21\x5d\xc3\xab\xc7\xff\xdf\xeb\xfa\x6f\xa5\xb6\x6e\x49\x32\x2d\x3c\x0e\xbd\x9e\xc3\xcd\x33\xcb\x9d\xdc\x81\x56\x1e\x17\xc3\xdd\xba\x97\x1d\x92\x94\xd3\x89\x82\x10\xe1\x91\xae\x78\x6c\x81\x9e\x0c\xeb\x33\x15\x79\x1f\x6b\xc3\x20\x49\xb2\xcd\x1f\x87\xb9\xe3\x7f\xf6\xa9\xb6\x4b\x57\x05\xf5\x76\xfb\x8c\xeb\x6f\xa5\x4b\xd3\x8c\x8a\x75\xe9\x39\xa5\xaa\x5e\x6e\xa8\x75\x11\x0c\xf0\x39\xb6\x3a\xef\xee\x56\x5e\x93\xba\x22\xb5\x0a\x1b\xab\xf9\x57\x38\x5f\xcf\xa7\xf0\xf8\xd4\x64\x38\xdb\x3f\x4f\xab\x22\x0f\xc5\x7a\xa4\x0f\x42\xcd\x22\xeb\x9e\x38\x75\x5c\x1e\x1e\xbd\x3a\x32\x04\x06\x92\xed\x30\x34\x21\x42\x4b\x6f\xd8\xd7\x69\x84\x8c\xaa\xa4\xd6\x82\x49\xab\xfd\x69\x05\x1f\x3e\x03\xe3\xdc\xa0\xb5\x5e\xe7\x57\x21\x71\xf4\x20\x2d\x74\x6e\xa2\x48\xa3\x7b\xe8\x5c\x3c\x51\x4f\xaf\xc5\x3c\xa8\xd1\x54\xc2\x5a\x91\xf9\x6a\x06\x18\x79\xd5\x9c\xea\x20\xbf\x37\x5a\x61\x30\xfb\x91\x7d\x6b\x66\x7d\x5a\x63\x26\x13\xce\xb0\x3d\x9c\xb6\x15\x8a\xf7\xdb\x1e\xfa\x4c\x81\xb5\x66\x4e\x17\x15\x9c\x7a\x92\x42\xa0\x09\x92\x3a\x87\x55\xed\x22\x41\x62\x88\xd1\xbf\x86\xbc\x35\x63\x56\xe4\xc0\x1a\x57\x02\x99\x0e\xbe\x4e\x68\x65\x49\x1c\x6d\xb5\xe1\xff\xfb\x35\x55\x63\xf8\xb2\x85\x6c\xc7\xa4\xd4\x5b\xf2\xe1\xf7\x86\xad\x2b\x6a\xec\xe0\xd5\xd7\xc9\x77\xf3\xf9\xfc\xeb\xe4\xb5\xd7\x66\xc8\x0e\x35\x33\xac\x42\xe7\x3d\xe4\xeb\xe4\x87\xb0\x9e\x24\xcc\x0c\xf6\x29\x4f\x01\x7d\xcd\x95\x2c\x75\x46\x80\x6b\x10\x4b\x3a\x4e\x46\x1b\x9d\xc9\xe7\x8e\xe3\xd0\xfe\xa2\x6b\x51\xa4\x27\x8c\xd3\x6d\x47\x11\x3a\x20\xa5\x52\xd8\xd5\x59\x31\xc4\x9c\x50\x52\x28\x84\x5f\xae\x3e\xde\x2e\x7e\x5c\x7d\xba\x83\x9a\xed\xa4\x66\x3c\x92\x73\x86\x29\x2b\xa9\x7b\xa5\xf4\xad\x81\xf0\x77\xc3\x64\xaa\xa4\xf1\xa7\xdb\x51\x46\xc4\x91\x1e\xe7\x31\xde\x2d\xdc\x7d\x7a\x00\x8b\xb9\x21\x21\x0c\x84\x8e\x81\xc7\x94\x7b\x42\x74\x4b\x61\xa3\x78\x8b\x44\x77\x37\x5f\x6e\xee\xfb\x62\x96\x5a\x72\xca\xd9\x56\x38\xb1\x09\xdd\x34\x65\x26\xa1\xd5\x1c\x1e\x34\x69\xea\x84\x64\x5f\x65\x14\xd4\xd4\x5e\x33\x82\x8f\xc0\x53\x8f\xc4\xb4\x5f\xed\x5e\xdd\xfe\xdf\xd5\x2f\x2b\xb0\x4e\x9b\xd3\x00\xf2\x84\x7a\x27\x43\xec\xad\x3c\xc5\x13\x77\x19\xc9\x2d\xcf\xb3\x6e\xe0\x39\xc3\x2a\x43\xce\x91\xcf\xda\x59\xc6\x12\x9c\x69\x8e\x2f\x3f\x38\xd2\xce\xcf\x66\x4d\x18\xa0\xcd\x8a\x38\x41\x3b\x39\x18\xa4\x1d\xf5\xbb\x55\x54\x48\xaa\xe6\xf6\x2b\xe4\x66\x06\xa9\x95\x8b\x70\xdf\x0d\x00\x2e\x4f\x6b\x07\xd5\x4e\xed\x7a\xa5\xa6\x37\x9f\x4f\x14\x06\x3d\x4e\x30\x69\x81\x59\xab\x73\xe1\x7d\x6e\xdf\xbb\x77\x94\x8f\x51\x76\xac\xc7\x18\xea\x2f\x0e\x2b\xad\xbb\x9e\x64\xb1\x21\x73\xc9\xe9\xcc\xe1\x30\x9a\xeb\xdc\x2e\x72\xad\x72\xac\x9d\x5d\xe8\x0d\x25\x36\xdc\x2e\xb6\xda\x3c\x09\xb5\x9e\x11\xeb\xb3\x60\x64\xeb\x07\xd7\x76\xf1\x9d\xff\x2f\x09\x35\x0f\x9f\xde\x7d\x5a\xc2\x15\xe7\xa0\x7d\x5b\xd7\x58\x2c\x1a\x19\x82\xc6\xce\x7b\x63\xc9\xa9\x1f\x92\x4d\xa1\x11\xfc\x87\x54\x11\xf5\x47\x70\x28\x98\xf3\x81\x42\x9d\x3c\x78\x1c\x8d\x6e\x85\x0d\xe8\xd3\x6e\xf7\x0e\x1f\x63\x29\xc6\x4a\x86\xfb\x9a\x32\xe2\x4d\xcf\xbe\x27\x4c\xa7\xec\xbd\x0a\x65\x42\xb4\x39\x64\x58\x90\x39\x5c\x89\x3b\x8f\xca\x42\x59\x34\x7b\x50\x4a\xa5\xb4\x18\x7b\x47\xcf\x85\xc3\x53\xf1\x4e\x2a\xef\x43\x75\x44\xcc\x15\x6a\x2d\xf1\x48\xea\x18\xf7\xb6\x95\x36\x65\x8f\x13\xf9\xc1\xa0\x6b\x8c\x42\xde\xcd\x17\x33\xa3\x9f\xd0\x0c\x4a\x99\x20\xdb\xca\xdd\x06\xe9\x79\x1d\xce\xe1\x2d\xe6\x8c\x12\x2e\x17\x45\x70\xf2\x04\xdd\xc0\x09\xf5\x01\x7a\x23\x78\x3b\x51\xb5\x14\x21\xe4\x3e\x64\xf8\x76\x44\x41\x05\x05\xb2\xbc\x8c\xf2\x00\x1b\x25\xdc\x57\x80\x75\xa6\xf1\x63\xcb\xa9\x4f\xfd\x96\xaa\xaf\x58\x84\xee\xfc\x7d\x29\xdf\x4a\xd0\x1c\xf4\xb6\xd5\x1e\x9f\x18\x67\xb5\x03\xe1\x2c\xa0\x72\x86\xba\x29\xa7\x61\x5b\x32\x87\x9b\x64\xbd\xd2\x9f\xc1\xe4\x5a\xd9\xa6\x42\xaa\x74\x6a\x8a\xe2\x39\xbc\xef\x97\x3d\x43\x66\x4d\x69\x75\xd7\x37\x73\x98\x32\xe7\xb2\xe1\xa1\x26\x7e\xc2\x1d\x4c\x7e\x5e\xdd\xdc\xdf\x5d\x7d\xbc\x99\x4c\x21\x6b\xe2\xa0\xb9\xbd\x3f\x76\x3d\x29\xe4\xa0\x7d\xa4\x43\x8f\xce\x21\x65\xb7\xbd\x7b\xa3\xb8\x1f\x64\xc7\x0b\xde\xbd\xfd\x3b\xdd\x31\xe9\x95\xdc\x1a\x4a\xb6\x49\x76\x3f\x9d\xf7\xc0\x75\x78\x31\xd4\xd9\xa4\xa7\xe1\xa0\x84\x42\x53\x7d\x44\xbe\x72\x14\x39\x09\xca\x27\x2d\x07\xa5\x8e\x23\x47\xf5\x6f\xd0\x8e\x30\x69\x09\x33\xf8\x6d\x62\x90\xe4\xfc\x09\x77\x93\x14\xaa\xff\x36\xa1\x80\x9a\x2c\x0f\x94\x39\x71\x9a\x9e\xb4\xd2\xff\xfe\x3b\x7c\x52\x5d\xa3\xd4\x89\xb2\xbf\xe9\x32\x91\xba\x00\x2a\x4a\xc6\xf1\x0d\xc1\x41\xc7\x74\x8a\xc1\xe3\x43\x2f\xc6\xf9\x4f\x38\x30\xe9\x38\x1c\xa6\xfb\x8d\xbd\xd7\x14\xc0\x0e\x6c\xc0\x1c\xd1\x0a\xa5\xfa\xfe\xa5\xe6\x40\x57\x4d\x0e\x90\x00\xa2\x20\xf9\x40\x87\x31\x3e\xb9\xf3\x34\x87\x96\x12\x23\xa6\x7e\xf5\x10\xb9\x61\x3c\x3d\xfa\x86\x97\x4c\x52\x20\x2e\x7e\x61\xb2\x19\x1c\xa8\x24\xf8\x88\xbd\xcc\x2b\xa5\xd5\x2c\x13\x8a\x99\xdd\xeb\xf8\x1a\x29\x70\x74\x88\x20\x83\x74\xa1\x17\x5d\x9d\x2b\x3f\xe1\x6e\x68\x4a\xf6\x22\x91\x36\xdf\x28\x4c\x10\x20\xf2\xff\xaa\xd6\xbe\xef\xdb\x01\xc9\x16\xee\x79\x7d\x4e\xcf\x70\x84\x98\x43\x52\xc1\x87\x02\x32\xed\xca\x78\x17\x53\x63\x24\x7b\x96\xf1\x69\xec\x78\x26\x14\x68\x08\x0b\x62\xad\x34\xd5\xfe\xbe\xc0\xef\x0e\x8d\x90\xf6\xa3\x7e\x3a\x33\xac\xe7\x33\xaf\x3d\xa2\xd4\xe7\x8c\x31\x36\x68\x02\x98\x91\x52\xd2\xef\x6d\xc6\xe6\x46\x6d\xf4\xdb\xf7\x46\x57\x2f\x86\x00\xbf\x7b\x14\x07\x2a\x34\x6b\xb4\xfb\xb7\xf4\x09\xae\xfc\x3b\xcc\x90\xff\xc2\x2b\x67\x7c\x16\xd6\x75\x90\xdd\xd5\x13\x7f\x1a\x3e\x04\x00\xbf\xc7\xe2\x1b\x1c\xfa\xe4\x75\x48\x9b\xc8\x0f\x6b\x4b\x2f\xef\x98\x07\x8e\x48\x33\xec\x38\xe7\x45\x82\x33\xaf\x2c\x12\x52\x25\x7b\x8b\xd1\xe3\x2f\xc0\x09\xe8\x77\x53\xdf\xc8\x4c\xe8\xc0\xfe\x7c\x8e\xce\x38\xfe\xd9\x0d\x06\x2b\xbd\xc1\x97\xa5\xc6\xfb\x76\xef\x68\x54\x04\x8a\xb4\x30\xd6\x7c\x84\x4f\xf4\x33\x8a\x91\x74\x5c\xff\xd9\xf9\x30\xe6\xc0\xc0\x63\xd7\x08\x9c\x49\x3b\xff\x29\xd8\xda\xd7\x5d\x2f\x32\x4e\xdc\x7b\xc6\x38\xde\x81\xbf\xd9\x38\x97\x76\x50\x86\xf3\x26\x2a\x06\x61\xf7\x44\x8a\x81\x9a\x25\xb0\xfd\xef\xe4\x78\xa7\xbf\x85\x03\xdc\x06\x2e\xc2\xeb\x60\x1c\x94\xfd\x45\x97\x9f\x73\x0e\x52\xcf\xc0\x92\xd3\xdf\xee\x36\x23\x8b\x61\x89\x19\xc3\x0e\xc5\xf1\xcf\xc7\x86\x0f\x0f\xd4\x83\xb6\xb3\xbf\x82\xe5\x42\x0a\xc7\x1c\x92\xed\xd7\x86\x55\xd4\x2f\xe6\x50\x32\xc5\x25\xe5\x36\x4a\x75\xd4\x18\x86\x17\x2f\xc7\x20\x37\xa8\xaf\xcd\xe9\x0f\xc0\x4e\x18\x69\x7f\x00\xf6\xd7\xf2\x92\xb2\xd8\xec\xe0\x17\x32\x17\xa3\xfa\x3e\x7a\xd4\x0a\x06\x9b\x37\x4c\xd6\x25\x7b\xd3\x3d\x8b\x3f\x80\x0c\x3f\x2f\xec\x2d\x87\x9f\x36\x20\xef\x4d\x10\xa9\x00\x63\xeb\x76\x18\xf9\xaf\x00\x00\x00\xff\xff\x48\x78\xa3\xf7\x1f\x2a\x00\x00")

func appcatalogAppscodeCom_appbindingsYamlBytes() ([]byte, error) {
	return bindataRead(
		_appcatalogAppscodeCom_appbindingsYaml,
		"appcatalog.appscode.com_appbindings.yaml",
	)
}

func appcatalogAppscodeCom_appbindingsYaml() (*asset, error) {
	bytes, err := appcatalogAppscodeCom_appbindingsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "appcatalog.appscode.com_appbindings.yaml", size: 10783, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"appcatalog.appscode.com_appbindings.v1.yaml": appcatalogAppscodeCom_appbindingsV1Yaml,
	"appcatalog.appscode.com_appbindings.yaml":    appcatalogAppscodeCom_appbindingsYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"appcatalog.appscode.com_appbindings.v1.yaml": {appcatalogAppscodeCom_appbindingsV1Yaml, map[string]*bintree{}},
	"appcatalog.appscode.com_appbindings.yaml":    {appcatalogAppscodeCom_appbindingsYaml, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
