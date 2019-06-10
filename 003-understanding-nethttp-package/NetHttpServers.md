# Web servers using net/http

### Handler interface
```
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)  // important
}
```
Any type implementing `ServeHTTP(ResponseWriter, *Request)` method belongs to type `Handler`.
