# RomanAPI
REST API returning Roman numerals in a specified range. Perform a Get on the endpoint localhost:8080/api/v1/romans with query parameters "min" for the lower boundary of the range and "max" for the upper boundary, e.g.
```curl 127.0.0.1:8080/api/v1/romans?min=1&max=10```

# Useful commands

## Initialize Module
```go mod init $modulename```

## generate go.mod
```module $modulename```

## Run Test
```go test```

## Get Dependency on Gin (for quick RestAPI)
```go get -u github.com/gin-gonic/gin```

# Useful Reads
- https://hackmd.io/@_Cl3aSMeQf2V5wwqxEyfwg/SkR0yzU5s
- https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/
- https://www.convictional.com/blog/gin-validation
- https://dev.to/ankitmalikg/api-validation-in-gin-ensuring-data-integrity-in-your-api-2p40
- https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/
- https://github.com/swaggo/gin-swagger?tab=readme-ov-file
- https://github.com/swaggo/swag/blob/master/README.md#declarative-comments-format
- https://cloudvesna.com/developing-and-testing-rest-apis-with-golang-gin-and-github-actions-75996b3e264a

# Docker Commands
## Build Docker Image
```docker build -t romanapi .```

## Run Docker 
```docker run --rm -p 8080:8080 romanapi```

# Test API
```curl 127.0.0.1:8080```

# Use Local Go Module
```
go mod edit -replace $localpackage=../$localpackage
```
```
go mod get $localpackage
```
```
go mod tidy
```