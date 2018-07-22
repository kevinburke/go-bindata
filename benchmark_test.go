package bindata_test

import (
	"os"
	"path/filepath"
	"testing"

	bindata "github.com/kevinburke/go-bindata"
)

func dirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}

func BenchmarkBindata(b *testing.B) {
	size, err := dirSize("testdata/benchmark")
	if err != nil {
		b.Fatal(err)
	}
	b.SetBytes(size)
	cfg := &bindata.Config{
		Package: "assets",
		Input: []bindata.InputConfig{
			{Path: "testdata/benchmark", Recursive: true},
		},
		Output:     "testdata/assets/bindata.go",
		NoMemCopy:  false,
		NoCompress: true,
		Debug:      false,
		NoMetadata: true,
		Mode:       0x0, ModTime: 0,
		Ignore: nil,
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := bindata.Translate(cfg); err != nil {
			b.Fatal(err)
		}
	}
}
