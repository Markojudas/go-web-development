# Handler Interface

``` go
type Handler interface {
  ServeHTTP(ResponseWriter, *Request)
}
```
