# RomanAPI
REST API returning Roman numerals in a specified range. Perform a Get on the endpoint localhost:8080/api/v1/romans with query parameters "min" for the lower boundary of the range and "max" for the upper boundary, e.g.
```curl 127.0.0.1:8080/api/v1/romans?min=1&max=10```

# Description
## Roman Numerals 
The module romanapi/roman contains functions to convert a decimal number, to a Roman numeral, e.g., 4 to 'IV'. The accepted range of decimals is set to 1-3999. The module contains a unit test, testing mostly edge cases. In principle, I could also have imported an existing package such as https://pkg.go.dev/github.com/brandenc40/romannumeral.

### Room for improvement
Instead of restricting the service only to Roman numerals, one could think of generalizing these. A simple extension would be the creation of English spelled numbers from decimals, e.g., 123 to 'one hundred and twenty three'. In order to successfully reach this, the module should be configurable, were a mapping table and the limits for the input can be given. Addintionally, a few minor adaptations in the code would be required. Using the example before, this could look similar to the table below.
```
{
    1: 'one',
    2: 'two',
    ...
    100: 'hundred'
    ...
    1000: 'thousand'
    ...
}
```


## Rest API
The Rest api is built in the module romanapi/api using the Gin Web Framework https://gin-gonic.com/.
It uses the standard configurations, localhost:8080, and served two endpoints:
- the home page '/' (GET)
- the romans endpoint '/api/v1/romans/' (GET) with query parameters "min" (required) and "max" (required), to define the rang of roman numerals to return. 

The inputs are validated using data models.

## API Specifications
The API specifications (Swagger 2.0 / OpenAPI) are created, using General API annotations. https://github.com/swaggo/swag/blob/master/README.md#general-api-info

The specifications can be created, by running 
```
go install github.com/swaggo/swag/cmd/swag@latest
swag init
```

The docs can be viewed at http://127.0.0.1:8080/docs/index.html, following https://github.com/swaggo/gin-swagger.



# How to use it
The service is wrapped in a docker container. To use it, follow these steps, starting in the base directory of this repo.
## Build Docker Image
```docker build -t romanapi .```

## Run Docker 
```docker run --rm -p 8080:8080 romanapi```

## Test API
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

## Use Local Go Module
```
go mod edit -replace $localpackage=../$localpackage
```
```
go mod get $localpackage
```
```
go mod tidy
```

# Useful Reads
- https://hackmd.io/@_Cl3aSMeQf2V5wwqxEyfwg/SkR0yzU5s
- https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/
- https://www.convictional.com/blog/gin-validation
- https://dev.to/ankitmalikg/api-validation-in-gin-ensuring-data-integrity-in-your-api-2p40
- https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/
- https://github.com/swaggo/gin-swagger?tab=readme-ov-file
- https://github.com/swaggo/swag/blob/master/README.md#declarative-comments-format
- https://cloudvesna.com/developing-and-testing-rest-apis-with-golang-gin-and-github-actions-75996b3e264a


