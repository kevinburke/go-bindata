// Copyright 2013 Dmitry Chestnykh. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bindata

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func ensureFileContains(name, data string) error {
	b, err := os.ReadFile(name)
	if err != nil {
		return err
	}
	if string(b) != data {
		return fmt.Errorf("wrong data in file: expected %s, got %s", data, string(b))
	}
	return nil
}

func tempFileName(count int) string {
	return filepath.Join(os.TempDir(), fmt.Sprintf("safefile-test-%d-%x", count, time.Now().UnixNano()))
}

var testData = "Hello, this is a test data"

func testInTempDir() error {
	name := tempFileName(0)
	defer os.Remove(name)
	f, err := safefileCreate(name, 0666)
	if err != nil {
		return err
	}
	if name != f.OrigName() {
		f.Close()
		return fmt.Errorf("name %q differs from OrigName: %q", name, f.OrigName())
	}
	_, err = io.WriteString(f, testData)
	if err != nil {
		f.Close()
		return err
	}
	err = f.Commit()
	if err != nil {
		f.Close()
		return err
	}
	err = f.Close()
	if err != nil {
		return err
	}
	return ensureFileContains(name, testData)
}

func TestMakeTempName(t *testing.T) {
	// Make sure temp name is random.
	m := make(map[string]bool)
	for i := 0; i < 100; i++ {
		name, err := makeTempName("/tmp", "sf")
		if err != nil {
			t.Fatal(err)
		}
		if m[name] {
			t.Fatal("repeated file name")
		}
		m[name] = true
	}
}

func TestFile(t *testing.T) {
	err := testInTempDir()
	if err != nil {
		t.Fatal(err)
	}
}

func TestWriteFile(t *testing.T) {
	name := tempFileName(1)
	err := safefileWriteFile(name, []byte(testData), 0666)
	if err != nil {
		t.Fatal(err)
	}
	err = ensureFileContains(name, testData)
	if err != nil {
		os.Remove(name)
		t.Fatal(err)
	}
	os.Remove(name)
}

func TestAbandon(t *testing.T) {
	name := tempFileName(2)
	f, err := safefileCreate(name, 0666)
	if err != nil {
		t.Fatal(err)
	}
	err = f.Close()
	if err != nil {
		t.Fatalf("Abandon failed: %s", err)
	}
	// Make sure temporary file doesn't exist.
	_, err = os.Stat(f.Name())
	if err != nil && !os.IsNotExist(err) {
		t.Fatal(err)
	}
}

func TestDoubleCommit(t *testing.T) {
	name := tempFileName(3)
	f, err := safefileCreate(name, 0666)
	if err != nil {
		t.Fatal(err)
	}
	err = f.Commit()
	if err != nil {
		os.Remove(name)
		t.Fatalf("First commit failed: %s", err)
	}
	err = f.Commit()
	if err != errAlreadyCommitted {
		os.Remove(name)
		t.Fatalf("Second commit didn't fail: %s", err)
	}
	err = f.Close()
	if err != nil {
		os.Remove(name)
		t.Fatalf("Close failed: %s", err)
	}
	os.Remove(name)
}

func TestOverwriting(t *testing.T) {
	name := tempFileName(4)
	defer os.Remove(name)

	olddata := "This is old data"
	err := os.WriteFile(name, []byte(olddata), 0600)
	if err != nil {
		t.Fatal(err)
	}

	newdata := "This is new data"
	err = safefileWriteFile(name, []byte(newdata), 0600)
	if err != nil {
		t.Fatal(err)
	}

	err = ensureFileContains(name, newdata)
	if err != nil {
		t.Fatal(err)
	}
}
