# lru
A simple implementation of LRU for caching with old data preempted.

[![GoDoc](https://godoc.org/github.com/Kwynto/lru?status.svg)](https://godoc.org/github.com/Kwynto/lru)
[![Go Report Card](https://goreportcard.com/badge/github.com/Kwynto/lru)](https://goreportcard.com/report/github.com/Kwynto/lru)
[![GitHub](https://img.shields.io/github/license/Kwynto/lru)](https://github.com/Kwynto/lru/blob/master/LICENSE)
[![gocover.io](https://gocover.io/_badge/github.com/Kwynto/lru)](https://gocover.io/github.com/Kwynto/lru)


## Contents

- [LRU](#lru)
  - [Contents](#contents)
  - [What is LRU cache](#what-is-lru-cache)
  - [How to connect LRU](#how-to-connect-lru)
  - [How to use LRU](#how-to-use-lru)
  - [About the package  (documentation, testing and benchmarking)](#about-the-package)
  - [About the author](#about-the-author)


## What is LRU cache
LRU, or LRU cache (Least Recently Used) is an algorithm for storing a limited amount of data: information that has not been used for the longest time is forced out of storage. It is used in the organization of the cache.

**[⬆ back to top](#lru)** - **[⬆ back to the chapter](#what-is-lru-cache)**

## How to connect LRU
In your project folder, initialize the Go-module with the command
> go mod init your_app_name

Download and install LRU
> go get github.com/Kwynto/lru

Now you can add the LRU package to your Go-code file, for example in `main.go`
```go
import "github.com/Kwynto/lru"
```

**[⬆ back to top](#lru)** - **[⬆ back to the chapter](#how-to-connect-lru)**

## How to use LRU
To use the LRU package, you need to import it into your code.
```go
import "github.com/Kwynto/lru"
```

Before using the cache, it must be created with the `New()` constructor.  
This function takes one `size` argument to specify the number of elements of any type to store.
```go
myAppCache := lru.New(1000) // Creating a cache for 1000 items.
```

**[⬆ back to top](#lru)** - **[⬆ back to the chapter](#how-to-use-lru)**

## About the package

LRU has a description of its functionality in a `README.md` file and internal documentation.  
LRU is tested and has a performance check.  
You can use the LRU tests and documentation yourself.

Download the LRU project to your computer:
> git clone https://github.com/Kwynto/lru.git

Go to the project folder:
> cd ./lru

**Check out the documentation**

Look at the documentation in two steps.  
First, in the console, run:
> godoc -http=:8080

And then in your web browser navigate to the uri:
> http://localhost:8080

*The `godoc` utility may not be present in your Go build and you may need to install it  
command `go get -v golang.org/x/tools/cmd/godoc`*

You can also use Go's standard functionality to view documentation in the console via `go doc`.  
For example:  
> go doc New

If your IDE is good enough, then the documentation for functions and methods will be available from your code editor.

**Testing**

Run tests:
> go test -v

Run tests showing code coverage:
> go test -cover -v

You can view code coverage in detail in your web browser.  
To do this, you need to sequentially execute two commands in the console:
> go test -coverprofile="coverage.out" -v  
> go tool cover -html="coverage.out"

**Performance**

You can look at code performance tests:
> go test -benchmem -bench="." lru.go lru_test.go

**[⬆ back to top](#lru)** - **[⬆ back to the chapter](#about-the-package)**

## About the author

The author of the project is Constantine Zavezeon (Kwynto).  
You can contact the author by e-mail: kwynto@mail.ru  
The author accepts proposals for participation in open source projects,  
as well as willing to accept job offers.

**[⬆ back to top](#lru)**
