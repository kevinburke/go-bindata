// Code generated by go-bindata. DO NOT EDIT.
// sources:
// in/a/test.asset (15B)
// in/b/test.asset (15B)
// in/c/test.asset (15B)
// in/test.asset (15B)

package main

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	clErr := gz.Close()
	if clErr != nil {
		return nil, clErr
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _inATestAsset = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd7\x57\x28\x4e\xcc\x2d\xc8\x49\x55\x48\xcb\xcc\x49\xe5\x02\x04\x00\x00\xff\xff\x8a\x82\x8c\x85\x0f\x00\x00\x00"

func inATestAssetBytes() ([]byte, error) {
	return bindataRead(
		_inATestAsset,
		"in/a/test.asset",
	)
}

func inATestAsset() (*asset, error) {
	bytes, err := inATestAssetBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "in/a/test.asset", size: 15, mode: os.FileMode(0644), modTime: time.Unix(1136214245, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x67, 0xe5, 0x1a, 0x66, 0xe5, 0x1, 0x4e, 0x8, 0x46, 0xd8, 0xa6, 0x67, 0x32, 0xa8, 0xa4, 0xf2, 0x95, 0x5a, 0x81, 0xd4, 0xe7, 0x1d, 0x90, 0xfa, 0x78, 0x26, 0x53, 0x53, 0x62, 0xd0, 0x86, 0xd2}}
	return a, nil
}

var _inBTestAsset = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd7\x57\x28\x4e\xcc\x2d\xc8\x49\x55\x48\xcb\xcc\x49\xe5\x02\x04\x00\x00\xff\xff\x8a\x82\x8c\x85\x0f\x00\x00\x00"

func inBTestAssetBytes() ([]byte, error) {
	return bindataRead(
		_inBTestAsset,
		"in/b/test.asset",
	)
}

func inBTestAsset() (*asset, error) {
	bytes, err := inBTestAssetBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "in/b/test.asset", size: 15, mode: os.FileMode(0644), modTime: time.Unix(1136214245, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x67, 0xe5, 0x1a, 0x66, 0xe5, 0x1, 0x4e, 0x8, 0x46, 0xd8, 0xa6, 0x67, 0x32, 0xa8, 0xa4, 0xf2, 0x95, 0x5a, 0x81, 0xd4, 0xe7, 0x1d, 0x90, 0xfa, 0x78, 0x26, 0x53, 0x53, 0x62, 0xd0, 0x86, 0xd2}}
	return a, nil
}

var _inCTestAsset = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd7\x57\x28\x4e\xcc\x2d\xc8\x49\x55\x48\xcb\xcc\x49\xe5\x02\x04\x00\x00\xff\xff\x8a\x82\x8c\x85\x0f\x00\x00\x00"

func inCTestAssetBytes() ([]byte, error) {
	return bindataRead(
		_inCTestAsset,
		"in/c/test.asset",
	)
}

func inCTestAsset() (*asset, error) {
	bytes, err := inCTestAssetBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "in/c/test.asset", size: 15, mode: os.FileMode(0644), modTime: time.Unix(1136214245, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x67, 0xe5, 0x1a, 0x66, 0xe5, 0x1, 0x4e, 0x8, 0x46, 0xd8, 0xa6, 0x67, 0x32, 0xa8, 0xa4, 0xf2, 0x95, 0x5a, 0x81, 0xd4, 0xe7, 0x1d, 0x90, 0xfa, 0x78, 0x26, 0x53, 0x53, 0x62, 0xd0, 0x86, 0xd2}}
	return a, nil
}

var _inTestAsset = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd7\x57\x28\x4e\xcc\x2d\xc8\x49\x55\x48\xcb\xcc\x49\xe5\x02\x04\x00\x00\xff\xff\x8a\x82\x8c\x85\x0f\x00\x00\x00"

func inTestAssetBytes() ([]byte, error) {
	return bindataRead(
		_inTestAsset,
		"in/test.asset",
	)
}

func inTestAsset() (*asset, error) {
	bytes, err := inTestAssetBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "in/test.asset", size: 15, mode: os.FileMode(0644), modTime: time.Unix(1136214245, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x67, 0xe5, 0x1a, 0x66, 0xe5, 0x1, 0x4e, 0x8, 0x46, 0xd8, 0xa6, 0x67, 0x32, 0xa8, 0xa4, 0xf2, 0x95, 0x5a, 0x81, 0xd4, 0xe7, 0x1d, 0x90, 0xfa, 0x78, 0x26, 0x53, 0x53, 0x62, 0xd0, 0x86, 0xd2}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"in/a/test.asset": inATestAsset,
	"in/b/test.asset": inBTestAsset,
	"in/c/test.asset": inCTestAsset,
	"in/test.asset":   inTestAsset,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"in": {nil, map[string]*bintree{
		"a": {nil, map[string]*bintree{
			"test.asset": {inATestAsset, map[string]*bintree{}},
		}},
		"b": {nil, map[string]*bintree{
			"test.asset": {inBTestAsset, map[string]*bintree{}},
		}},
		"c": {nil, map[string]*bintree{
			"test.asset": {inCTestAsset, map[string]*bintree{}},
		}},
		"test.asset": {inTestAsset, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
