
# IAM-API - Authorize middleware

A project written in Go (Golang) with applying Casbin & Go Gin Framework. It features a Authorize middleware API with performance that is up to 40 times faster ^^!
## Installation

1. To edit casbin policy please update in acl-model directory
2. To change jwt field parse from custToCD to another field edit return f from function subjectFromJWT
```go
claim.Get("custodyCD").(string), nil
```
3. To run project:
```sh
$ go run main.go
```
3. To cross complie project, consider using Go Gox https://github.com/mitchellh/gox