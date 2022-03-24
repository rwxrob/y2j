# Convert YAML to JSON with Go 1.18+

![Go Version](https://img.shields.io/github/go-mod/go-version/rwxrob/yaml2json)
[![GoDoc](https://godoc.org/github.com/rwxrob/yaml2json?status.svg)](https://godoc.org/github.com/rwxrob/yaml2json)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)

## Install

This command can be installed as a standalone program or composed into
a Bonzai command tree.

Standalone

```
go install github.com/rwxrob/yaml2json/yaml2json@latest
```

Composed

```go
package cmds

import (
	"github.com/rwxrob/bonzai"
	"github.com/rwxrob/yaml2json"
)

var Cmd = &bonzai.Cmd{
	Name:     `cmds`,
	Commands: []*bonzai.Cmd{help.Cmd, yaml2json.Cmd},
}
```

## Embedded Documentation

All documentation (like manual pages) has been embedded into the source
code of the application. See the source or run the program with help to
access it.
