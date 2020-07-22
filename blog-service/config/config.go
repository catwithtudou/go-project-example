// Code generated for package config by go-bindata DO NOT EDIT. (@generated)
// sources:
// config.yaml
package config

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

// Mode return file modify time
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

var _configYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x52\x5f\x6f\xe2\x3e\x10\x7c\x47\xe2\x3b\xac\xf4\x7b\x6e\x08\xff\xc1\x4f\xbf\x96\x52\x95\x53\xb9\x43\x97\x54\x7d\x3c\x99\x64\x09\xae\x9c\xac\xb1\x37\x69\xb8\x4f\x7f\x8a\x83\x4a\xe8\xdd\x5b\x3c\x33\xde\xc9\x8c\x37\x42\x5b\xa1\x15\xfd\x1e\xc0\xcf\xb2\xd8\x52\x8a\x02\x52\xdc\x97\x59\x83\x3c\x33\x9b\x1d\x59\x16\xb0\x08\xc3\xd0\x6b\x50\xa6\xb1\xca\x91\x4a\x16\x30\xf3\xd0\x9b\x55\x8c\x37\xd8\xbd\x31\x7e\xe0\x23\x1e\x64\xa9\x79\x27\x33\x8c\xd4\x6f\x14\x30\xf4\x17\xb6\xb2\xee\x42\x61\x47\xba\xa2\x82\xb1\xe6\x2f\x0e\x2f\x94\x45\xb2\xc2\x9d\xe4\xa3\x00\xc7\x64\x65\x86\x03\x4d\x99\xbb\x90\x4f\x4a\xe3\x77\x99\xa3\x00\x69\x4c\x07\x5b\xd7\x2c\x20\xd0\xe4\xb3\xbc\x1a\x4d\x32\xfd\x7b\x4e\xe9\x71\xd7\x91\xf8\x46\x5e\xad\x16\x70\x64\x36\x62\x30\x18\x8e\xe6\x41\x18\x84\xc1\x50\x34\x35\x0c\x1c\x4b\x56\xc9\xf5\xc2\x26\x97\x19\x6e\x65\xdd\x06\x9a\x02\xfc\x07\xdb\x87\x2f\xf4\xbd\xd6\xf4\xb1\xae\xd9\xf9\x62\x00\xee\x20\x78\x37\x59\xe7\x1b\xaf\x07\x53\x64\xfd\xde\x3a\x97\x4a\x7b\xf1\x33\x39\x16\xe0\x72\x36\xc1\xe9\x14\x24\x94\x37\x60\xfb\x2a\x93\xd9\xd4\xfb\x38\xb4\x6d\xfe\xe5\x64\xb9\x18\x8e\x26\xf3\xc5\xff\x1d\xa9\x74\xee\x83\x6c\x2a\x80\x72\xe5\xb2\x43\xa5\x52\x57\x56\x7b\x4c\x0e\x0d\xbb\x71\x51\xf4\x22\x80\x6d\x89\xcd\xf1\xc9\x52\xfe\xef\x31\x31\x7d\xfe\xfa\x64\x39\x0e\xe7\xc3\x70\x3a\xfa\xa4\xbf\xbd\xc5\x9e\x8d\x30\xb1\xc8\x02\x30\x4d\xcf\xc9\xfb\xb9\x35\x70\x25\x5a\x01\x7b\x4d\xd9\x9d\x43\x5b\xa9\xc4\x3b\xad\x6b\xa3\x2c\x0a\x98\x8f\x9a\x0d\x78\x94\x2c\xf7\xd2\x61\xbb\x37\x0f\xf1\xd9\xa0\x80\xfc\xec\x4e\xfa\x36\xa0\x25\xe2\xdb\x50\x29\xe6\x74\xad\xe9\xfa\x54\xe3\x71\x38\x6b\x87\xb5\x37\x1b\xff\x5f\x1d\xff\x58\xee\x35\xee\x2c\x1e\x54\x7d\x21\x1b\x74\x75\x94\xd6\x35\x09\x4a\x3e\x2c\x5a\x23\xeb\xfc\x76\x0b\x88\x2f\x1d\x6d\x65\xbd\x49\x35\xae\xa8\x28\x5c\x67\xa7\x7f\x18\x2c\x2e\xd8\x38\xec\xf7\xfe\x04\x00\x00\xff\xff\xcd\xfe\x00\x4e\x59\x03\x00\x00")

func configYamlBytes() ([]byte, error) {
	return bindataRead(
		_configYaml,
		"config.yaml",
	)
}

func configYaml() (*asset, error) {
	bytes, err := configYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config.yaml", size: 857, mode: os.FileMode(438), modTime: time.Unix(1595313967, 0)}
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
	"config.yaml": configYaml,
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
	"config.yaml": &bintree{configYaml, map[string]*bintree{}},
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
