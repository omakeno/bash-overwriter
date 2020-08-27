# Bash Overwriter

[![Go Report Card](https://goreportcard.com/badge/github.com/omakeno/bashoverwriter)](https://goreportcard.com/report/github.com/omakeno/bashoverwriter)

Overwrite previous lines on bash terminal.

This implements io.Writer interface.

## Usage

```golang
writer := bashoverwriter.GetBashoverwriter()

fmt.Fprint(writer, "first line")
fmt.Fprint(writer, "second\nline")
fmt.Fprint(writer, "third\n\nline")
```
