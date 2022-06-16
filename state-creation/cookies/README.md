# type Cookie

<!-- markdownlint-disable -->

```go
type Cookie struct {

    Name    string
    Value   string

    Path        string  // optional
    Domain      string  // optional
    Expires     string  // optional
    RawExpires  string  // for reading cookies only

    // MaxAge=0 means no "Max-Age" attribute specifiec
    // MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0"
    // MaxAge>0 means Max-Age attribute present and given in seconds
    MaxAge      int
    Secure      bool
    HttpOnly    bool
    Raw         string
    Unparsed    []string //Raw text of unparsed attribute-value pairs
}

```

A Cookie represents an HTTP cookie as sent in the Set-Cookie header of an HTTP response or the Cookie header of an HTTP request.

<hr>

<h2>func (*Cookie) String</h2>

```go
func (c *Cookie) String() string
```

String returns the serialization of the cookie for use in a Cookie header (if only Name and Value are set) or a Set-Cookie respons header (if other fields are set).
if c is nil or c.Name is invalid, the empty string is returned.

<hr>

<h2>func SetCookie</h2>

```go
func SetCookie(w ResponseWriter, cookie *Cookie)
```

SetCookie adds a Set-Cookie headerto the provided ResponseWriter's headers. The provided cookie much have a valid Name. Invalid coookies may be silently dropped.

<hr>

<h2>Get a Cookie: func (*Request) Cookie</h2>

```go
func (r *Request) Cookie(name string) (*Cookie, error)
```

Cookie returns the named cookie provided in the request or ErrNoCookie if not found. If multiple cookies match the given name, only one cookie wll be returned.

<hr>

<h2>Deleting a Cookie</h2>

To delete a cookie, set the "MaxAge" field to either Zero or negative number. You can expire a cookie by setting one of thse two fields: Expires or MaxAge Expires sets the exact time when the cookie expires. Expires is Deprecated. MaxAge sets how long the cookie should live (in seconds)
