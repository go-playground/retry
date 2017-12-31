Retry library
=============
![Project status](https://img.shields.io/badge/version-1.0.0-green.svg)
[![Build Status](https://travis-ci.org/go-playground/retry.svg?branch=master)](https://travis-ci.org/go-playground/retry)
[![Coverage Status](https://coveralls.io/repos/github/go-playground/retry/badge.svg?branch=master)](https://coveralls.io/github/go-playground/retry?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-playground/retry)](https://goreportcard.com/report/github.com/go-playground/retry)
[![GoDoc](https://godoc.org/github.com/go-playground/retry?status.svg)](https://godoc.org/github.com/go-playground/retry)
![License](https://img.shields.io/dub/l/vibe-d.svg)

Retry library provides a set of standardized common components and abstracts away some code that normally is duplicated.
My motivation is not to reinvent the wheel but to standardize a set of components for reuse in other libraries.

Example
------------
```go
package main

import (
	"errors"
	"fmt"

	"github.com/go-playground/retry"
)

func main() {
	err := retry.Run(5, func() (bool, error) {
		return false, errors.New("ERR")
	},
		func(attempt uint16, err error) {
			fmt.Printf("Attempt: %d Error: %s\n", attempt, err)
		},
	)
	if err != nil {
		panic(err)
	}
}

```

Package Versioning
---------------
I'm jumping on the vendoring bandwagon, you should vendor this package as I will not
be creating different version with gopkg.in like allot of my other libraries.

Why? because my time is spread pretty thin maintaining all of the libraries I have + LIFE,
it is so freeing not to worry about it and will help me keep pouring out bigger and better
things for you the community.

License
------
Distributed under MIT License, please see license file in code for more details.
