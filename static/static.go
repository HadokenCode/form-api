// Code generated by go-bindata.
// sources:
// static/migrations/1_initial.sql
// static/template/error.html
// DO NOT EDIT!

package static

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

var _staticMigrations1_initialSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\x5f\x4f\xdb\x3e\x14\x7d\xf7\xa7\x38\x8a\x90\x28\xfc\x80\x87\xdf\x78\x40\x44\x7b\x48\x9b\xdb\x92\x2d\x75\x2a\xc7\x16\xa0\x69\xaa\xbc\xc6\x74\x91\x9a\x3f\x24\xce\x18\xdf\x7e\x4a\xdb\xd0\x12\x0a\xda\xde\x22\x9f\x73\xee\x3d\x39\xf7\xda\xe7\xe7\xf8\x2f\x4b\x97\x95\xb6\x06\xaa\x64\x6c\x24\xc8\x93\x04\xba\x93\xc4\xe3\x20\xe2\x70\x9a\x26\x4d\xce\x8b\xba\x2e\x1d\x97\xb1\x7d\x7a\x6c\xb5\x35\x99\xc9\xed\xd0\x2c\xd3\xbc\x53\x8e\x15\x1f\xc9\x56\xd8\x94\x89\xb6\x66\x6e\xd3\xcc\xd4\x56\x67\xe5\xe0\x04\x82\xa4\x12\x3c\x86\x14\xc1\x64\x42\x02\x5e\x8c\xa3\x3e\xed\x88\x0d\x69\x12\x70\x06\x70\xba\xbd\xd8\xa0\xc9\x5c\x5b\x5c\x7f\xc6\xa2\xa9\x2a\x93\xdb\x1d\xd9\x65\xd8\x16\x6d\xd9\x6e\x6b\xef\xdb\xe5\xff\x57\x57\x9f\xbe\x83\x84\x88\xc4\x35\x16\x45\xb3\x4a\x90\x17\x16\x69\x62\x72\x9b\x3e\x3c\x43\xe7\x30\x8f\x8d\x5e\xa5\xf6\x19\x45\x69\x2a\x6d\x8b\x0a\x0f\x45\x05\xfb\x5c\x1a\xfc\xce\x56\x6d\x19\x20\x18\xaf\x1d\x9c\x22\x88\xe1\x07\xb1\x0c\xf8\x48\x62\x2c\xa2\x29\xa2\xd0\xbf\x38\x85\xbc\x21\xbe\x61\xfe\xa5\xd5\x2d\xb7\xe7\x17\xa0\x30\xa6\x1e\x18\x85\x7e\x07\x72\x1f\xc1\xd8\x65\xc4\x7d\x97\xbd\x0d\x0b\xa1\xc7\x27\xca\x9b\x10\xca\x55\xb9\xac\x1f\x57\xee\xe1\x11\x51\x9e\xbc\xcc\x56\xde\xcf\x08\xb1\xf4\xa4\x8a\xdb\x09\x10\x57\x53\x0c\x8e\x4d\xae\x7f\xac\x4c\x72\x7c\x76\x9c\xa4\xf5\xe6\xf3\xc4\xdd\x69\xbc\x61\x48\x70\x1e\x8a\x2a\x9b\xd7\x8b\x9f\x26\xd3\x0e\x06\x0c\x9b\xed\x70\xa0\x54\xe0\x83\x47\x12\x5c\x85\x21\x66\x22\x98\x7a\xe2\x1e\x5f\xe9\x1e\x3e\x8d\x3d\x15\x4a\xb4\xbc\xf9\xd2\xe4\x6d\xdc\x66\xfe\xeb\x72\x70\x72\xb6\x96\xd7\xa6\xea\xc9\xd7\xe7\x5d\x8f\xbb\x69\xd8\x03\xac\xb6\x4d\xed\x74\xfe\x5f\x7a\x76\x7d\x76\xff\xd1\xb2\x17\x95\xd9\xce\xc4\x81\x0c\xa6\x14\x4b\x6f\x3a\x7b\x2b\xca\x8b\xa7\xce\xcf\xcb\x10\xf7\x04\x6c\x3f\x88\xed\xea\x6e\xa2\x48\xb4\xd5\xf3\xad\xc6\xc1\x90\xc6\x91\x20\xa8\x99\xdf\x32\xdb\xbb\xf3\x2a\xaf\x71\x24\x40\xde\xe8\x06\x22\xba\x05\xdd\xd1\x48\x49\xc2\x4c\x44\x23\xf2\x95\xa0\x03\xf7\xe5\x70\xfc\x6d\xcf\x6d\xf8\x6d\xf4\x31\x89\xc0\x7b\x95\xf9\xd9\x3b\x73\x59\x9f\x6f\xd4\x5f\xe2\x88\x0f\x5f\x03\xff\x12\xd5\x3a\x90\xfd\x3d\xf3\x8b\xa7\x9c\x31\x5f\x44\xb3\xb7\x56\xdd\x0e\xf8\x20\xb9\x7e\x56\xee\x81\x62\x7d\x68\xb7\xc5\xdd\xd1\x07\x4f\x4f\x47\x79\xe7\x5d\xfb\x13\x00\x00\xff\xff\xab\x0e\x95\x1f\x06\x05\x00\x00")

func staticMigrations1_initialSqlBytes() ([]byte, error) {
	return bindataRead(
		_staticMigrations1_initialSql,
		"static/migrations/1_initial.sql",
	)
}

func staticMigrations1_initialSql() (*asset, error) {
	bytes, err := staticMigrations1_initialSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/migrations/1_initial.sql", size: 1286, mode: os.FileMode(420), modTime: time.Unix(1511457221, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _staticTemplateErrorHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\x51\x74\xf1\x77\x0e\x89\x0c\x70\x55\xc8\x28\xc9\xcd\xb1\xe3\xb2\x01\x51\x0a\x39\x89\x79\xe9\xb6\x4a\xa9\x79\x4a\x20\x81\xd4\xc4\x14\x3b\x2e\x05\x05\x05\x05\x9b\xdc\xd4\x92\x44\x85\xe4\x8c\xc4\xa2\xe2\xd4\x12\x5b\xa5\xd0\x10\x37\x5d\x0b\x25\xa8\x54\x49\x66\x49\x4e\xaa\x9d\x6b\x51\x51\x7e\x91\x42\x75\xb5\x82\x9e\x73\x7e\x4a\xaa\x42\x6d\xad\x8d\x3e\x44\x82\xcb\x46\x1f\x62\x8e\x4d\x52\x7e\x4a\x25\x88\x0b\xa3\x21\xd6\x02\x02\x00\x00\xff\xff\x6d\x33\x58\x59\x87\x00\x00\x00")

func staticTemplateErrorHtmlBytes() ([]byte, error) {
	return bindataRead(
		_staticTemplateErrorHtml,
		"static/template/error.html",
	)
}

func staticTemplateErrorHtml() (*asset, error) {
	bytes, err := staticTemplateErrorHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/template/error.html", size: 135, mode: os.FileMode(420), modTime: time.Unix(1511438765, 0)}
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
	"static/migrations/1_initial.sql": staticMigrations1_initialSql,
	"static/template/error.html": staticTemplateErrorHtml,
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
	"static": &bintree{nil, map[string]*bintree{
		"migrations": &bintree{nil, map[string]*bintree{
			"1_initial.sql": &bintree{staticMigrations1_initialSql, map[string]*bintree{}},
		}},
		"template": &bintree{nil, map[string]*bintree{
			"error.html": &bintree{staticTemplateErrorHtml, map[string]*bintree{}},
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

