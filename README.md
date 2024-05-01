# RomanAPI
REST API returning roman numerals in a specified range

# Useful commands
Run these in the ./src folder

## Initialize Module
```go mod init $modulename```

## Run Test
```go test```

## Get Dependency on Gin (for quick RestAPI)
```go get -u github.com/gin-gonic/gin```

# Useful Reads
- https://hackmd.io/@_Cl3aSMeQf2V5wwqxEyfwg/SkR0yzU5s

# Docker Commands
## Build Docker Image
```docker build -t romanapi .```

## Run Docker 
```docker run --rm -p 8080:8080 romanapi```

# Test API
```curl 127.0.0.1:8080```