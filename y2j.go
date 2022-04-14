package y2j

import (
	"fmt"
	"io"
	"os"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/compfile"
	"github.com/rwxrob/help"
	y2j "github.com/rwxrob/y2j/pkg"
)

var Cmd = &Z.Cmd{

	Name:      `y2j`,
	Summary:   `convert basic YAML to JSON`,
	Usage:     `[help|<file>]`,
	Comp:      compfile.New(),
	Version:   `v0.4.0`,
	Copyright: `Copyright 2021 Robert S Muhlestein`,
	License:   `Apache-2.0`,
	Commands:  []*Z.Cmd{help.Cmd},

	Description: `
		The y2j command converts YAML (including references and
		anchors) to compressed JSON (with a single training newline) using
		the popular Go yaml.v3 package and its special <yaml:",inline"> tag.
		Because of this only YAML maps are supported as a base type (no
		arrays). An array can easily be done as the value of a map key.

		If no filename is passed standard input is assumed.

		JSON encoding is done using rwxrob/json instead of the "standard"
		encoding/json which unnecessarily escapes HTML characters by
		default.

		If h or help are passed as the first argument this help is assumed
		(preventing the uncommon use of files named 'h' or 'help').

		The intent of y2j is to get YAML into a form quickly for using
		with the more powerful jq command line utility until an equally
		powerful yaml alternative is available.

		`,

	Call: func(_ *Z.Cmd, args ...string) error {
		var buf []byte
		var err error

		if len(args) == 0 {
			buf, err = io.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
		} else {
			buf, err = os.ReadFile(args[0])
			if err != nil {
				return err
			}
		}

		buf, err = y2j.Convert(buf)
		if err != nil {
			return err
		}
		fmt.Println(string(buf))

		return nil
	},
}
