# Bloom Filter in Go

A simple integer-based Bloom filter implementation in Go.

[![Go Report Card](https://goreportcard.com/badge/github.com/lovung/bloomfilter)](https://goreportcard.com/report/github.com/lovung/bloomfilter)
[![GoDoc](https://godoc.org/github.com/lovung/bloomfilter?status.svg)](https://pkg.go.dev/github.com/lovung/bloomfilter)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/lovung/bloomfilter/blob/main/LICENSE)

## Overview

A Bloom filter is a space-efficient probabilistic data structure used to test whether a given element is a member of a set. This implementation is designed for use with integers.

## Installation

To use this Bloom filter in your Go project, you can install it using `go get`:

```shell
go get github.com/lovung/bloomfilter
```

## Usage
Here's a simple example of how to use the Bloom filter:

```go
package main

import (
	"fmt"
	"github.com/yourusername/bloomfilter"
)

func main() {
	var bf bloomfilter.BloomFilter[int]

	// Add some items to the Bloom filter
	itemsToAdd := []int{1, 2, 3, 4, 5}
	for _, item := range itemsToAdd {
		bf.Add(item)
	}

	// Check if an item may exist in the Bloom filter
	itemToCheck := 3
	if bf.MayExist(itemToCheck) {
		fmt.Printf("Item %d may exist in the Bloom filter\n", itemToCheck)
	} else {
		fmt.Printf("Item %d does not exist in the Bloom filter\n", itemToCheck)
	}
}
```

## API Documentation
For detailed documentation, please refer to the GoDoc page.

## Contributing
Contributions are welcome! Please feel free to open issues or pull requests for bug fixes, improvements, or new features.

## License
This project is licensed under the MIT License - see the LICENSE file for details.
