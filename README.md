# 🌳 Go YAML to JSON Converter

[![GoDoc](https://godoc.org/github.com/rwxrob/y2j?status.svg)](https://godoc.org/github.com/rwxrob/y2j)

## Install

This command can be installed as a standalone program or composed into
a Bonzai command tree.

Standalone

```
go install github.com/rwxrob/y2j/y2j@latest
```

Composed

```go
package cmds

import (
	Z "github.com/rwxrob/bonzai"
	"github.com/rwxrob/y2j"
)

var Cmd = &Z.Cmd{
	Name:     `cmds`,
	Commands: []*Z.Cmd{help.Cmd, y2j.Cmd},
}
```

## Embedded Documentation

All documentation (like manual pages) has been embedded into the source
code of the application. See the source or run the program with help to
access it.

## Design Considerations

* **Nothing but maps.**

  It is extremely rare to encounter a YAML array in the wild. It is
  common, however, to have a map with a single array as it's value.
