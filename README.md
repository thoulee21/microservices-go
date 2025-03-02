# Golang Microservices Boilerplate - Clean Architecture

[![issues](https://img.shields.io/github/issues/gbrayhan/microservices-go)](https://github.com/gbrayhan/microservices-go/tree/master/.github/ISSUE_TEMPLATE)
[![forks](https://img.shields.io/github/forks/gbrayhan/microservices-go)](https://github.com/gbrayhan/microservices-go/network/members)
[![stars](https://img.shields.io/github/stars/gbrayhan/microservices-go)](https://github.com/gbrayhan/microservices-go/stargazers)
[![license](https://img.shields.io/github/license/gbrayhan/microservices-go)](https://github.com/gbrayhan/microservices-go/tree/master/LICENSE)
[![CodeFactor](https://www.codefactor.io/repository/github/gbrayhan/microservices-go/badge/main)](https://www.codefactor.io/repository/github/gbrayhan/microservices-go/overview/main)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/6c10cc49928447f38952edaab67a94a4)](https://www.codacy.com/gh/gbrayhan/microservices-go/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=gbrayhan/microservices-go&amp;utm_campaign=Badge_Grade)

Example structure to start a microservices project with golang. Using a MySQL database. Using a Hexagonal
Architecture that is a Clean Architecture.

## Table of Contents

- [Features](#features)
- [Commands](#commands)

## Features

- **Golang v1.21**: Stable version of go
- **Framework**: A stable version of [gin-go](https://github.com/gin-gonic/gin)
- **Token Security**: with [JWT](https://jwt.io)
- **SQL databaseSQL**: [MariaDB](https://mariadb.org/) using internal sql package of
  go [sql](https://golang.org/pkg/databaseSQL/sql/)
- **Testing**: unit and integration tests using package of go [testing](https://golang.org/pkg/testing/)
- **API documentation**: with [swaggo](https://github.com/swaggo/swag) @latest version that is a go implementation
  of [swagger](https://swagger.io/)
- **Dependency management**: with [go modules](https://golang.org/ref/mod)
- **Environment variables**: using [viper](https://github.com/spf13/viper)
- **Docker support**
- **Code quality**: with [CodeFactor](https://www.codefactor.io/) and [Codacy](https://www.codacy.com/)
- **Linting**: with [golangci-lint](https://golangci-lint.run/usage/install/) an implementation of a Golang linter

## Commands

### Build and run image of docker

```bash
docker-compose up  --build  -d
```

### Swagger Implementation

To visualize the swagger documentation on local use

http://localhost:8080/v1/swagger/index.html

### Unit test command

```bash
# run recursive test
go test  ./test/unit/...
# clean go test results in cache
go clean -testcache
```

### Lint inspection of go

```bash
golangci-lint run ./...
```
