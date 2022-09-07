package bindata

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

var sanitizeTests = []struct {
	in  string
	out string
}{
	{`hello`, "`hello`"},
	{"hello\nworld", "`hello\nworld`"},
	{"`ello", "(\"`\" + `ello`)"},
	{"`a`e`i`o`u`", "(((\"`\" + `a`) + (\"`\" + (`e` + \"`\"))) + ((`i` + (\"`\" + `o`)) + (\"`\" + (`u` + \"`\"))))"},
	{"\xEF\xBB\xBF`s away!", "(\"\\xEF\\xBB\\xBF\" + (\"`\" + `s away!`))"},
}

func TestSanitize(t *testing.T) {
	for _, tt := range sanitizeTests {
		out := []byte(sanitize([]byte(tt.in)))
		if string(out) != tt.out {
			t.Errorf("sanitize(%q):\nhave %q\nwant %q", tt.in, out, tt.out)
		}
	}
}

func TestEncode(t *testing.T) {
	t.Skip("used to test unicode ranges")
	data, err := os.ReadFile("testdata/fa.js")
	if err != nil {
		t.Fatal(err)
	}
	w := new(bytes.Buffer)
	uncompressed_memcopy(w, &Asset{}, bytes.NewReader(data))
	fmt.Println(w.String())
	t.Fail()
}

func TestEmptyFile(t *testing.T) {
	buf := new(bytes.Buffer)
	c := &Config{NoCompress: true, NoMemCopy: false}
	err := writeReleaseAsset(buf, c, &Asset{Func: "hello", Path: "testdata/empty/empty_file"})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(buf.String(), `[]byte("")`) {
		t.Errorf("should have got an empty string, got %s", buf.String())
	}
}
