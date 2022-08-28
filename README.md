# lru
A simple implementation of LRU for caching with old data preempted.

[![GoDoc](https://godoc.org/github.com/Kwynto/lru?status.svg)](https://godoc.org/github.com/Kwynto/lru)
[![Go Report Card](https://goreportcard.com/badge/github.com/Kwynto/lru)](https://goreportcard.com/report/github.com/Kwynto/lru)
[![GitHub](https://img.shields.io/github/license/Kwynto/lru)](https://github.com/Kwynto/lru/blob/master/LICENSE)
[![gocover.io](https://gocover.io/_badge/github.com/Kwynto/lru)](https://gocover.io/github.com/Kwynto/lru)


## Contents

- [GoSession](#gosession)
  - [Contents](#contents)
  - [What are sessions and why are they needed](#what-are-sessions-and-why-are-they-needed)
  - [How to connect GoSession](#how-to-connect-gosession)
  - [How to use GoSession](#how-to-use-gosession)
  - [Examples of using](#examples-of-using)
    - [Example 1](#example-1)
    - [Example 2](#example-2)
    - [Example 3](#example-3)
  - [About the package  (documentation, testing and benchmarking)](#about-the-package)
  - [About the author](#about-the-author)


## What are sessions and why are they needed

**[⬆ back to top](#gosession)** - **[⬆ back to the chapter](#what-are-sessions-and-why-are-they-needed)**

## How to connect GoSession
In your project folder, initialize the Go-module with the command
> go mod init your_app_name

Download and install GoSession
> go get github.com/Kwynto/gosession

Now you can add the GoSession package to your Go-code file, for example in `main.go`
```go
import "github.com/Kwynto/gosession"
```

**[⬆ back to top](#gosession)** - **[⬆ back to the chapter](#how-to-connect-gosession)**

## How to use GoSession
To use the GoSession package, you need to import it into your code.
```go
import "github.com/Kwynto/gosession"
```

All operations for working with sessions must be called from handlers.  
Each time you start working with the session store, you need to call `gosession.Start(w *http.ResponseWriter, r *http.Request)`, since this function returns the identifier of the store and allows you to access the elements of the store through the identifier.
```go
id := gosession.Start(&w, r)
```

**[⬆ back to top](#gosession)** - **[⬆ back to the chapter](#how-to-use-gosession)**

## About the package

GoSession has a description of its functionality in a `README.md` file and internal documentation.  
GoSession is tested and has a performance check.  
You can use the GoSession tests and documentation yourself.

Download the GoSession project to your computer:
> git clone https://github.com/Kwynto/gosession.git

Go to the project folder:
> cd ./gosession

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
> go doc Start

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
> go test -benchmem -bench="." gosession.go gosession_test.go

*The slowest of all functions is `cleaningSessions()`, but this should not scare you, as it is a utility function and is rarely executed. This function does not affect the performance of the entire mechanism, it is only needed to clean up the storage from lost sessions.*

**[⬆ back to top](#gosession)** - **[⬆ back to the chapter](#about-the-package)**

## About the author

The author of the project is Constantine Zavezeon (Kwynto).  
You can contact the author by e-mail: kwynto@mail.ru  
The author accepts proposals for participation in open source projects,  
as well as willing to accept job offers.

**[⬆ back to top](#lru)**
