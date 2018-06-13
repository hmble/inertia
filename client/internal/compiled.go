// Code generated by go-bindata.
// sources:
// client/scripts/daemon-down.sh
// client/scripts/daemon-up.sh
// client/scripts/docker.sh
// client/scripts/keygen.sh
// client/scripts/token.sh
// DO NOT EDIT!

package internal

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
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _clientScriptsDaemonDownSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\xce\xb1\x6e\xc2\x30\x10\xc6\xf1\xfd\x9e\xe2\x2b\xb0\x1a\x9e\x20\x43\x2a\x22\x54\xa9\xb8\x12\x52\x87\x4e\x60\xe2\x73\x38\xd1\x9c\xa9\x7d\xa8\xaf\x5f\x45\x74\x88\xd8\x6e\xf8\xdf\x4f\xdf\xf2\x65\x73\x16\xdd\xd4\x0b\xd1\x12\xaf\xa1\x4a\x8f\xda\x17\xb9\x19\x52\x2e\x38\x17\xd1\x41\x74\x40\xcc\xbf\x0a\xbb\x30\x62\xe0\x31\xeb\x9a\xa8\xb2\xc1\x31\xd1\xb6\xed\xf6\x1f\xfe\xe8\xdb\x7d\xd7\x88\x72\x31\x09\xee\x11\x4d\xe2\x8e\xed\xff\x05\x7d\x56\x0b\x53\x81\xa0\x11\x16\xae\x0c\xb1\x07\x2c\x69\x3a\xa5\xa2\xdc\x55\x45\x87\x35\xb5\xef\x87\xae\xdd\x7e\x1d\x0f\x9f\xde\xbf\xf9\x5d\x73\xaa\xf7\x98\x11\x73\x7f\xe5\x82\x5b\x85\xfb\x81\x73\x49\xbe\x8d\x0b\x16\x1a\x46\x6e\x56\xb3\x21\x8b\x13\xcd\xfb\x32\xc2\x25\xac\x9e\x4c\xfa\x0b\x00\x00\xff\xff\x62\xed\x0f\x4a\xfb\x00\x00\x00")

func clientScriptsDaemonDownShBytes() ([]byte, error) {
	return bindataRead(
		_clientScriptsDaemonDownSh,
		"client/scripts/daemon-down.sh",
	)
}

func clientScriptsDaemonDownSh() (*asset, error) {
	bytes, err := clientScriptsDaemonDownShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "client/scripts/daemon-down.sh", size: 251, mode: os.FileMode(420), modTime: time.Unix(1528351521, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _clientScriptsDaemonUpSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x55\xdd\x6e\xe3\x36\x13\xbd\xe7\x53\x4c\x64\xe7\xcb\x57\xa0\xb2\xba\x9b\x5e\x79\xa1\x05\xdc\xb5\xb0\x09\x1a\xdb\x81\x9d\x6d\x51\xa4\x81\x97\x16\xc7\x16\x6b\x89\xd4\x92\x23\xab\xed\xd3\x17\xa4\x24\xcb\xce\x4f\x7b\xa5\x84\x33\x73\x78\xe6\xf0\xcc\x78\x70\x11\x6d\xa4\x8a\x6c\xc6\xd8\x00\x7e\xe2\x56\xa6\x60\x53\x23\x4b\x82\xad\x36\x60\x91\x48\xaa\x1d\x54\x25\xdc\x2a\x34\x24\x39\x18\xfc\x56\x49\x83\x05\x2a\xb2\xf0\x7f\x21\x0d\xa6\xa4\x8d\x44\xfb\x3d\x20\xa5\xdf\xb1\x01\x70\x25\x60\x63\xa4\x72\x85\x94\x21\x08\x8e\x85\x56\xa0\x55\x2e\x15\x8e\x18\xb3\x48\x10\xa2\xbb\xef\x8b\x45\x03\xdc\xec\x2a\x8f\x36\x62\xd3\x49\x32\x5b\xcc\xd7\xcb\xe4\x2e\x99\xac\x92\xf8\xf2\xf1\xdd\x93\xed\x0e\xef\x17\xcb\x87\xf8\xf2\xf1\xfd\x93\x65\x37\x8b\xd5\xc3\x7a\x32\x9d\x2e\x93\xd5\x2a\xbe\x7c\xbc\x7e\xb2\x0e\xac\x23\x28\x0b\xbe\x43\x10\x48\x5c\xe6\x3d\xe6\x7c\x32\x4b\x62\xd9\xa4\x84\x0d\x23\x76\x3b\x9b\x7c\x4e\xe2\x6a\x93\xe6\xbc\x52\x69\x56\x72\x11\xb5\x19\xe3\xe1\x39\x15\x8f\x4f\x20\x34\x5a\x75\x45\x50\x70\x22\x34\x50\x67\x9c\xa0\xd4\x86\x4e\xdb\x34\x95\xb2\xa0\x15\x48\xe5\x4f\x53\xad\x88\x3b\x54\x27\x8c\x85\x5c\xab\x9d\xfb\x4a\x02\x69\xa1\xe0\x65\x89\x02\x48\xb7\xa9\xc6\x89\x09\x27\x0d\x8f\xd8\xa7\xc5\xfc\x61\x72\x3b\x4f\x96\x8d\x00\x3f\x5e\xff\x70\x7d\x54\xae\x34\xfa\x0f\x4c\x89\x15\x7b\x21\x0d\x84\x25\x0c\x6f\x16\xb3\xa4\xeb\xe1\x8d\xe3\xa8\xab\x3a\x91\x4c\x70\x7a\x33\xdd\xc7\xd8\x00\x3e\x69\xb5\x95\xbb\xca\x70\x92\x5a\xbd\x95\x9c\xfa\xa4\x7f\x8f\x46\xd6\xe6\x1e\x30\xc3\x74\x0f\x72\x0b\x3c\x37\xc8\xc5\x5f\x4e\x39\xef\x19\xe7\x1f\xe2\x7b\x04\xa1\x6b\x05\xf8\xa7\xb4\xde\x83\x8d\xbe\x23\x36\xb9\x5b\x26\x93\xe9\x6f\xeb\xe5\x97\xf9\xfc\x76\xfe\x39\xfe\x6a\x2b\xa1\x41\xe8\x74\xef\x24\xb1\x10\x7e\x83\x30\xdc\xca\xdc\xbd\x50\xa0\x78\x81\xf1\xf0\xc4\x03\xc1\x57\x26\xb7\xf0\x08\x17\x10\xfe\x0d\xc1\xf0\x19\x58\x00\x4f\x1f\xdc\x5b\x28\x06\x00\x80\x69\xa6\x21\xb8\xaf\x9a\x19\x38\x12\xe9\x55\xf3\x0f\x4e\x1a\x6c\x8e\x58\x06\xbe\xe6\x94\x8c\x29\x20\xdc\xc2\xf3\x3b\xe0\x23\x44\x02\x0f\x91\xaa\xf2\x1c\xde\x7f\xfc\xdf\x3b\xb6\x95\x1f\x58\xc3\x2a\x78\xe6\xbb\x00\x2e\x62\x08\x08\x2d\x9d\x33\x1b\xc0\x54\xd7\x2a\xd7\x5c\xf8\x79\x44\x4b\x28\x3a\x3e\xde\xff\xa3\x93\x06\xba\x54\xc7\x7d\xe8\x3d\xff\x92\x6a\xe9\xc8\x34\xc1\x97\xfc\x30\xb7\xd8\x5e\x7b\xe7\xae\x74\x7c\x60\x53\xc9\x5c\x00\xb9\x11\xb0\x99\xae\x72\x01\x19\x3f\x20\x6c\x10\x15\xd8\xb4\xbc\x12\x20\x15\xe9\xb6\xcc\xd9\xfb\x97\xfb\x15\x70\x82\xa8\xa1\x19\xbe\xa0\x79\xf7\x1f\x14\x7d\xb7\xa1\x3c\x07\x78\x4d\x4c\x67\xae\x65\xa5\xfa\xd1\x83\x5a\x52\x06\x3c\x4d\xd1\xda\x6e\xd8\x32\x6d\xa9\x43\xb6\xee\x43\xde\x77\x6c\x00\x06\x73\x3c\x70\x45\x6d\x4a\xbf\xe0\x5c\x29\xcf\x73\x5d\xfb\xd5\xd8\xa3\xbb\xbf\x8c\xce\x47\x6c\x00\x2b\x44\x8f\xee\x1e\x7c\x96\xf8\xbc\x42\x9b\xe3\x36\x72\x7b\x21\xd3\x35\x50\x26\x2d\xd4\xda\xec\xed\x98\x0d\x20\x23\x2a\xed\x38\x8a\x76\x92\xb2\x6a\x33\x4a\x75\x11\xbd\xb6\x93\x06\x99\xae\x43\x49\xa1\xaf\x63\x8d\x66\xcb\x76\x64\x8e\x2b\xb6\x59\x47\xc3\x93\x15\x12\xb0\x33\x4b\x56\x0a\x42\x01\x61\x68\x0a\xf8\xdd\x4b\x1c\x96\xbd\xeb\x7c\xfe\x38\x18\x9e\x2f\x9d\xa0\xcb\x3c\x40\x74\xe0\x26\x32\x95\x8a\x1a\xb8\x91\x53\x6e\xfc\xda\x61\x5f\x12\xf8\x1d\x10\x8c\x23\x5e\x96\x91\x97\xb4\x0d\x21\xb8\x40\xdc\xc6\xfb\xd3\xd5\xea\x66\xfd\xf3\x7c\xf1\xeb\x7c\xed\xf6\xfc\x2a\xbe\x3a\x56\x46\x23\x6b\xb3\x68\xaf\x74\xad\xd6\xee\x7f\x7b\xd5\x55\x85\x6e\xce\xfb\x3e\xfc\xa0\xb7\xb1\xa0\xf5\x93\x27\xd2\xff\x6e\x04\x2f\x9d\xf3\x4f\x00\x00\x00\xff\xff\x0d\x84\xc7\xe3\x0b\x07\x00\x00")

func clientScriptsDaemonUpShBytes() ([]byte, error) {
	return bindataRead(
		_clientScriptsDaemonUpSh,
		"client/scripts/daemon-up.sh",
	)
}

func clientScriptsDaemonUpSh() (*asset, error) {
	bytes, err := clientScriptsDaemonUpShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "client/scripts/daemon-up.sh", size: 1803, mode: os.FileMode(493), modTime: time.Unix(1528351521, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _clientScriptsDockerSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\x4f\x6f\xd3\x40\x10\xc5\xef\xfb\x29\x1e\x4d\xd4\x82\x84\x63\x9a\x23\x55\x91\x4a\x92\x13\x95\x22\x25\xed\x19\x6d\xed\xb1\xbd\xaa\xbd\xeb\xce\x8c\x5b\x0c\xe2\xbb\x23\xff\x85\x82\x80\xe6\x14\xef\xbc\x79\xf3\xf6\xb7\xb3\x78\x15\xdf\x39\x1f\x4b\x61\xcc\x02\x1f\x43\x50\x51\xb6\xb5\xc0\xa2\xb2\x49\xe1\x3c\x21\x0b\x8c\x34\x24\xf7\xc4\x2b\x63\x84\x14\x11\x19\xb3\xdd\x6f\x3e\xed\x0e\x9f\x8f\xfb\xdb\xc3\x66\x77\x99\x93\xae\x46\x49\x12\xaa\xa9\xb8\xdd\x1d\x6f\x2e\xcf\x62\xad\xea\x38\x27\x8d\x46\x81\x14\x67\xdd\xa8\xe3\xbd\xab\xe1\xbc\xa8\x2d\x4b\xab\x2e\x78\xb8\x0c\xdb\x5e\x02\x27\xb0\x25\x93\x4d\xdb\x49\x41\xe9\xca\xb8\x0c\x85\x95\x62\xcc\x82\xf5\x87\x38\xa5\xc7\xd8\x37\x65\x79\x01\x2d\xc8\x1b\x00\xa0\x2f\x4e\xf1\xce\x64\xee\xc2\x98\x8c\x34\x29\x32\x57\xd2\xeb\x37\xf8\xd6\x57\x17\xb8\xe2\x5c\xde\x8f\xff\x81\xe5\x39\x24\x34\x9c\x10\x6e\x0f\xd7\x3f\x4f\xd7\x48\x49\xd4\xf9\x21\x57\xe7\xb0\xea\x8b\x53\x82\xa4\xe1\xf2\x2f\xf3\xbb\x9f\x34\x69\x18\x34\x51\x26\xc7\xeb\x6e\x4a\x14\x70\xb2\x5c\x9f\x0c\x11\xcb\xc9\xe7\x29\x27\xfd\x9f\x4f\xaf\x89\xf6\x7d\x3b\x96\xe7\xa3\x83\xd0\xac\x62\xd2\x86\x3d\x86\x4a\x77\xef\xef\x1d\xde\xab\xca\x7e\x0d\x1e\xbb\xcd\x71\x40\xe8\x13\x12\x30\x3d\x34\x8e\x09\x49\x23\x1a\xaa\x89\x6d\x07\x36\x67\xaa\x11\x3d\x4c\x6d\x31\x69\x12\x4b\x2b\x4a\x55\xc4\x54\x92\x15\xfa\x25\x5b\x9f\xab\x6d\x66\x03\x44\xed\xf8\x28\x66\x4e\xb6\xc0\x0d\xb7\xd0\x80\x34\x3c\xf9\x32\xd8\x14\x8d\x38\x9f\x0f\x58\x02\xf7\xd7\x7a\x3b\x2a\xef\x28\x0b\x4c\x60\x92\xc0\xda\x89\x34\x4c\xd6\x53\xcb\xcc\x7f\x7e\x53\x2c\x9f\xad\xe0\xfc\xd9\x2d\xdd\xef\x1c\x8b\x67\xd5\x3f\x11\xda\x5a\xa3\x0e\x73\x53\xa7\x56\x09\xa7\xa7\xf3\x49\x34\x6f\x60\x1f\x63\xee\x78\x51\x8c\x7f\x06\xc8\x9c\xc9\x9c\x31\x3d\x4b\x21\x7e\x74\x09\x4d\x9b\x2d\x6a\x59\xcd\x8f\x00\x00\x00\xff\xff\x6e\x88\x52\x3d\x99\x03\x00\x00")

func clientScriptsDockerShBytes() ([]byte, error) {
	return bindataRead(
		_clientScriptsDockerSh,
		"client/scripts/docker.sh",
	)
}

func clientScriptsDockerSh() (*asset, error) {
	bytes, err := clientScriptsDockerShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "client/scripts/docker.sh", size: 921, mode: os.FileMode(493), modTime: time.Unix(1528351521, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _clientScriptsKeygenSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x91\x41\x6b\xdc\x30\x10\x46\xef\xf3\x2b\xbe\xb0\x81\x5c\x2a\xef\xbd\xa1\x0b\x6d\x13\xda\x3d\x74\x13\x68\x7a\x2a\xc5\x68\xad\x71\x34\xd8\x91\x8c\x66\xdc\xad\x2f\xfd\xed\xc5\xeb\xa5\x24\x64\x2f\xd1\x55\x6f\x78\x4f\x9a\xd5\xc5\x7a\x2f\x69\xad\x91\x68\x85\xfb\x92\xc3\xd8\xb0\xc2\x63\x18\xf7\xbd\x34\x6e\x28\xf2\xdb\x1b\xa3\xe3\xc9\x0d\x5e\x0a\x7c\x0a\xc8\xa3\x0d\xa3\x29\x2c\xf2\x89\x9b\xef\x2b\x22\x65\x83\x63\xa2\xed\x4d\x7d\x73\xfb\xfd\x61\xbb\xfb\xf8\xb0\xbd\xdb\x7d\xb8\xfc\x7a\xf7\xed\x76\x5d\xa9\xc6\xb5\x84\xba\xa8\xaf\x25\x71\x31\xf1\x75\xe0\xa1\xcf\x13\xdd\xff\xf8\x54\xbf\x71\xa6\x1a\xc6\xfd\x9c\xfc\x39\x72\xd3\x41\x5a\x04\x56\x93\xe4\x4d\x72\x42\x2b\x3d\xc3\xf7\x85\x7d\x98\xc0\x7f\x44\x4d\xdf\x93\xb4\xf8\x09\xd7\xe2\xf2\xa5\x09\xbf\xae\xe7\x87\x24\x02\x80\x23\x73\x71\xa4\x5e\x37\xbd\x20\xe7\xb3\xc2\xb6\x7d\xf6\x01\x08\x99\x35\xd9\x22\x7c\x87\x27\xdf\x31\xc4\xaa\xff\xb8\x6a\x74\x1d\x4f\x8f\x9c\xe0\xa6\x73\x25\x9b\x73\xd6\xe3\x78\x2b\xd7\xc4\xbd\x32\x2d\xda\x2f\x9c\xb8\x9c\xd6\x82\x83\x58\x44\xca\x18\xbc\xea\x21\x97\xb0\x08\x9f\xcb\x5e\x9b\x9c\xa1\xa8\x87\xdb\xe1\xea\x8a\x5a\x21\x3a\xe1\xda\xf8\x84\x47\xb1\x38\xee\xab\x26\x3f\x61\xb3\xc1\xdf\x65\x09\x5d\xca\x87\x54\xc7\xac\xa6\x44\x8d\xb7\xb3\xa9\xff\x02\x00\x00\xff\xff\x37\x00\x91\x4b\x4d\x02\x00\x00")

func clientScriptsKeygenShBytes() ([]byte, error) {
	return bindataRead(
		_clientScriptsKeygenSh,
		"client/scripts/keygen.sh",
	)
}

func clientScriptsKeygenSh() (*asset, error) {
	bytes, err := clientScriptsKeygenShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "client/scripts/keygen.sh", size: 589, mode: os.FileMode(493), modTime: time.Unix(1528351521, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _clientScriptsTokenSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\xcf\x41\x4b\xfb\x40\x10\x05\xf0\xfb\x7e\x8a\xf7\xa7\xfd\xd3\x53\xb2\xf7\x42\x0e\x45\x82\x29\xd6\x46\x8c\xe2\x45\x08\xdb\x64\x6c\x42\xcc\x6c\x9c\xd9\x55\xfc\xf6\x12\x4d\xf5\x38\xf3\x9b\x79\xf0\x56\xff\xec\xa9\x67\xab\x9d\x31\x4a\x01\x09\x19\xb3\xc2\xa3\x92\xc0\xc9\x39\x8e\xc4\x21\x35\xf7\xf9\x21\xdf\x55\x79\xf6\x5f\x67\xbc\x26\x26\x71\x81\xe0\xd0\x3a\x1a\x3d\x23\xf8\x81\x18\x51\x7b\x3e\xe3\xea\xb0\xc7\x8b\x17\xec\xee\xf6\x10\x7a\x8b\xa4\x41\x53\xa3\xb1\xf5\x68\x7d\x33\x90\x40\x22\x23\x49\x64\xc4\xb3\x01\x80\xe4\x1d\xeb\xa2\xbc\xcd\xb7\xd6\x4d\x93\xed\xbc\x86\x0b\x10\xaa\xaa\xa8\x6f\x8e\xe5\xd3\xb1\x2e\xca\xea\xa1\xca\x36\xbf\x37\x36\x55\xed\xec\xc0\xfe\x83\xeb\x79\xd6\xcd\xdf\xd7\x9c\x96\x7d\x67\x5e\x76\x09\x71\x90\xcf\xc9\xf7\x1c\xb2\x9e\x49\x42\xef\xda\xc5\xe2\xa9\x79\x75\x91\x9b\x6e\x72\xad\x5d\x6c\xbb\x5e\x1a\xff\x34\x33\x5f\x01\x00\x00\xff\xff\xb8\x50\x60\x94\x24\x01\x00\x00")

func clientScriptsTokenShBytes() ([]byte, error) {
	return bindataRead(
		_clientScriptsTokenSh,
		"client/scripts/token.sh",
	)
}

func clientScriptsTokenSh() (*asset, error) {
	bytes, err := clientScriptsTokenShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "client/scripts/token.sh", size: 292, mode: os.FileMode(493), modTime: time.Unix(1528351521, 0)}
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
	"client/scripts/daemon-down.sh": clientScriptsDaemonDownSh,
	"client/scripts/daemon-up.sh": clientScriptsDaemonUpSh,
	"client/scripts/docker.sh": clientScriptsDockerSh,
	"client/scripts/keygen.sh": clientScriptsKeygenSh,
	"client/scripts/token.sh": clientScriptsTokenSh,
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
	"client": &bintree{nil, map[string]*bintree{
		"scripts": &bintree{nil, map[string]*bintree{
			"daemon-down.sh": &bintree{clientScriptsDaemonDownSh, map[string]*bintree{}},
			"daemon-up.sh": &bintree{clientScriptsDaemonUpSh, map[string]*bintree{}},
			"docker.sh": &bintree{clientScriptsDockerSh, map[string]*bintree{}},
			"keygen.sh": &bintree{clientScriptsKeygenSh, map[string]*bintree{}},
			"token.sh": &bintree{clientScriptsTokenSh, map[string]*bintree{}},
		}},
	}},
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

