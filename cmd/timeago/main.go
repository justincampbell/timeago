package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/justincampbell/timeago"
)

func main() {
	flag.Usage = usage
	flag.Parse()
	input := flag.Arg(0)

	if input == "help" {
		usage()
		os.Exit(0)
	}

	if input == "" {
		usage()
		os.Exit(1)
	}

	d, err := time.ParseDuration(input)
	if err != nil {
		fmt.Printf("%s\n\n", err)
		usage()
		os.Exit(1)
	}
	fmt.Println(timeago.FromDuration(d))
}

func usage() {
	fmt.Println("github.com/justincampbell/timeago")
	fmt.Println("usage: timeago DURATION")
	fmt.Println("A duration can be in s (seconds), m (minutes), or h (hours)")
	fmt.Println("Examples: 5s 1m30s 4h30m")
}
