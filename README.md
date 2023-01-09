# wnc-final

## Quick start
Local development:
```sh
# Run app with migrations
docker compose -f local.yml up
```

Production:
```sh
docker compose -f production.yml up
```

Integration tests (can be run in CI):
```sh
docker compose -f integration_test.yml up --abort-on-container-exit --build --exit-code-from http_v1_integration
```

Unit tests (can be run in CI):
```sh
go test -cover -race $(go list ./... | grep -v /integration_test/)
```

Generate locale files:
```sh
goi18n extract -sourceLanguage=en-US -outdir=./locales/en-US/ -format=yaml ./
```

Generate docs customer app files:
```sh
swag init --exclude ./internal/controller/http/v1/services/employee -o ./docs/v2/customer/ --instanceName customer
```
Generate docs employee app files:
```sh
swag init --exclude ./internal/controller/http/v1/services/customer -o ./docs/v2/employee/ --instanceName employee
```

Convert OpenApi v2 to v3 (yaml):
```sh
# https://github.com/swaggo/swag/issues/386
docker run --rm -v $(PWD)/docs:/work openapitools/openapi-generator-cli:latest-release \
	  generate -i /work/v2/swagger.yaml -o /work/v3 -g openapi-yaml --minimal-update \
    && mv ./docs/v3/openapi/openapi.yaml ./docs/v3/ \
    && rm ./docs/v3/README.md \
    && rm ./docs/v3/.openapi-generator-ignore \
    && rm -rf ./docs/v3/.openapi-generator \
    && rm -rf ./docs/v3/openapi
```

Convert OpenApi v2 to v3 (json):
```sh
# https://github.com/swaggo/swag/issues/386
docker run --rm -v $(PWD)/docs:/work openapitools/openapi-generator-cli:latest-release \
	  generate -i /work/v2/swagger.json -o /work/v3 -g openapi --minimal-update \
    && rm ./docs/v3/README.md \
    && rm ./docs/v3/.openapi-generator-ignore \
    && rm -rf ./docs/v3/.openapi-generator 
```

### Changing between openapi v2/v3:
Change [docs.go](https://github.com/TcMits/wnc-final/blob/master/docs/docs.go)

```go
package docs

import _ "github.com/TcMits/wnc-final/docs/v3" // replace this
// import _ "github.com/TcMits/wnc-final/docs/v2"
```


## Overview

### Web framework
[Iris](https://www.iris-go.com/) is an efficient and well-designed, cross-platform, web framework with robust set of features. Build your own high-performance web applications and APIs powered by unlimited potentials and portability.

### Database - ORM
[ent](https://entgo.io/docs/getting-started/) is a simple, yet powerful entity framework for Go, that makes it easy to build and maintain applications with large data-models and sticks with the following principles:

-   Easily model database schema as a graph structure.
-   Define schema as a programmatic Go code.
-   Static typing based on code generation.
-   Database queries and graph traversals are easy to write.
-   Simple to extend and customize using Go templates.

