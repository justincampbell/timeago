# timeago
[![Build Status](https://travis-ci.org/justincampbell/timeago.svg?branch=master)](https://travis-ci.org/justincampbell/timeago)
[![GoDoc](https://godoc.org/github.com/justincampbell/timeago?status.svg)](https://godoc.org/github.com/justincampbell/timeago)

A port of Rails' `time_ago_in_words` to Golang

## Installation

```
go get github.com/justincampbell/timeago
```

## Examples

### Golang

```go
package main

import (
	"fmt"
	"time"

	"github.com/justincampbell/timeago"
)

func main() {
	var d time.Duration

	d, _ = time.ParseDuration("25s")
	fmt.Println(timeago.FromDuration(d))
	d, _ = time.ParseDuration("55m")
	fmt.Println(timeago.FromDuration(d))
	d, _ = time.ParseDuration("72h")
	fmt.Println(timeago.FromDuration(d))
}
```

```
$ go run main.go
less than a minute
about 1 hour
3 days
```

### Command-line

```
$ go get github.com/justincampbell/timeago/cmd/timeago
$ timeago 25s
less than a minute
$ timeago 55m
about 1 hour
$ timeago 72h
3 days
```
