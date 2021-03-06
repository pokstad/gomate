// Code generated by go-bindata.
// sources:
// assets/remarkdown.css
// DO NOT EDIT!

package main

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

var _assetsRemarkdownCss = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x5f\x73\xdb\x36\x12\x7f\xe7\xa7\xd8\x2a\xd3\x19\x5b\x15\x24\x5a\xb6\x73\x2d\x33\xf5\xdc\x4d\x93\xce\xf4\xc1\x77\x33\xd7\x9b\x7b\x49\x33\x17\x88\x5c\x89\x38\x81\x00\x0f\x00\x2d\xbb\x39\xe7\xb3\xdf\x80\x20\xcd\x7f\x20\x2d\x27\x3d\xa7\x0f\x55\x34\x13\x0a\x58\x2c\x7e\xfb\x07\x8b\xdd\xa5\x57\xf3\xaf\xe0\xef\x78\x4d\xd5\x3e\x91\x07\x01\xe7\xcb\x70\x19\xc2\xc9\xf5\x4f\xff\x38\x85\xd4\x98\x5c\x47\xab\xd5\xf6\x46\xc7\xe9\x72\xc7\x4c\x5a\x6c\x96\x4c\xae\x14\x66\x15\xf9\x0a\xe6\xab\x60\x35\x0f\x00\x7e\x12\x31\x2f\x12\x4c\x22\xa0\x9c\x83\x36\x77\x1c\x75\x00\xf0\x1a\xb7\xb4\xe0\x46\x47\x90\x0a\xa2\x50\xa3\xb1\x0f\x29\xd5\x29\x14\x9c\xa4\x77\x79\x8a\x02\x24\x27\x09\xc6\x2c\xa3\x1c\xfe\x53\x48\x83\x64\x67\x00\x33\xa2\x0d\x55\xa0\x8d\x92\x62\xe7\x9e\x29\xd9\x28\x1a\xef\xd1\x40\x2c\x13\x24\x86\xc5\x7b\xc8\x15\x12\x26\x12\x14\x06\x52\x55\x92\x05\xf3\x55\xb0\x6c\x20\xc2\x87\x00\x80\x33\x81\x24\x45\xb6\x4b\x4d\x04\x67\xcb\xcb\x57\x01\xc0\x56\x0a\x43\xb6\x34\x63\xfc\x2e\x82\x4c\x0a\xa9\x73\x1a\xe3\xa2\x79\x7c\x15\xdc\x07\x1d\x4e\xb9\xc2\x45\x67\xc0\xc2\xe8\x8e\xec\x37\x49\x77\x40\xd3\x2c\x2f\x21\x74\xb6\x63\x22\x45\xc5\xcc\x60\x87\xf4\xac\xa1\xd5\xec\x57\x6c\x51\x56\xa3\x87\x4a\x88\xd6\x78\x46\xd5\x8e\x09\x62\x64\x1e\xc1\x7a\xb9\xbe\xc4\xac\x35\xba\x91\xc6\xc8\xac\x14\xda\x8e\xf7\xf7\x5b\x3f\xf3\x7e\xe7\x5d\xed\xa4\x17\xbd\xdf\x97\xbd\xdf\x2f\x3f\x13\x5f\x05\xe3\x58\x78\x67\x51\xb4\xc1\xad\x54\x58\x6e\x1b\x4b\x61\x50\x98\x08\x66\x2f\x60\xe6\xd1\xdd\x18\xb1\x97\xfa\x7c\x94\xda\x4b\x7e\x31\x4e\xee\xa5\xbf\x9c\xa0\xf7\x2e\x78\x39\xb5\xc0\xb3\x62\x99\x9e\x91\x42\x24\xa8\xec\x59\x82\xf4\x6c\xd1\x9d\x5c\xb7\x27\x9d\x57\x25\x4c\xe7\x9c\xde\x45\x60\xe8\x86\xa3\x35\x82\xbc\x41\xb5\xe5\xf2\x10\x41\xca\x92\x04\x85\x1d\xcb\xa5\x66\x86\x49\x11\x81\x42\x4e\x0d\xbb\x29\x29\x73\x9a\x24\x4c\xec\xa6\xed\xd5\x87\x54\x4b\x34\x09\xcd\x2f\xb6\x90\x62\x70\xde\x87\xec\xe9\xd6\xa0\x7a\x84\x7b\x49\x53\x32\x6f\x24\xa3\x1b\x2d\x79\x61\x4a\xc9\x6a\x89\x42\xfb\x83\xe3\xd6\x54\x8f\xad\xe8\xe4\x3c\xf6\xc0\x12\x93\x46\x70\x16\x86\x5f\x8f\xe9\xee\x20\x55\x42\x36\x0a\xe9\x3e\x82\xf2\x3f\x42\x39\xb7\x13\x71\xa1\xb4\x54\x11\x24\x2e\xfc\x1e\x29\x59\xcf\x13\xbe\xff\x42\x1f\x8f\xeb\x8d\x2b\xb9\x81\x4b\xbe\xd0\x67\x78\xb6\x5c\xc0\xff\xf4\x30\xb4\x65\xbb\xa2\x72\x50\x47\x5f\x11\x42\xd8\x5a\x90\x3a\xf9\x07\x6e\xd3\xbf\xea\xda\xfb\x0e\xf8\x58\x77\x54\x09\xaa\xfa\x00\x58\x7d\x72\xeb\x38\x23\x17\xd4\x32\x55\x24\x46\x61\x95\x5f\xed\x6f\xf0\xd6\x10\xca\xd9\x4e\x44\xe0\x66\x86\xa1\x46\x8d\x84\x9a\x39\x94\xff\x3c\xf6\x56\x75\x6a\x30\xba\x76\x54\xf9\x05\xf7\x29\x2e\x4e\xab\xa7\x8b\x38\x6d\xc5\x97\xf2\xec\xf5\x18\x70\xd6\xe5\xe1\x2c\x18\x7a\xac\x37\x58\xbb\x2c\x38\xc9\x79\xa1\x6b\x06\x9c\x69\x43\xca\x5c\x68\x24\xc2\xb4\xe8\xaf\x80\xb3\x8e\xac\x5b\x2e\xa9\x89\xca\x10\xd1\x0e\x07\x0e\x7f\x85\xc4\xc5\x0f\x72\xd1\x19\x54\x95\xed\xdd\x60\xa3\xb2\x6f\x86\x8a\x2e\xb8\xcb\xac\x9e\x80\xb7\xa6\xff\xff\xe3\x1d\x3a\xc6\xb1\x38\x9f\x07\x9f\xc7\xf9\xe4\x6f\xe0\x7c\xf2\x13\x9d\xaf\x5e\x38\x54\x8e\x05\x5d\xd8\xa3\xe9\x32\xf0\x08\x54\x96\x10\xc9\x7d\x0c\xac\xd6\xaa\x93\xe6\x56\x30\x11\x2b\xcc\x4a\x81\xa7\x57\x4d\xe9\xfa\x78\xed\x66\x4c\x90\xca\x32\xeb\x9e\xbe\x2b\x44\x27\x0e\xc6\x29\xcc\x96\x43\x7f\x96\x9c\xfc\x8a\x4a\x8e\xab\x62\x9c\xfe\x59\x64\x98\x85\x7e\xd0\x94\xe7\x29\xfd\x6c\x03\x76\x38\x3d\xdd\x92\x83\xe5\xcf\x6c\xd2\x05\x70\x79\x40\xe5\x30\x78\xed\x0b\x74\xe4\x36\x78\xeb\x25\xf5\x65\x09\xef\x86\xfa\xa7\x44\xa7\xf2\x50\x28\x0e\xf4\x6d\xaa\x70\xfb\xee\xe8\x4d\x7c\x4b\xbd\x9b\x9e\xcc\x80\x1a\xa3\x4e\x2c\xcd\x29\xcc\x4e\x67\x13\xd9\x5b\x4f\x0e\xcc\x7a\x05\x65\x59\x10\xb7\xca\xa2\xda\x51\x54\x46\xf9\xa0\x2c\xaa\x87\x07\x4c\xbd\xb9\x72\x39\xee\x13\xc0\x73\x47\x63\xe6\x72\x32\x1d\x5b\x4d\x8d\x30\xf4\x50\xf9\xd8\xff\x6b\x68\x3e\x27\xa6\x1f\x65\x3d\xe7\x45\xea\x81\x5a\xf5\x10\x5a\x40\x26\xb8\x4f\x50\x7b\xa1\x7b\xb0\xc7\x32\x41\x3f\x72\x37\xe3\xe3\xf3\xde\x93\x46\x2a\x9c\x60\xd5\xcc\x36\xec\x1e\x2a\x2e\xef\xa5\x98\x77\x72\xca\xa3\x12\x53\xb0\xb5\x5b\x55\x79\x5f\x3c\xc6\xd0\xc5\x02\x6f\x28\x08\xed\x60\xcf\x2a\xb9\x6a\x5a\x38\x7e\x4b\x74\x29\x3c\xc5\x97\x23\xe0\x09\x3e\xc6\xe3\x81\x64\x8a\x09\xd9\x16\x9c\x1f\xc1\xa9\x45\xe7\x51\xfd\x86\xcb\x78\xef\xa9\xdf\x06\x29\xf9\x63\x35\xda\x27\xe9\xa7\xef\x55\xef\xdf\xcf\xc6\x8b\x6c\x97\x5b\x54\x1d\x1c\xef\xee\x4f\x56\x6d\x6f\xff\x8f\x1f\x3f\xce\x9e\x58\xbc\x8e\xe1\xf8\x2c\xeb\x74\x10\x7d\x91\xcf\xe7\xa9\xc1\x79\x55\xd9\x1c\xf5\x24\x98\xeb\x87\x04\x73\xfd\x68\x82\xd9\x63\xe4\x73\x8c\xd1\x15\x9d\x7b\xd1\xdf\xdd\x68\xd2\xd5\x2a\x33\xa8\x10\x75\xba\x1e\xdd\xd4\xdb\xab\x94\x94\x19\x24\x65\x17\x36\xb2\xb6\xec\x26\x51\x57\xbf\xd0\x3f\xbe\x7f\x7c\x7f\xf3\xef\xec\x98\xb0\x9c\x20\x27\xbb\x6d\x06\x09\xf2\xa6\x01\x92\x60\x2c\x15\x75\xc7\xc1\x5b\x68\xb4\x56\xf9\xe3\x57\x87\xc0\x1f\xb8\x86\x89\x4d\xd9\x58\x25\xf6\x27\x2a\xd7\x65\x7d\xfa\x05\xef\x1a\x40\x24\x96\x9c\xd3\x5c\xa3\x4d\xca\xdd\xd3\x23\xbb\x29\xb8\x82\xb9\xdb\xae\x49\xf0\xcf\xab\xf8\x33\xd5\xd3\xad\x82\x40\xbb\x71\xd4\x6a\xde\xdf\xa0\x32\x2c\xa6\xbc\x9e\x32\x32\x1f\x4d\x6a\xfb\xdd\xab\x23\xf0\x7e\x53\x61\xae\xfb\xcb\x2e\x4a\x9d\x7b\xd2\x93\xde\xea\x68\xcb\x94\x36\x24\x4e\x19\x4f\xe0\x0a\x4c\x1a\x09\x69\x4e\xde\xea\x58\xe6\xf8\xfd\x4c\xc9\xc3\xec\xdd\x69\x87\xf3\x74\xe7\xfa\x13\xb8\xff\x1e\xba\x9e\xbd\x64\xe2\x98\xbe\x76\x9d\x01\x3e\xb9\x87\xdd\x4b\x98\x9e\xa4\x40\x6b\xe8\x47\x95\x78\xb4\xed\x6b\xcf\x19\xa9\x0b\xff\xfb\x0b\x9d\xf8\x4e\xa8\xcc\xa7\x0e\x07\xaa\xba\x1a\x9b\xdb\xb4\x7b\x7f\x0e\x6e\xc7\xfb\x20\x58\xcd\xe7\x01\xcc\xe1\xe7\xf2\xcd\x2b\x6c\xa5\x6a\xbf\xd5\x4d\x30\x93\x90\xd3\x1d\x6a\x4b\xf3\x1a\x73\x14\x89\x06\x29\xea\x28\x07\x2d\xe1\x63\x6d\x89\x56\x41\xf0\xe7\x0c\x13\x46\xe1\xa4\x75\xbe\xbf\x0b\xc3\xfc\xd6\xf9\x79\x6a\x32\x17\x00\x01\x36\x34\xde\xef\x94\x2c\x44\x42\xaa\xf6\xf1\x8b\x37\x6f\xde\x58\xa4\xf7\x16\x19\xad\xf4\xe5\x66\x7e\x0c\xd7\xe1\x65\x13\x01\x7c\x91\x93\x46\xa9\x55\x4d\xb9\xac\x0a\x4f\x0f\xa7\x29\xbf\x05\x2d\x39\x4b\x1c\xdd\x56\xc6\x85\x5e\x00\x8d\x68\x6c\xa3\x4c\x7b\xa3\x52\x49\xa5\xea\x86\xe8\x6a\x0c\xf7\x41\x20\xe8\x4d\x27\xa3\x0a\x21\x84\x8b\xaa\xf4\x69\x45\xa8\xd2\x8b\xed\x02\x4b\x7f\x05\x9b\x91\x7a\xcb\xcd\xea\x9c\x8a\x61\xcf\xe4\xbe\x51\x28\xbd\xf5\x28\xd4\xad\x9d\xc3\x87\x7e\x2d\x31\x08\x29\x4b\x0b\xd0\xe9\x76\x89\xb7\x34\xcb\xb9\x2f\x2b\x24\xe7\x61\x7e\xdb\x89\xbb\xe5\x70\x3d\xda\xa8\xa5\x34\xd7\x8f\x95\x13\xc1\x35\xdd\x23\x98\x14\xa1\x76\x1e\xc2\xd9\x1e\xc1\x1d\x06\x0d\x5c\xca\x3d\xc4\x52\x72\xe8\xbd\x66\xf7\x57\xa9\xed\xf3\x26\x73\x1a\x33\x73\x17\xc1\xf2\x65\xeb\x4d\xc3\x0b\x6b\x8a\x6f\xc3\x7a\xfb\xbf\x15\x0a\xca\x77\x20\xcd\xeb\xf2\x25\xfc\x85\x6b\x09\x1a\x8d\x61\x62\x07\x52\x60\xe9\xde\xb6\xf8\x05\xe4\x65\x4b\x4b\xf7\xd1\x0c\xde\xb8\x5f\xa3\xe0\x72\x01\xd7\x52\xd0\x58\x2e\x60\xf6\x1a\xff\x4d\xff\x59\xc0\xcf\x54\x68\x3b\x28\x67\x0b\xf8\x41\x0a\x2d\x39\xd5\x8f\xfc\x09\x40\xc9\xbb\x7f\x04\xc9\x41\xd1\xbc\x2d\xd5\xb7\xe1\x3a\x5c\x87\x83\xbf\x33\xf8\x41\x16\x8a\xa1\x5a\xc0\xac\x7a\x82\xbf\xe2\x61\x36\xb5\x63\x29\xe6\x87\x21\x67\x80\xe0\x1e\x02\x9f\x4f\xfd\xa9\xf1\x29\x1f\xf4\x5e\xbd\x1e\xc2\x57\x2c\xcb\xa5\x32\x54\x98\x87\x23\xdb\x5e\xd7\xcf\x2b\x22\x58\xbb\xb6\xf6\xab\xee\x9b\xf9\xef\x2e\xbf\xae\x8d\x58\x05\x21\x1b\x70\xca\x48\xf2\xa2\x6c\x52\x11\xce\xb4\x69\x27\x0e\x1d\x25\x0a\xe9\x54\x78\x1f\xfc\x2f\x00\x00\xff\xff\x6f\xf3\x7d\x94\x92\x22\x00\x00")

func assetsRemarkdownCssBytes() ([]byte, error) {
	return bindataRead(
		_assetsRemarkdownCss,
		"assets/remarkdown.css",
	)
}

func assetsRemarkdownCss() (*asset, error) {
	bytes, err := assetsRemarkdownCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/remarkdown.css", size: 8850, mode: os.FileMode(420), modTime: time.Unix(1520148926, 0)}
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
	"assets/remarkdown.css": assetsRemarkdownCss,
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
	"assets": &bintree{nil, map[string]*bintree{
		"remarkdown.css": &bintree{assetsRemarkdownCss, map[string]*bintree{}},
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

