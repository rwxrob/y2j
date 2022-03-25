package yaml2json

import (
	"os"
	"testing"
)

func TestCmd_file_Map(t *testing.T) {

	out := `testdata/map.json`

	// redirect stdout temporarily
	orig := os.Stdout
	defer func() { os.Stdout = orig }()
	file, _ := os.Create(out)
	defer os.Remove(out)
	os.Stdout = file

	Cmd.Call(nil, `testdata/map.yaml`)
	want := `{"one":1,"other":"<thing>","some":"<thing>","true":true}` + "\n"
	buf, _ := os.ReadFile(out)
	if want != string(buf) {
		t.Errorf("\nwant: %v\ngot:  %v\n", want, string(buf))
	}
}
