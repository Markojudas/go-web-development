# Understanding the exmaple from pkg.go.dev/io#copy
<!-- markdownlint-disable -->

The examples provided by the standard library uses the <strong>strings package</strong>, the <strong>NewReader</strong> function to create a Reader.

from the strings package:

```go
func NewReader(s string) *Reader
```

This is very simple. We pass a string to the strings.NewReader function and we are get in return a pointer to a Reader!