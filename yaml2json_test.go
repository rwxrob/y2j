package yaml2json

import (
	"os"
	"testing"
)

func TestCmd_file(t *testing.T) {

	out := `testdata/simple.json`

	// redirect stdout temporarily
	orig := os.Stdout
	defer func() { os.Stdout = orig }()
	file, _ := os.Create(out)
	defer os.Remove(out)
	os.Stdout = file

	Cmd.Call(nil, `testdata/simple.yaml`)
	want := `{"one":1,"other":"<thing>","some":"<thing>","true":true}` + "\n"
	buf, _ := os.ReadFile(out)
	if want != string(buf) {
		t.Errorf("\nwant: %v\ngot:  %v\n", want, string(buf))
	}
}

/*
func TestCmd_error(t *testing.T) {
	// capture the output
	buf := new(bytes.Buffer)
	log.SetFlags(0)
	log.SetOutput(buf)
	defer log.SetFlags(log.Flags())
	defer log.SetOutput(os.Stderr)
}
*/
