# dlogger

Dlogger is a logger for debugging. It outputs log only when `DEBUG` env is specified.

[![CircleCI](https://circleci.com/gh/y-yagi/dlogger.svg?style=svg)](https://circleci.com/gh/y-yagi/dlogger)


## Usage

```go
package main

import (
	"os"

	"github.com/y-yagi/dlogger"
)

func main() {
	logger := dlogger.New(os.Stdout)
	logger.Printf("Debug: %v\n", 1) // Nothing output

	os.Setenv("DEBUG", "1")

	logger = dlogger.New(os.Stdout)
	logger.Printf("Debug: %v\n", 1) // Log output
}
```
