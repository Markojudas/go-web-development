# Redirects

<!-- markdownlint-disable -->

```go
func Redirect(w ResponseWriter, req *Request, url string, code int)
```

Redirect replies to the request with a redirect to URL, which may be a path relative to the request path.

The provided code should be in the 3xx range and is usually StatusMovedPermanently, StatusFound, or StatusSeeOther.

if the Content-Type header has not been set, Redirect sets it to "text/html; charset=utf-8" and writes a small HTML body. Setting the Content-Type header to any value, including nil, disables that behavior

<hr>
<h2>Note</h2>
On the web, we have a client/server architecture. Clients make requests, and servers write responses to thos clients. The request and respons are both just text that much conform to the rules of HTTP. Both the request & response have a start line, headers, and a body.

The request start line is called the "request line". it consists of:

`request-line = method SP request-target SP HTTP-version CRLF`

The response start line is called the "status line". it consists of:

`status-line = HTTP-version SP status-code SP reason-phrase CRLF`

You can see that the status line wants a status-code and a reason-phrase. For redirects we have several status-codes and reason-phrases from which we can choose, but primarily you should choose either

<ul>
    <li> 301 moved permantely</li>
    <li> 303 see other &emsp;&emsp;&emsp;&emsp;&emsp;&emsp;//changes method to GET </li>
    <li> 307 temporary redirect &emsp;&emsp;//keeps same method</li>
</ul>

<hr>
<h2>HTTP status Codes: 3xx range</h2>

```go
const (
    //...
    StatusMultipleChoices = 300
    StatusMovedPermanently = 301
    StatusFound = 302
    StatusSeeOther = 303
    StatusNotModified = 304
    StatusUseProxy = 305

    StatusTemporaryRedirect = 307
    StatysPermanentRedirect = 308

    //...
)
```
