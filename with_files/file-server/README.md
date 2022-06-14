# http.FileServer

<!-- markdownlint-disable -->

<h2>func FileServer</h2>

```go
func FileServer(root FileSystem) Handler
```

The function FileServer returns a handler so we can use it on http.Handle()
the function takes in a FileSystem

<h2>Type FileSystem</h2>

```go
type FileSystem interface {
    Open(name string) (File, error)
}
```

A type that has the method <strong>Open(name string) (File, Error)</strong> will implement this interface:

<h2>type Dir</h2>

```go
type Dir string
```

A Dir implements the FileSystem using the native file system restricted to a specific directory tree.

While the FileSystem.Open method takes '/' - sperated paths, a Dir's string value is a filename on the native file system, not a URL, so it
is separated by filepath.Seperator, which ins't necesarily '/'.

An empty Dir is treated as ".".

```go
func (d Dir) Open(name string) (File, error)
```

<h2>func StripPrefix</h2>

```go
func StripPrefix(prefix string, h Handler) Handler
```

StripPrefix returns a hanlder that serves HTTP requests by removing the given prefix from the request URL's Path (and RawPath if set)
and invoking the handler h. StripPrefix hanldes a request for a path that doesn't begin with prefix by replying with an HTTP 404 not found error.
The prefix MUST MATCH exactly: if the prefix in the request contains escaped characters the reply is also an HTTP 404 not found error
