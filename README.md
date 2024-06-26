# RomanAPI

RomanAPI is a REST API that returns Roman numerals within a specified range. By performing a GET request on the endpoint localhost:8080/api/v1/romans with query parameters "min" for the lower boundary of the range and "max" for the upper boundary, you can retrieve the corresponding Roman numerals. For example:
```curl 127.0.0.1:8080/api/v1/romans?min=1&max=10```

# Roman Numerals 

The module romanapi/roman contains functions to convert a decimal number, to a Roman numeral, e.g., 4 to 'IV'. The accepted range of decimals is set to 1-3999. The module contains a unit test, testing mostly edge cases. Alternatively, existing packages such as [romannumeral](https://pkg.go.dev/github.com/brandenc40/romannumeral) could be imported to achieve similar functionality.

## Room for Improvement

Instead of restricting the service to Roman numerals, consider generalizing it to support other numeral systems. For instance, extending functionality to generate English spelled numbers from decimals (e.g., converting 123 to 'one hundred and twenty-three') would require configuring the module with a mapping table and input limits.
Using this example, this could look similar to the table below.
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


# Rest API

The REST API is built in the module `romanapi/api` using the [Gin Web Framework](https://gin-gonic.com/). It utilizes the standard configurations of localhost:8080 and serves three endpoints:
- `/` (GET): Home page
- `/api/v1/romans/` (GET): Retrieves Roman numerals within a specified range using query parameters "min" (required) and "max" (required).
- `/api/v1/docs/index.html/` (GET): Retrieves the Swagger 2.0 API specifications.

The inputs are validated using data models, and simple unit tests are implemented to validate the handler functions in this module.


## API Specifications

The API specifications (Swagger 2.0 / OpenAPI) are created, using General API annotations. See [swaggo](https://github.com/swaggo/swag/blob/master/README.md#general-api-info)

The specifications can be created, by running 
```
go install github.com/swaggo/swag/cmd/swag@latest
swag init
```

The docs can be viewed at http://127.0.0.1:8080/api/v1/docs/index.html, following the instructions on [gin-swagger](https://github.com/swaggo/gin-swagger).



# Room for Improvement & Production Readiness

This section discusses possible measured to improve the API and make it production ready.

## Selecting a Hosting Provider

Choose a hosting provider that meets your requirements in terms of scalability, reliability, security, and budget. Common options include cloud service providers like AWS, Google Cloud Platform, Microsoft Azure, or dedicated server providers.

## Setting up Infrastructure

Provision servers or serverless resources to host your API. Configure networking, firewalls, load balancers, and other necessary infrastructure components to ensure high availability and security.

## Deployment: CI/CD

To achieve production readiness, it's crucial to establish a robust deployment pipeline (CI/CD) that automates the deployment process. Currently, the CI pipeline primarily focuses on running unit tests and integration tests. In a productive environment there's a need to extend it to include deployment procedures. The steps consist of
- Expanding the existing CI/CD pipeline to incorporate deployment automation.
- Set up multiple environments like Development, Staging, and Production. Each environment serves a distinct purpose, facilitating proper testing, validation, and deployment of changes.
- Implement a mechanism to manage environment configurations effectively. This involves maintaining separate configuration files or using environment variables to customize settings for each environment.

## Configuring DNS

Assign a domain name to the API and configure DNS settings for routing traffic from the internet. 


## Logging

Implement logging functionality for debugging and monitoring in production.


## Error Handling

Enhance error handling to provide more informative and user-friendly error messages.

## Security

Improve input validation mechanisms and implement authentication (e.g., JWT) for a production-ready API. Consider authorization logic for resource access control. In the case of this simple API this would probably be a binary access to all or none resources, or even no restrictions at all. If there were configuration tables or data to be sent via PUT/POST, the privileges could be split into the roles of a consumer (GET only) a developer (GET and restricted PUT/POST) and an admin (access to all).

## Monitoring and Alerting

Set up monitoring and alerting systems for proactive detection and response to performance issues, errors, and downtime.

## Prevention of DDOS

Implement measures such as rate limiting, traffic filtering, and utilization of a CDN to prevent DDOS attacks.

## Compliance and Regulations

In a production environment, ensuring compliance with relevant regulations and standards is required. In this case, however, the data is of a very insensitive nature, which would simplify this part.

## Documentation

Improve API documentation, especially error message documentation and parameter validation rules.

# Integration Testing

Integration tests are implemented using the Go Testify Suite combined with Go testcontainers [testcontainers-go](https://pkg.go.dev/github.com/testcontainers/testcontainers-go). Make sure to build the docker container before running these tests.



# How to use this repo

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

- [Developing and Testing REST APIs with Golang, Gin, and GitHub Actions](https://hackmd.io/@_Cl3aSMeQf2V5wwqxEyfwg/SkR0yzU5s)
- [Gin Binding in Go: A Tutorial with Examples](https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/)
- [Gin Validation](https://www.convictional.com/blog/gin-validation)
- [API Validation in Gin: Ensuring Data Integrity in Your API](https://dev.to/ankitmalikg/api-validation-in-gin-ensuring-data-integrity-in-your-api-2p40)
- [Gin Binding in Go: A Tutorial with Examples](https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/)
- [gin-swagger Readme](https://github.com/swaggo/gin-swagger?tab=readme-ov-file)
- [Swagger: Declarative Comments Format](https://github.com/swaggo/swag/blob/master/README.md#declarative-comments-format)
- [Developing and Testing REST APIs with Golang, Gin, and GitHub Actions](https://cloudvesna.com/developing-and-testing-rest-apis-with-golang-gin-and-github-actions-75996b3e264a)