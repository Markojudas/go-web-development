<!-- markdownlint-disable MD033 -->

# http.ServeMux

<h2>func NewServeMux</h2>

```go
func NewServeMux() *ServeMux
```

NewServeMux allocates and returns a new ServeMux

<h2>func (*ServeMux) Handle</h2>

```go
func (mux *ServeMux) Handle(pattern string, handler Handler)
```

Handle registers the handler for the given pattern. If a handler already exists for pattern, Handle panics

<h2>A Mux is both a ServeMux AND a Handler</h2>

```go
func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request)
```

By having the <strong>ServeHTTP(w ResponseWriter, r *Request)</strong> Method, a TYPE ServeMux is also TYPE Handler as it implements the http.Handler interface
