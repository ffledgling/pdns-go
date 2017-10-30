# About

Swagger based client bindings for PDNS's Authoritative HTTP API for Go.

Generated using https://github.com/go-swagger/go-swagger

# Steps

Install go-swagger: `go get -u github.com/go-swagger/go-swagger/cmd/swagger`

Then:
- `swagger validate https://raw.githubusercontent.com/ffledgling/pdns/5854-swagger-api/docs/http-api/swagger/authoritative-api-swagger.yaml`
- `swagger generate client -f https://raw.githubusercontent.com/ffledgling/pdns/5854-swagger-api/docs/http-api/swagger/authoritative-api-swagger.yaml -A pdns-go`

