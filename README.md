# LRU
![Coverage](https://img.shields.io/badge/Coverage-92.3%25-brightgreen)
A simple and fast implementation of LRU for caching with old data preemption and constant access time.

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
  - [Usage example](#usage-example)
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

To store a value in the cache, use the `Store((key any, value any)` method. The key must be an input, such as serialized data or a regular structure. The value must be the result of the input.
```go
myAppCache.Store("input data", "output")
```

When you need a calculation, you can reuse the data via the `Load(key any)` method. The result of executing this function will be the value from the cache and an error message.
```go
value, err := myAppCache.Load("input data")
```

If the cache has the required value of the previous operations, then the error will be equal to `nil` and can be checked, but if there is no value, then you need to calculate and add it to the cache.
```go
if err != nil {
  // There is no cache value, you need to perform calculations and put the result into the cache
  key := "input data"
  value = "output" // there should be a calculation of the correct result
  myAppCache.Store(key, value)
}
```

**[⬆ back to top](#lru)** - **[⬆ back to the chapter](#how-to-use-lru)**

## Usage example

Let's calculate a sufficiently large number of values of the Fibonacci number series, put the results in the cache and get some of them.
```go
package main

import (
	"fmt"

	"github.com/Kwynto/lru"
)

func fiboInternal(n uint, a, b uint) uint {
	// Internal function for use in Fibo(n)
	// This function implements the final recursion.
	if n == 1 {
		return b
	}
	return fiboInternal(n-1, b, a+b)
}

// The Fibo() function is a fast implementation of the Fibonacci number via finite recursion.
func Fibo(n uint) uint {
	if n == 0 {
		return 0
	}
	return fiboInternal(n, 0, 1)
}

func main() {
	cacheFibo := lru.New(100)

	for i := 40; i < 141; i++ {
		value := Fibo(uint(i))
		cacheFibo.Store(i, value)
	}

	fmt.Println("Reading the correct value.")
	value1, err := cacheFibo.Load(140)
	fmt.Printf("Result: %v. Err: %v\n", value1, err)
	fmt.Println("")

	fmt.Println("Reading an invalid value.")
	value2, err := cacheFibo.Load(142)
	if err != nil {
		fmt.Printf("Error handling: %v\n", err)
	} else {
		fmt.Printf("Result: %v.\n", value2)
	}
	fmt.Println("")

	fmt.Println("Writing and reading.")
	if cacheFibo.Store(142, Fibo(142)) {
		value3, err := cacheFibo.Load(142)
		fmt.Printf("Result: %v. Err: %v\n", value3, err)
	} else {
		fmt.Println("Write error.")
	}
	fmt.Println("")

	fmt.Println("Reading a value from the cache and calculating if there is no value, then writing a new value.")
	var result uint
	input := 143
	value4, err := cacheFibo.Load(input)
	if err != nil {
		result = Fibo(uint(input))
		cacheFibo.Store(input, result)
	} else {
		result = value4.(uint)
	}
	fmt.Printf("Result: %v\n", result)
	fmt.Println("")

	fmt.Println("The end.")
}

```

**[⬆ back to top](#lru)** - **[⬆ back to the chapter](#usage-example)**

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
