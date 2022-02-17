# ![Conduit App](golang.png)


[![Build Status](https://travis-ci.org/wangzitian0/golang-gin-realworld-example-app.svg?branch=master)](https://travis-ci.org/wangzitian0/golang-gin-realworld-example-app)
[![codecov](https://codecov.io/gh/wangzitian0/golang-gin-realworld-example-app/branch/master/graph/badge.svg)](https://codecov.io/gh/wangzitian0/golang-gin-realworld-example-app)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/gothinkster/golang-gin-realworld-example-app/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/gothinkster/golang-gin-realworld-example-app?status.svg)](https://godoc.org/github.com/gothinkster/golang-gin-realworld-example-app)

> ### Golang/Gin codebase of Conduit App that adheres to the [RealWorld](https://github.com/gothinkster/realworld) spec and API.


This codebase to demonstrate a fully fledged fullstack application built with **Golang/Gin** including CRUD operations, authentication, routing, pagination, and more.


# Directory structure
```
// Add later
```

# Getting started

## Install Golang

Make sure you have Go 1.13 or higher installed.

https://golang.org/doc/install

## Environment Config

Set-up the standard Go environment variables according to latest guidance (see https://golang.org/doc/install#install).


## Install Dependencies
From the project root, run:
```
go build ./...
go test ./...
go mod tidy
```

## Testing
From the project root, run:
```
go test ./...
```
or
```
go test ./... -cover
```
or
```
go test -v ./... -cover
```
depending on whether you want to see test coverage and how verbose the output you want.

## Todo
- More elegance config
- Test coverage (common & users 100%, article 0%)
- ProtoBuf support
- Code structure optimize (I think some place can use interface)
- Continuous integration (done)
