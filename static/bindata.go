// Code generated by go-bindata.
// sources:
// static/migrations/1_initial.sql
// static/migrations/demo/data.sql
// static/templates/error.html
// static/templates/redirect.html
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

var _staticMigrations1_initialSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x94\x4f\x6f\xda\x40\x10\xc5\xef\xfb\x29\x9e\xac\x48\x21\x69\x92\x43\x9b\x43\x14\xab\x07\x83\xc7\xc4\xad\x59\xa3\xdd\xb5\x42\x54\x55\xc8\xc5\x1b\x6a\x09\xff\x89\xbd\x6e\x9a\x6f\x5f\x19\x70\x20\x86\x46\xed\x0d\xed\x7b\x6f\x66\xf8\xcd\xc0\xe5\x25\x3e\x64\xe9\xb2\x8a\x8d\x46\x54\x32\x36\x12\xe4\x28\x02\xcd\x14\x71\xe9\x87\x1c\xbe\x07\x1e\x2a\xd0\xcc\x97\x4a\xc2\x6a\x9a\x34\xb9\x2c\xea\xba\xb4\x6c\xc6\xf6\xc3\xd2\xc4\x46\x67\x3a\x37\x43\xbd\x4c\xf3\xae\x8e\x17\xf1\x91\x6a\xcb\x34\x65\x12\x1b\x3d\x37\x69\xa6\x6b\x13\x67\xe5\xe0\x0c\x82\x54\x24\xb8\x84\x12\xfe\x78\x4c\x02\x8e\xc4\x49\xdf\x76\xc2\x86\x34\xf6\x39\x03\x38\xdd\x5f\x6d\xd4\x64\x1e\x1b\xdc\x7e\xc6\xa2\xa9\x2a\x9d\x9b\x9d\xd9\x66\xd8\x16\x6d\xdd\x76\x3b\xde\xb7\xeb\x8f\x37\x37\x9f\xbe\x83\x84\x08\xc5\x2d\x16\x45\xb3\x4a\x90\x17\x06\x69\xa2\x73\x93\x3e\xbe\x20\xce\xa1\x9f\x9a\x78\x95\x9a\x17\x14\xa5\xae\x62\x53\x54\x78\x2c\x2a\x98\x97\x52\xe3\x77\xb6\x6a\xcb\x60\x8d\x81\xee\xaf\xce\xe1\x4b\xb8\xbe\x54\x3e\x1f\x29\x78\x22\x9c\x20\x0c\xdc\xab\x73\xa8\x3b\xe2\x1b\xe7\x3f\x8e\xba\xf5\xf6\xe6\x05\x28\x90\xd4\x13\xc3\xc0\xed\x44\xee\xc2\xf7\x6c\x46\xdc\xb5\xd9\x21\x2c\x04\x0e\x1f\x47\xce\x98\x50\xae\xca\x65\xfd\xb4\xb2\x8f\xaf\x88\xf2\xe4\x75\xd3\xea\x61\x4a\x90\xca\x51\x91\x6c\x37\x40\x3c\x9a\x60\x70\xaa\xf3\xf8\xc7\x4a\x27\xa7\x17\xa7\x49\x5a\x6f\x3e\x9e\xd9\xbb\x8c\x33\x0c\x08\xd6\x63\x51\x65\xf3\x7a\xf1\x53\x67\xb1\x85\x01\xc3\xe6\x3a\x2c\x44\x91\xef\xae\x8f\x86\x47\x41\x80\xa9\xf0\x27\x8e\x78\xc0\x57\x7a\x80\x4b\x9e\x13\x05\x0a\xad\x6f\xbe\xd4\x79\x8b\x5b\xcf\x7f\x5d\x0f\xce\x2e\xd6\xf1\x5a\x57\xbd\xf8\xfa\xbd\xeb\x31\x9b\x04\x3d\xc1\xc4\xa6\xa9\xad\x6e\xfe\xd7\x9e\x5d\x9f\xdd\xf7\x68\xdd\x8b\x4a\x6f\x77\x62\x41\xf9\x13\x92\xca\x99\x4c\x0f\x43\x79\xf1\xdc\xcd\xf3\xba\xc4\xbd\x00\xdb\x07\xb1\x3d\xdd\x0d\x8a\x24\x36\xf1\x7c\x9b\xb1\x30\x24\x2f\x14\x84\x68\xea\xb6\xce\x90\xf7\x78\x79\xa1\x00\x39\xa3\x3b\x88\xf0\x1e\x34\xa3\x51\xa4\x08\x53\x11\x8e\xc8\x8d\x04\x1d\xf9\xbd\x1c\xc7\xdf\xf6\xdc\xc2\x6f\xd1\x4b\x12\xbe\xf3\x86\xf9\xc5\x5f\xf6\xb2\x7e\xdf\xa4\xbf\xc8\x90\x0f\xdf\x0a\xff\x83\x6a\x0d\x64\xff\xce\xdc\xe2\x39\x67\xcc\x15\xe1\xf4\x70\x54\xbb\x13\xde\x21\xd7\x67\x65\x1f\x29\xd6\x97\x76\x57\xdc\x3d\xbd\xf3\xd7\x63\xb3\x3f\x01\x00\x00\xff\xff\x54\xf2\x54\x53\xf7\x04\x00\x00")

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

	info := bindataFileInfo{name: "static/migrations/1_initial.sql", size: 1271, mode: os.FileMode(420), modTime: time.Unix(1521979138, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _staticMigrationsDemoDataSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x95\xd1\x6a\xe3\x46\x17\xc7\xef\xf5\x14\x87\x21\xe0\x84\xcf\x92\x63\x5b\x76\xec\x24\x0e\xdf\xee\x5a\xbb\x1b\x48\x9d\xe2\x24\x5d\x4a\x58\xc2\x68\x34\xb6\x07\x4b\x23\x55\x1a\xad\xb3\x2c\x85\x34\xbd\x69\xaf\x42\x7b\x55\x0a\xa5\xed\x1b\xb8\xa5\xa6\xe9\x6e\xec\x7d\x85\xa3\x37\x2a\x92\xe3\xd4\x71\x1c\xda\xc0\x5e\x94\xde\xd8\xcc\x68\xce\xff\x7f\xf4\x3b\x47\x73\x74\x1d\xfe\xe7\x89\x6e\x48\x15\x87\xa3\x40\xd3\xe6\xd7\x07\x8a\x2a\xee\x71\xa9\x1e\xf3\xae\x90\x5a\x73\x1f\x56\x56\x34\x80\xa6\xf5\x64\xef\x51\xdb\x02\x87\x7b\xfe\x51\xc4\x43\x38\x3a\xda\x6d\xc2\x66\x03\xe2\x58\x38\x27\x5d\x2e\x79\x1a\x7d\xf2\xca\x5c\x5d\xdb\x9a\x3b\xce\x3d\x2a\x5c\x00\xb8\x39\x9e\x33\x8b\x8c\x56\xf8\x7a\x5d\x2f\x33\x5e\xd2\xcd\xf5\xba\xa9\xdb\xc5\xf5\x9a\x5e\xe6\x8c\x95\x2a\x1b\xac\xda\xa1\x66\x6e\x5e\xa2\x2b\x54\x2f\xb6\x2d\x39\x2f\xb1\x5e\xa9\x57\x37\xaa\x7a\x95\xd9\x55\xdd\x2c\xd9\x4c\xaf\x77\xec\x8a\x5e\x73\x3a\xc5\x6a\xbd\x6e\x57\x6b\xb5\xe2\x12\x89\x76\x3c\x2f\xe1\x94\xcb\x76\xd9\xd1\x4b\xd5\x0e\xd7\xcd\x72\x85\xe9\x75\x73\xbd\xa4\xb3\x1a\x37\x69\xb1\x53\xdb\x60\x6c\x2a\xf1\xd8\x7a\xb6\xdb\xd2\xd2\x37\xd8\x6d\x1d\x58\xed\x43\xd8\x6d\x1d\xee\x03\xe9\xf8\xa1\x77\x12\xb1\x1e\xf7\x28\x81\x55\x92\x42\x20\x79\x12\x47\x3c\x24\x79\x72\xbd\xbf\x06\x9f\x3c\xda\x3b\xb2\x0e\xb2\x68\x80\xd5\x8c\x45\x7e\x06\x30\x9f\xd3\xb6\x53\x15\x70\xa9\xec\x36\x08\x97\x04\x94\x50\x2e\x6f\x10\x2b\x63\x16\xc5\x76\xc4\x42\x11\x28\xe1\x4b\x02\x94\xa5\xff\x0d\xd2\x53\x2a\x88\x36\x0b\x85\x3e\xf5\x84\x6b\x44\xd4\x13\xdd\xd8\x75\x85\x34\x84\xec\xf8\x05\x02\x1e\x57\x3d\xdf\x69\x90\xc0\x8f\x14\x01\x2e\x99\x7a\x1d\xf0\x06\xa1\x41\xe0\x0a\x46\x53\x8d\xc2\xa9\x3e\x18\x0c\xf4\xd4\x5a\x8f\x43\x97\x4b\xe6\x3b\xdc\x21\x3b\x1a\xc0\xb6\x90\x41\xac\x40\x52\x8f\x37\x48\x96\x2d\x81\x69\xfc\x6c\x31\x97\x20\x01\x8f\x9e\xba\x5c\x76\x55\xaf\x41\xaa\x26\x81\x90\x7f\x16\x8b\x90\x3b\x0d\x52\x24\x85\x1d\x6d\xbb\x90\x5a\xec\x68\xb9\xb5\xfc\x0c\xc0\xac\x92\xff\x80\xc1\x33\xa1\x9e\xc7\x76\xd6\x6d\x10\xd0\x2e\xbf\x87\x40\xd4\x37\xa6\xa2\x86\xf0\x33\x3f\x9d\x06\xe2\x83\x62\x48\x7f\x67\x14\x14\x3f\x55\x37\x19\xb6\xb2\x07\x81\x4b\x19\xef\xf9\xae\xc3\xc3\xe9\x96\x61\x18\xb7\xc8\x94\x2a\x77\xc8\x2c\x38\x74\x38\x77\x6c\xca\xfa\x4b\x5d\x9e\xde\x3c\xbc\xe5\xf4\xa9\x1f\x87\x30\x0b\xbc\x6b\x79\xd7\xf3\xde\x6a\xb4\xe3\xfb\xaa\x11\xc6\x8b\xd5\xc0\xdf\x70\x84\x57\x38\xf9\xf7\x15\x03\xbf\xc3\xab\xe4\x62\x01\xd2\x74\xf3\x03\x17\x04\xbf\xc7\x09\x5e\xe1\x15\x8e\x70\x9c\x9c\xe3\x30\x39\xc3\x4b\xfc\x63\xd1\xf9\x5b\x1c\x26\x5f\x01\xbe\x5d\x76\xf6\x61\xe5\xda\xba\xe7\xee\x71\xa8\x9a\xbf\x79\xb2\xe5\xf2\x0b\x27\xf7\xe6\xfa\xf3\xdd\x3c\x26\x8a\x47\xea\xff\xde\x6b\x63\xba\xf1\xf2\xf3\x65\x1f\x67\xee\xcd\x94\xf3\xe6\x31\x79\x62\x40\xcb\x0f\x55\x8f\xf9\x8a\xc3\xc7\x34\xec\x0b\x19\xf9\x92\xbc\xcc\xff\xc5\x68\xf3\x98\xbc\xf0\xc3\x3e\x30\x5f\xaa\x90\x32\x15\x81\xf2\xa1\x23\x14\x08\x09\xaa\xc7\x41\x09\x8f\xc3\x80\x43\x57\xbc\xe2\x20\x94\xb1\xcc\xb4\x1d\xcf\x9b\xe2\xcf\x78\x99\x91\x7a\x07\xf8\x03\x4e\x92\xb3\xe4\x3c\x05\x99\x9c\x03\xfe\x98\x31\x7c\x8b\x97\x38\x4e\xbe\xc0\x09\x8e\x17\x53\xc1\x9f\x70\x88\xbf\xa4\x87\x71\x08\xf8\x3b\x0e\xf1\x3d\x4e\xf0\x1d\x8e\x93\x0b\x1c\xa5\x0a\xbf\x26\x67\x69\x0f\x27\x17\x79\xc8\x4e\xbd\x4f\xbe\x4c\xbe\x4e\xeb\x83\x63\x9c\xe0\x08\x70\x9c\x06\x8e\x71\x94\x7c\x33\x4d\x35\xe5\x6f\xb5\x9a\x5b\xda\xca\xca\xd6\xf2\x39\x69\x49\xe7\xf6\x04\x6d\xfa\x03\xf9\xa0\x99\xfa\x1f\x1a\x92\x4d\x6b\xcf\x3a\xb4\xe0\x69\x7b\xff\xa3\x5b\x8d\xfa\xe2\xb9\xd5\xb6\x60\xda\xad\xb0\xdb\x9a\xf5\xe6\x4d\xd3\xcd\xcc\xaf\xfb\xfd\xae\xcc\x6c\xd6\x3e\x48\xe8\xef\x0b\xf7\x67\x00\x00\x00\xff\xff\xb4\x87\x17\xcb\x0a\x09\x00\x00")

func staticMigrationsDemoDataSqlBytes() ([]byte, error) {
	return bindataRead(
		_staticMigrationsDemoDataSql,
		"static/migrations/demo/data.sql",
	)
}

func staticMigrationsDemoDataSql() (*asset, error) {
	bytes, err := staticMigrationsDemoDataSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/migrations/demo/data.sql", size: 2314, mode: os.FileMode(420), modTime: time.Unix(1517513033, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _staticTemplatesErrorHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x58\x5b\x57\xe3\x38\x12\x7e\xe7\x57\xd4\x7a\x99\x49\xb2\x8d\xed\x40\xb8\x13\x73\x4e\xb6\xb9\x0c\xd0\x30\x24\x81\x86\x9e\x97\x3d\x8a\xad\xd8\x0a\xb6\xe4\x96\xe4\x5c\x36\xc7\xff\x7d\x8f\xe4\x4b\x1c\x92\xb0\x3d\xdd\xe9\x07\x4b\x25\xe9\xd3\x57\xa5\xaa\xa2\xaa\xdb\xff\xb8\xf8\xf3\xf3\xd3\xb7\xc7\x4b\x08\x64\x14\x9e\x6f\xb5\xd5\x07\xdc\x10\x09\xe1\x18\x94\x99\x23\x61\x40\x88\xa8\xef\x18\xf3\x39\x58\x7d\x37\xc0\x11\xb2\xbe\x20\xea\x27\xc8\xc7\x90\xa6\x86\x3a\x82\x91\x77\xbe\x05\x00\xd0\x8e\xb0\x44\xe0\x06\x88\x0b\x2c\x1d\xe3\xf9\xe9\xca\x3c\x36\xf2\x25\x49\x64\x88\xcf\x2b\x28\x4f\x4a\x00\x69\xda\xb6\xb3\xa5\x0a\x42\x20\x65\x6c\xe2\xef\x09\x19\x3b\x06\xc7\x43\x8e\x45\x60\x80\xcb\xa8\xc4\x54\x66\x4c\x2e\x70\x88\x66\x56\x1f\xbb\x8c\x7a\x02\xd2\xf4\x2c\xe1\xa1\xa3\x16\x7a\xd8\x23\x1c\xbb\x32\xe3\xa6\x21\x43\x42\xdf\x80\xe3\xd0\x31\x84\x9c\x85\x58\x04\x18\x4b\x03\x02\x8e\x87\x8e\xa1\x6e\x12\xa7\xb6\x1d\xa1\xa9\xeb\x51\x6b\xc0\x98\x14\x92\xa3\x58\x4d\x5c\x16\xd9\xa5\xc0\xde\xb7\x9a\x56\xd3\x76\x85\x58\xc8\xac\x88\x50\xcb\x15\xc2\xd0\xf7\x64\x3f\x42\x25\xf6\x39\x91\x33\xc7\x10\x01\x6a\x1d\xef\x9b\xd7\xf4\xa0\x75\xbc\x3f\xfd\xde\xdd\x45\xec\xe5\xb5\xf3\xa9\x79\x70\xdc\x7b\x7d\x9c\x3e\xfa\x87\xc3\xd9\xfe\xcd\xcb\xf8\xe9\x21\x68\x5e\xee\x1d\xb6\x5e\xa3\x2b\xf7\x36\xec\x77\x26\xe4\xda\xbf\xea\xbc\xd8\x5e\x87\xf4\x0f\x6f\x5f\x23\x03\x5c\xce\x84\x60\x9c\xf8\x84\x3a\x06\xa2\x8c\xce\x22\x96\x88\x42\x3d\xad\xd4\x79\x49\x41\x3d\xe0\x0e\x0c\x98\x37\x83\x39\x04\x98\xf8\x81\x3c\x85\xdd\x66\xf3\xb7\x33\x48\xcb\x4d\xf9\xb2\x47\x44\x1c\xa2\xd9\x29\x0c\x43\x3c\x3d\x03\x14\x12\x9f\x9a\x44\xe2\x48\x9c\x82\x8b\xa9\xc4\xfc\x0c\x06\xc8\x7d\xf3\x39\x4b\xa8\x67\xba\x2c\x64\xfc\x14\xfe\x39\x3c\x50\xff\xaa\x78\x96\x8b\xb8\x07\x73\x98\x10\x4f\x06\xc5\x75\x11\x9a\x9a\xb9\xe0\xb0\xd9\x8c\xa7\x4a\xc2\x7d\x42\x4f\xa1\x09\x28\x91\x6c\x15\xc0\x8a\x39\xf3\x39\x16\xe2\x03\xea\x96\xf6\xca\xa5\xad\xa5\x1a\x94\x51\x5c\xec\x6d\xdb\x15\xc3\xb4\x85\xcb\x49\x2c\x17\x56\x9a\x10\xea\xb1\x89\xd5\x7f\xea\xf4\x9e\xc0\x81\x0b\x24\xb1\x45\xd9\xa4\xde\x38\x2b\xb7\xd4\x87\x09\x75\x25\x61\x14\xea\x94\x79\xb8\x01\x73\x18\x23\x0e\x2e\x38\xa0\xe6\x96\x0e\x92\x2f\x44\xc8\x33\x70\x2d\x8e\x23\x36\xc6\xf5\x9a\x26\x57\x6b\x28\x11\xf2\xbc\x7a\x2d\x9b\xa4\xf5\xfc\x3e\x8f\xb9\x49\x84\xa9\x2c\x07\x97\x21\x56\x9f\x46\x7e\x6f\xdb\x2e\x88\xb6\xed\x2c\xac\xda\xea\xa9\xce\xb7\xe6\x73\x13\xb6\xe1\xd4\x01\x0b\xcc\x34\xd5\xd3\x09\x91\x41\x11\x4c\x90\xa6\xf3\x39\x6c\x0f\x19\x8f\x16\x9b\xda\x1e\x19\x17\xa1\xac\xcc\x5b\x38\xcc\x3b\xb1\xa9\x2e\xc2\xdc\xd0\xb1\xb9\x08\x4a\x8f\x8c\x37\xec\x57\x8c\x8c\x85\x29\xdb\xfa\xd6\xf9\x3c\x27\x74\x73\x01\x69\x4a\xbc\x2c\x4c\x55\x14\xaa\x25\x4c\x3d\x4d\x09\x96\x7e\x8b\xc4\x52\xcd\x28\xa0\xd3\x41\x26\x2f\xe8\x18\x80\xf4\x53\x64\xd2\x4e\xf6\x2c\x69\x6a\xbc\x03\x5c\x98\xe5\x1e\xcb\x80\x79\x90\xa6\x10\xe9\xd1\x82\xce\x46\x36\x8b\xc3\x97\xd4\x65\x1e\xa1\xfe\xd3\x2c\x56\x97\x03\xa6\xae\x9c\xc5\x78\x1d\xc6\xf9\xd6\x7b\x08\x8e\xa8\x8f\xc1\xba\xa1\x71\x22\xc5\x9a\x6b\x96\x0d\xaa\x6c\x67\xaa\xf0\x8a\x81\xb3\x89\x71\xbe\xb2\xb9\x80\x25\x43\x40\xd4\xd3\xe6\xcd\xad\xb2\x0e\x5a\xc3\x87\x68\x80\x43\x18\x32\x9e\x11\xd6\x0f\x62\x94\x2f\xc8\x42\x53\x44\x66\x0b\xd4\x40\xdf\x9e\x6d\x5f\x9e\x9a\xa1\xff\xde\x1f\xb4\x7c\x3d\xbf\x25\x0f\xc9\xf0\x4f\x3e\x50\x65\xbd\xf9\x8b\x75\xd8\x1e\xa3\x90\x78\xca\x8d\x29\x93\x50\xdf\xb6\x2e\x39\x67\xdc\xfa\x03\x09\x3d\x00\xab\x01\x9b\x34\x27\xca\xe8\x4b\xa6\x55\x7f\x34\x38\xd3\xd6\x28\x27\x66\xe8\x03\x11\xe6\x7c\xae\x8c\x9a\xdf\x96\xa6\xfa\xab\x1e\x36\x14\x4a\x61\x42\x4b\x01\xf5\xd6\xb8\xda\x8a\x56\x15\xef\x87\x1f\x73\xff\xa5\x1f\x45\x51\xee\x60\x0f\x28\xca\xe3\xa0\xf4\xb9\xdc\x11\x7f\x90\x43\xf1\x66\xd5\x40\xfa\xd0\xf3\x37\xe0\x3c\x86\xc8\xc5\x01\x0b\x3d\xcc\x15\x5a\xbc\x98\xfe\x3c\xe6\x57\x14\x26\x9a\xdb\x58\x0d\x7e\x1e\xe7\x9e\xd0\x2f\x98\xfa\x32\xd0\x21\x4e\x68\xa8\x27\xbf\x80\x87\xa6\x15\x3c\x34\xfd\x55\xbc\x9e\xaa\x5e\x38\xd6\x19\x88\xe7\xe3\x4d\x59\xa3\x7a\x7a\x35\xcc\x8b\x5c\xfc\x01\x83\x4a\xb2\xae\x62\xad\xdb\x9e\x87\x48\xe6\x6d\xff\xe1\x79\xb1\x54\xf8\x5a\x40\x3c\x0f\x53\xa3\xf2\x38\xdb\x6b\x0a\xaa\x12\xeb\xef\xe4\xb1\x35\x39\xa2\xb5\x21\x47\x2c\x87\xf1\x40\x52\x18\x48\x6a\x7a\x2a\xaf\xf2\x82\xa9\x48\x06\x11\x91\xeb\xee\x59\xb5\xc5\x86\xeb\x37\xa5\xa8\xea\xd6\xb2\xc8\x90\x78\x2a\xcd\xac\x24\xda\x70\x6c\xd3\x51\x73\x80\x38\x54\x27\xa6\x90\x9c\xc4\xd8\x5b\x16\x22\x4a\x22\x24\xb1\x07\x03\xdf\x14\x89\xeb\xe2\xe5\xb2\x72\xcd\x4f\x97\x37\x8e\x51\x29\xb9\x74\xca\xde\x5e\xa9\x8d\x37\x98\xe4\x23\x6b\xad\x8a\xdf\x89\xda\xb6\x7a\xf1\xbc\x40\xc8\x96\xf2\x4f\xd5\xf3\xf2\xb2\x0b\x04\x77\x17\xa5\xb6\xab\x8a\xa7\xd1\xf7\x04\xf3\x99\x2e\xb1\xb3\xa1\xd9\xb2\xf6\xac\x5d\x4b\x84\x24\xd2\x65\xf5\xa8\xa2\xfe\x6a\x4d\x7d\x77\xdb\x62\x7b\x17\x77\xf2\xe6\x6d\xfc\xed\xe6\xae\xf5\x7c\xf9\xf0\xdf\xe8\xfe\xe8\xee\xf3\x5b\x8f\xdb\xfc\xf2\xc4\xee\xc6\xfe\x21\xea\xfc\x75\x7d\x3b\xb9\xba\xb8\xff\xfa\xd0\xb1\xaf\xe3\xeb\xab\xab\x93\x56\xf0\x1a\x5f\x1f\xdc\xbd\x3d\x2c\xb0\x37\xd5\xd6\x95\x4a\x6c\xad\x12\x1e\x1d\x09\xcb\x0d\x59\xe2\x0d\x43\xc4\xb1\xd6\x04\x8d\xd0\xd4\x0e\xc9\x40\xd8\x31\x8b\x63\xcc\xad\x91\xb0\x77\xad\xdd\x3d\xeb\xc4\x4e\x22\xaf\x10\xfe\x7f\xed\x3a\xf1\xc3\xc0\x0f\x4e\xfe\xfd\xe9\xdb\x6e\xf7\x4e\x8e\x5b\x3d\x7a\xf4\xd2\x8a\xfc\xc7\x69\xf0\x7c\x72\x67\xf7\xdd\xae\xe8\x3c\x1e\x05\xcf\x64\xf0\xda\x3a\x19\x1d\x0d\xd1\xdb\xd5\xa3\x78\x1b\xbf\x26\x62\x3c\x44\xcd\xc1\x7e\xf7\x97\xb5\xfb\xd1\x6e\x68\xf4\xbe\x19\xfa\x58\xaf\xdb\xbf\x7a\x87\xfd\x18\x8f\x82\xfd\xe7\xe6\x9e\x77\x3c\xfa\x53\x1e\x8e\xbf\x5c\xfe\x31\xc4\xf6\x6d\xf7\x9a\xf4\x7a\xfd\x6e\x77\xda\x1f\x5e\xbd\xc4\x64\xf7\xfe\x7b\xf2\xd5\xeb\xcc\x46\xcf\x88\x1f\x7c\x3a\x3a\x7c\xfc\xfa\x39\xfa\x16\xfe\x84\x5e\xe7\x5b\x95\x32\x7e\x7b\x07\x84\x44\x5c\x36\x60\xae\x91\x6a\x89\xc0\xa0\xa2\xd1\x95\xb5\xac\xfe\x76\x19\x15\x12\xb6\xcb\xb0\x77\x60\xbb\x5e\x7b\xd7\x9b\xa8\x70\xad\xe5\xf5\xba\x6a\x08\x24\x89\x30\x4b\xe4\x8e\x1e\xf0\x1d\x70\x19\x57\xf9\x92\x30\x5a\x60\x16\x73\x70\xe0\x1e\xc9\xc0\xd2\xdd\x54\xbd\xbe\x68\x3a\xc0\x2c\x88\xd9\x2a\x90\x9b\x39\x7a\x8e\x0c\x4e\x06\x0d\x0e\xc4\xaa\xb3\xbe\xa1\xb2\x5e\x52\xb4\x54\x66\xaa\x37\x2a\x27\x38\x98\xce\x0a\x09\x81\xe5\x8d\x4a\x5e\x63\x14\x56\x0c\x52\x18\xa2\x3c\x6a\x9a\x8b\xfe\x27\xb3\x45\x8c\xb9\x4a\x7b\xaa\x36\x5f\x62\xaf\x68\x36\xe1\x5f\xf9\x8d\x76\xc1\x35\xd7\x60\x81\xf2\x8e\xa8\xde\xde\x50\x4d\x73\xbd\xa6\x33\x57\x6d\xa7\x7a\xc5\x27\xa8\xfd\x56\xab\xf4\x60\x64\x08\xf5\xca\x72\x1b\xf6\x9b\xf0\xfb\xef\x15\xd0\x00\x89\xcf\x2a\xdf\xd6\x6b\x03\xdf\x9c\x20\x4e\x09\xf5\x6b\x8d\xaa\x5e\xcb\x24\xb2\x3e\x6d\xcd\x11\xd5\xae\x2d\xc4\xd9\x1f\x9b\x2a\x93\x34\x2b\x0c\x57\x08\x1d\x7f\x44\x28\x4f\xe4\x7f\x87\x50\x79\x64\x99\x50\xc9\xb3\xc2\x48\x8f\xd2\x9d\xc2\x63\xca\x26\x73\xd4\x55\x09\x75\x67\xa9\xc7\x6d\x34\xb6\xaa\xbd\x65\xd6\x54\xb6\xed\xec\x3f\x7d\xfe\x17\x00\x00\xff\xff\x6a\xc8\xc3\xe0\x05\x12\x00\x00")

func staticTemplatesErrorHtmlBytes() ([]byte, error) {
	return bindataRead(
		_staticTemplatesErrorHtml,
		"static/templates/error.html",
	)
}

func staticTemplatesErrorHtml() (*asset, error) {
	bytes, err := staticTemplatesErrorHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/templates/error.html", size: 4613, mode: os.FileMode(420), modTime: time.Unix(1517593349, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _staticTemplatesRedirectHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x94\xdf\x6e\xe2\x3a\x10\xc6\xef\xfb\x14\x3e\xb9\xad\xb0\x81\xf0\xf7\x1c\x82\x94\x53\x08\x5b\x68\x29\x21\x50\xa0\x77\xc6\x71\x12\x87\xc4\x0e\xb6\x43\xc9\x56\xbc\xfb\x0a\xd2\xdd\x56\xbb\xea\x6e\xb5\xbe\xb1\xfd\x8d\x35\xf3\xfd\x46\x23\xf7\xfe\x19\x3c\xdc\x2c\x36\xb3\x21\x88\x74\x9a\xf4\xaf\x7a\xe7\x0d\x24\x98\x87\x96\xf1\xf2\x02\xa0\x47\x22\x9a\x62\x78\x87\x79\x98\xe3\x90\x82\xd3\xc9\x38\x3f\xa2\xd8\xef\x5f\x01\x00\x40\x2f\xa5\x1a\x03\x12\x61\xa9\xa8\xb6\x8c\xe5\xc2\xa9\x74\x8c\xd7\x90\x66\x3a\xa1\xfd\x77\x59\x16\x67\x01\x9c\x4e\x3d\x54\x86\xde\x65\x88\xb4\xce\x2a\x74\x9f\xb3\x83\x65\x48\x1a\x48\xaa\x22\x03\x10\xc1\x35\xe5\xba\x74\x32\xa0\x09\x2e\xa0\x47\x89\xe0\xbe\x02\xa7\xd3\x7f\xb9\x4c\xac\x73\x60\x4e\x7d\x26\x29\xd1\xa5\xb7\x4b\xca\x84\xf1\x1d\x90\x34\xb1\x0c\xa5\x8b\x84\xaa\x88\x52\x6d\x80\x48\xd2\xc0\x32\xce\x95\xd4\xbf\x08\xa5\xf8\x48\x7c\x0e\xb7\x42\x68\xa5\x25\xce\xce\x17\x22\x52\xf4\x43\x40\x0d\x58\x85\x55\x44\x94\x7a\xd3\x60\xca\x38\x24\x4a\x19\x97\x3a\xe5\x62\x5c\xd3\x50\x32\x5d\x58\x86\x8a\xb0\xd9\x69\x54\x46\xbc\x69\x76\x1a\xc7\xbd\x5b\xc3\x62\xb5\xb6\xaf\xab\xcd\xce\x7c\x3d\x3b\xce\xc2\x56\x50\x34\x6e\x57\x87\xc5\x34\xaa\x0e\xeb\x2d\x73\x9d\x3a\x64\x9c\x78\xf6\x33\x1b\x85\x8e\xbd\x42\xbe\xcd\xbc\xd6\x78\x9d\x1a\x80\x48\xa1\x94\x90\x2c\x64\xdc\x32\x30\x17\xbc\x48\x45\xae\xce\xad\x47\x65\xef\x7b\x5b\xe1\x17\xfd\xab\x9e\x22\x92\x65\x1a\x28\x49\xde\xc8\x88\xf0\x29\x8c\xf7\x39\x95\xc5\x85\xa8\x3c\x56\x4c\x58\x87\x35\xa8\x12\x96\x5e\x28\xe2\x77\x10\xbf\x22\x4c\xc6\xa6\xa8\x0f\x26\xfa\x76\x77\xd8\xdc\x4e\xcc\xe5\x70\xfa\x35\xbd\x6f\x4f\x6e\x76\x73\x89\xe4\xb0\x8b\xdc\x2c\x6c\x61\xfb\x69\x34\x7e\x76\x06\xf7\x8f\x53\x1b\x8d\xb2\x91\xe3\x74\xcd\x68\x9d\x8d\x9a\x93\xdd\xf4\x63\x84\x1e\x2a\x3d\x7f\x64\xde\xe7\xb1\x82\x24\x11\xb9\x1f\x24\x58\xd2\x0b\x01\x8e\xf1\x11\x25\x6c\xab\x50\x26\xb2\x8c\x4a\x18\x2b\x54\x83\xb5\x3a\xec\xa2\x3c\xf5\xbf\x8b\x7f\xa6\xb2\xb3\xe9\x36\x8c\xba\xff\x5f\x6f\x6a\xee\x44\x1f\xcc\x39\x6f\xaf\xcc\x34\x9c\x1d\xa3\x65\x77\x82\x3c\xe2\x2a\x7b\xd6\x8e\x96\x6c\xbb\x36\xbb\x71\x3b\xc0\x3b\x67\xa6\x76\x87\x75\xae\x0e\x01\xae\x6e\x1b\xee\x5f\x53\x7d\x76\xd8\xe2\x9f\x67\xed\xf7\x3c\xe3\xa7\x79\xcb\xcb\x68\x1c\x35\x96\xd5\xba\xdf\x89\x1f\x74\xeb\x70\x37\xfc\x12\x50\x34\x76\x47\x6c\x3e\xf7\x5c\xf7\xe8\x05\xce\x2a\x63\xb5\xfb\x7d\xfe\xe8\xdb\x45\xbc\xc4\xb2\x79\xdd\x6e\xcd\x1e\x6f\xd2\x4d\xf2\x29\x1e\xf4\x3a\x6a\xa8\xfc\x21\xbe\x05\x00\x00\xff\xff\xb8\x76\xa4\x4e\x32\x04\x00\x00")

func staticTemplatesRedirectHtmlBytes() ([]byte, error) {
	return bindataRead(
		_staticTemplatesRedirectHtml,
		"static/templates/redirect.html",
	)
}

func staticTemplatesRedirectHtml() (*asset, error) {
	bytes, err := staticTemplatesRedirectHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/templates/redirect.html", size: 1074, mode: os.FileMode(420), modTime: time.Unix(1517467783, 0)}
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
	"static/migrations/demo/data.sql": staticMigrationsDemoDataSql,
	"static/templates/error.html": staticTemplatesErrorHtml,
	"static/templates/redirect.html": staticTemplatesRedirectHtml,
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
			"demo": &bintree{nil, map[string]*bintree{
				"data.sql": &bintree{staticMigrationsDemoDataSql, map[string]*bintree{}},
			}},
		}},
		"templates": &bintree{nil, map[string]*bintree{
			"error.html": &bintree{staticTemplatesErrorHtml, map[string]*bintree{}},
			"redirect.html": &bintree{staticTemplatesRedirectHtml, map[string]*bintree{}},
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

