package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindata_read(data []byte, name string) ([]byte, error) {
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

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _pr_review_js = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5a\x5b\x73\xd3\xc8\x12\x7e\xcf\xaf\x98\x40\x15\xb6\xc1\x28\x09\x8f\x71\xe5\x5c\x0a\x0e\x55\xe7\x02\x49\x11\xa8\xf3\x60\x5c\x46\x91\xc6\xb1\x88\xad\xf1\xd1\x48\xf1\xf1\x82\xff\xfb\x76\xcf\x7d\x46\xe3\x4b\x48\x76\x97\xdd\x45\x0f\xc1\xd6\xf4\x7d\xba\xbf\x9e\x69\x33\x69\xca\xac\x2e\x58\x49\xae\x69\xfd\x8e\xde\x16\x74\x49\x2b\xfe\x9a\x55\xaf\x8b\x19\xed\x4e\xe0\x0f\xef\x93\x45\x33\x9b\xbd\xa3\xff\x6b\x28\xaf\x3f\x70\x5a\xc1\x8b\x8a\x4e\x68\x55\xd1\xdc\x70\xf4\xc8\x97\x03\x42\x6e\xd3\x8a\x54\xea\x15\x0a\xe0\xe4\x8c\x7c\x59\x0f\xd4\x0a\x4a\x33\x0c\x6a\x05\x96\xc6\x09\x4d\xb3\x69\xb7\x2d\xb3\x4f\x8c\x71\xda\x12\x2d\x5c\xaa\x33\xbc\x6a\xd5\x27\xd7\x34\xc4\x57\x3c\xc4\x6f\x23\x50\x1f\x7b\xfb\xf5\x2b\x19\x8e\x06\x9b\xd9\x92\x45\xc3\xa7\x5d\x63\x85\xa4\x5c\x8b\x7f\xf1\xef\xc1\x7e\x16\xb9\x61\xc2\x40\xb4\x5d\x17\xda\x12\xfc\x53\xa6\x73\x6b\x97\xe0\x2e\x26\xc4\x58\xc0\x13\xa0\xb8\xae\xa7\xe4\xec\xec\x8c\x1c\x5b\x8f\x51\xc3\x22\xad\x6b\x5a\x95\x20\x7f\x0c\x92\xca\xbc\x9b\xb1\x72\x52\x5c\x27\xea\x7d\xcb\xba\x97\xb0\x6c\x25\x78\x32\xb8\x0a\x17\x92\x68\xfe\x81\x47\xd8\x40\x5e\xa0\xa9\x1e\xa5\xb1\x52\x59\xae\x89\xe7\x69\x9d\x4d\xad\x5d\x5a\x49\xc2\x17\xb3\xa2\xee\x76\xfa\x9d\x9e\x6b\x9b\x5a\x76\x4d\x23\xe4\xe8\x08\xbc\xe1\x0c\x82\x34\x63\xd7\x9a\xa4\x4f\xbc\xb0\xf5\x06\x0e\x43\x45\xeb\x06\x82\x21\x54\x63\x6a\x5e\xa4\xf5\x74\x37\x9f\xda\x54\xad\x93\xe8\x0d\x22\x13\x26\xf3\x99\x64\x69\x59\x93\x2b\x4a\xea\x29\x25\x1c\x98\x49\xca\xc5\x67\xac\x19\x50\x2a\x8a\x86\x70\x5a\xe6\xb4\x32\x82\x70\x07\x65\x0c\x9e\x3c\x21\x5d\x13\xba\x04\x82\x41\xff\x7f\x3e\xe9\x06\xf5\xd6\x13\xbb\xfb\xfc\x04\xb3\xc0\x12\xab\x8d\xff\x0b\x39\xe9\xf9\xa1\x51\x9e\x8e\x93\x65\x51\x4f\x59\x53\x5b\x05\xad\x4a\x76\x3d\x3d\x68\x79\x8c\x66\x1e\x46\xa2\x2f\x15\x68\xd6\xb5\x26\xf7\x32\x5a\x72\x45\x32\x00\xf7\x3f\x9d\x55\x34\xcd\x57\x32\x98\x45\x79\x2d\x52\xa1\x28\x81\x83\x53\xb1\xe9\x5d\x47\x97\x06\x03\xf3\x6a\x9c\xdc\xd0\x15\xef\x7a\x38\xd3\x53\xab\xbe\xf1\xa1\x1e\x1b\xb3\x63\xdf\x1f\x6b\x78\xc8\x62\xbd\x74\x7c\x15\xf5\x9b\x96\x39\x9b\x9b\x7c\x38\xb3\x52\x86\x6f\x20\xb5\x92\xc9\x8c\xb1\xaa\x2b\x3e\x4a\xd2\x6e\x8f\x3c\x25\x61\xdd\xf6\x74\x51\x7b\xde\x0c\x7d\xe1\x23\x47\x7a\x7c\xdd\x81\xad\xad\x84\x12\xbe\xda\xc9\xbe\xee\x43\xce\x16\x5c\x46\x4f\xe5\x8f\x27\x68\x70\x00\xae\xbb\xad\x02\xdf\x5e\x4e\xd3\x37\xe9\x42\x42\x9d\xc5\xff\x89\x59\x31\xe0\xbf\x0f\x26\x5a\xb6\x10\xfb\x24\xa0\x24\x7c\x9a\x6a\x98\x35\x36\x5a\x26\xdf\x40\x27\xcd\xcf\x17\xb4\xa4\x79\x97\xde\xd2\xb2\x7e\x95\xd6\xa9\x35\x34\xa8\x05\xd0\x63\x88\x12\x5c\x1b\xab\xea\x4d\xb0\x7e\x10\x66\x8a\x52\xc4\x47\x3b\x89\xe9\x02\x91\xb8\xb0\x62\x44\xa8\xac\xaa\xa4\x6c\xe6\x57\xaa\xc8\x62\x9d\xf1\x2e\x1d\xf7\xcb\xba\x37\x88\x05\x38\xb6\x15\xc2\x4a\xcc\xff\x68\xa5\x44\x8a\x40\x60\x32\xe5\x3c\xbd\xa6\x20\xb2\xf3\xf7\x25\xe5\x0c\x90\x6c\xc9\xaa\x9b\x43\xf2\x96\x2d\xc9\xe7\x06\x41\xac\x00\xa0\x4b\xb3\x1b\x02\x39\x45\x96\x29\x7c\x43\x10\x7c\x5f\xa5\xb7\x05\x40\x1e\x83\x9a\xe7\x5c\xac\x31\xc0\xbf\x4a\xbc\x92\x9a\xc9\x8a\x35\x15\xc9\x58\x4e\x93\x8f\xe5\xc7\xb2\x23\x13\x55\x2b\x7c\x06\x1a\x1f\x3f\x36\xc8\xca\x05\x81\xdb\xdc\x3d\xf3\xf7\x38\x13\xf8\xa2\xff\xd6\x21\xcf\x0c\x09\x7c\xec\x18\x03\xf6\x3d\x3d\xf8\xf2\x9e\x13\x14\x28\xa0\xdf\x17\xb6\x36\x60\xea\x92\x1b\x02\x83\xab\xfe\xea\x63\xe3\x3a\x04\x68\x3e\x87\xf0\xf1\x68\x84\x9e\x93\x34\xcb\xe8\xa2\x3e\x25\x9f\x16\xd5\x58\xc5\xf5\xfc\xdf\x9f\x36\x11\xe7\xb9\x71\xda\x63\xc1\x85\xa1\x6e\x08\xa3\x4d\xec\xd9\x34\x2d\xe1\x4b\x54\x82\x5a\x1b\xb2\x59\x3e\x36\x82\xc8\xb0\xa4\xcb\xf1\x4e\xb9\x8b\x0a\x50\x9e\xf0\x3a\xad\x1b\xee\x09\x95\xaf\x3e\xd9\xbd\xcf\x00\x86\x6b\xfa\x4f\xce\x1b\x38\x4a\xcc\xe7\x50\x51\xad\xba\xea\x6b\xe1\x2a\xf0\x92\xe5\x52\x48\xd2\x0d\x64\x43\x45\x4f\x01\xe3\x11\x50\x74\x53\xe9\x00\x48\xe4\x80\xf6\x9d\xfe\x76\xbe\xa6\x9a\x19\x96\x0b\xb7\xc1\x4b\x37\x4e\x45\x6e\xc4\xab\xee\x33\x2b\x4a\x38\xda\x10\x38\xdb\x18\xa5\xda\xfd\xce\x81\xea\x5d\xf8\x0f\x4f\x6f\x29\x6a\x6e\xf9\x8b\xf9\x76\x8a\x75\x07\x86\x2a\xc6\xf1\xa4\x62\x73\x10\xfa\xaf\xcb\xf3\xb7\x09\xaf\x21\xba\xd7\xc5\x64\x15\xa8\x96\x08\x8f\x61\xdd\x2e\x5a\x94\x41\x5b\x98\x85\x1b\x94\xb4\x01\x65\x3f\x2c\x72\x88\xfd\xc3\xc1\xec\x3d\x50\xd6\xda\x7b\x59\xa3\xb6\x19\x4b\xf3\x5d\x4e\x1b\x66\x48\x6a\x03\x44\xfb\xb0\xc7\xb6\x23\x0e\xd4\xf6\x1a\xe4\xea\xb0\xb7\xa0\xba\x5a\xb5\x3a\x21\xac\x8a\xcd\x58\xa4\x70\x30\xea\x7a\x8e\xc9\x4d\x85\xd3\x27\x9c\x22\xbb\x08\x56\x6b\x5f\x48\xa0\xc4\x11\x13\xb8\x18\x17\xa4\x23\x8f\x27\xf4\x19\x9c\xcb\xb6\x63\x24\xf6\x9a\xcd\xfd\xfb\xf0\xcc\x76\x70\xff\x4e\xb6\x4f\xbf\x97\x67\x1a\xd1\xed\xeb\xaa\xa1\x0a\x4c\xed\x7d\xeb\xbe\x8d\xd5\x0d\xc7\xbd\x3b\xe7\x25\xb6\x4d\x36\x11\x17\x00\x19\x42\x60\xa6\xa4\x91\xd5\xb1\x77\xf7\xfb\xe3\x36\x3f\xb8\x42\xbd\xa1\x15\x2c\x43\xd8\xb9\x0a\x11\x9e\x29\x5c\x97\xdd\x1d\xd9\xcb\xe3\xe0\xc8\xbb\xf9\xd4\x6c\x56\xfe\x0a\xba\x9a\x12\x2f\x1a\x81\xd0\x80\xb0\x47\x4e\xa5\x91\x81\x1b\x77\xec\x4e\x3f\xda\xd3\x6f\xdc\x9e\x5a\x9b\x64\xdb\x53\x91\x7b\x1d\xa9\xc0\x3d\x55\x6a\x35\x60\xcb\xdb\xfb\x5b\xbc\xdc\xbb\xa4\xf2\xb5\xdf\xb1\xae\x58\xbe\xf2\x88\x32\xa9\x39\xc1\x85\x04\x8c\x86\x6b\xa0\xa6\xcd\xe6\xb9\x3b\x31\x12\x14\x6a\x02\x02\x85\xe3\x8d\x40\x74\x2a\x39\x78\xab\x5e\x99\x99\x81\xb3\x5d\xbd\x70\x1a\x04\x9a\x5c\x2c\x77\x6e\xd2\x73\xb8\x34\x68\x49\xbc\xb9\x92\x61\xed\x9e\x1c\xf7\xb4\x25\x24\x9c\xc5\x54\x35\xc8\xd5\x88\x8c\x5f\xb5\x57\x50\x1d\xf6\x8e\x6e\x39\x40\x77\x74\x3c\x71\x78\x98\xe5\xf3\xf6\xf8\xc1\xe0\x08\x5f\x16\xd8\x91\x80\x7d\x78\x3c\x4a\x6a\xf6\x1f\x06\x19\xf3\x32\x85\x06\xe6\x8d\x3b\x32\x78\x43\x3a\xec\xa6\x73\xea\xa8\x10\x9b\x56\x57\xb8\x01\x6e\x07\x87\xad\xde\xda\xb3\xf5\x83\xe1\x55\xec\xbe\xe9\x6e\x83\xf5\xb5\x85\x1d\xc8\xe9\xb7\x5a\xd2\x20\x60\xcb\xe9\x8c\xd6\x34\x40\x1e\x9b\x6a\xa3\x90\xde\x54\xc7\x16\x37\x76\x97\x5a\xc4\xf6\x45\xf5\x8a\xd6\x69\x31\x6b\x9f\xb4\xd4\x7b\xd0\xd8\xb2\x7e\x67\x93\x0c\x52\xd0\x3e\x1e\x0e\x1a\xdd\x16\xf3\x48\x87\x37\x70\xd5\xe1\x58\xeb\x76\x19\xa1\x2d\x0a\x6a\x88\x59\x4e\xea\x87\x76\x46\xb1\xba\xc8\xfb\xe4\x91\xbd\x75\x2d\x70\x23\xf2\xc3\x47\x2d\xe6\x35\xa1\x33\x48\xaf\x1d\x3e\xb4\x56\x09\x89\xb8\x15\xa1\x6a\xc1\x7a\x5c\x84\x83\xea\x1e\xf7\xbd\x11\xde\xb7\xc5\x47\x7b\xf7\x69\xc7\xc5\x4f\x23\xff\xf4\xe8\x92\xe9\x4f\x57\x10\xaf\x9b\x41\x50\xb4\x70\x15\xf5\xaa\x16\x53\x0a\xcb\xfd\x64\x14\xa6\x0d\x66\x29\xdc\x2f\x75\x71\x0d\x83\xda\xb8\x67\xb9\x6f\x2d\xf8\x78\xc9\xdf\xa1\xe8\x5b\x9c\xda\xcd\x17\x2d\x37\xe5\xe3\x38\x6a\x10\x3b\x76\x01\x82\xa2\xdc\x76\x6c\x73\x9f\xe8\xec\x5b\xda\xb0\x6d\x64\xae\x9f\x75\xe4\xed\xba\xed\x98\x8f\x64\x72\x27\x47\xc2\x0b\x79\xd4\x8a\xaf\x8b\xd1\x65\xdf\x78\x1d\x51\xf5\x30\xc0\x17\xdb\x86\x7d\x4f\xf9\xfe\xf3\xcd\x98\x89\xcf\x2e\xe4\xd8\x17\x3b\x76\xa0\xc7\x1e\xf8\xf1\xd0\x08\xb2\x1d\x43\x22\x28\xe2\x02\x84\xfa\xbe\x01\x48\x7c\xca\x8d\xb0\x72\xf4\x34\x40\x98\x8a\xce\xd9\x2d\xbd\x03\xc8\x7c\xb7\x28\x12\x32\x46\x0f\x0f\xaa\xa4\x7e\x9d\x0a\xba\xe7\x5e\x3d\x3d\x0a\xf6\x4a\x8e\x15\x37\xec\x15\xfe\x4e\x16\xc7\xcc\xef\x7a\xd7\xa2\xa0\x13\xdd\xb3\x38\xd2\xb4\x49\x5f\xec\x40\xd4\x17\x16\x51\xe3\x7a\x62\x98\x74\xb7\x64\xfa\xa5\x00\xf9\x81\xc1\x40\x23\x00\x03\x98\x74\xb2\x4a\xcc\xe5\xd2\x1b\xfa\x0f\xbc\xa1\xe1\xf4\x4d\x5e\xf5\x4e\x09\x9e\x09\xdd\xcb\xf4\xe9\x66\x5c\x5f\x5b\x2d\xed\x1f\x9a\x8c\x70\xc7\xc5\xa8\x69\x72\xe4\xfd\xfb\xbb\xb7\x04\x23\xaf\xff\x4a\xab\xf4\x4f\x3d\x68\x95\x33\xe6\xb2\xcf\x37\xcf\xb3\xec\xb3\xef\x64\xab\xa5\x74\xe7\x8c\x2b\xae\x63\xe3\xb4\xcb\x3e\xb1\xa3\x51\x74\x02\xe6\xb3\x04\xaf\x36\x5c\x51\x82\x01\x92\x7d\xbe\xfd\xf2\x75\x8f\xb3\xcb\x8f\x7b\x9b\x7e\xfe\xa4\xf7\x36\xfd\x5f\x20\x64\x0a\x83\xac\xa3\x23\x72\x51\x31\xdc\x72\x32\x65\xec\x86\xe4\x00\x17\x07\x98\x9b\x62\xfc\xf5\x7e\xb5\xa0\x22\x11\x3b\x2e\xac\x76\xb0\x95\xdb\xf1\x58\x2a\x0b\x52\x90\x31\x01\xa0\x1d\x99\xb6\xdb\x7e\xc1\x07\xe5\x72\x83\x23\xba\xc4\xfc\x6e\xac\xa6\x6e\x5b\x94\xc9\x4c\xc8\x03\x0a\x39\xfd\x73\xed\x6d\x59\xd3\x1e\x25\x6e\x33\x67\x4f\xd7\xf9\xaa\xcc\xa6\x15\x2b\x8b\x9f\x68\xdb\xff\xf6\x6f\x6b\x18\xfd\x9f\x03\x00\x00\xff\xff\xe9\x47\xbf\x78\xb8\x27\x00\x00")

func pr_review_js_bytes() ([]byte, error) {
	return bindata_read(
		_pr_review_js,
		"pr_review.js",
	)
}

func pr_review_js() (*asset, error) {
	bytes, err := pr_review_js_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "pr_review.js", size: 10168, mode: os.FileMode(420), modTime: time.Unix(1441698677, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"pr_review.js": pr_review_js,
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

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"pr_review.js": &_bintree_t{pr_review_js, map[string]*_bintree_t{
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

