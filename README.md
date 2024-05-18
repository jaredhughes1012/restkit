# restkit

Simple, lightweight framework for building JSON-consuming and producing RESTful services. This framework
is fully compatible with the built-in http framework


## Get Started

```
go get -u github.com/jaredhughes1012/restkit
```

## Request Binding

```
type Body struct {
  ReturnStatus int `json:"returnStatus"`
}

func handler(w http.ResponseWriter, r *http.Request) {
  restkit.UseRequestBody[Body](w, r, func(body Body) {
    r.WriteHeader(body.ReturnStatus)
  })
}
```