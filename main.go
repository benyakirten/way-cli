package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Usage = customUsage
	d, t, w := parseFlags()
	fmt.Printf("d=%d, t=%d, w=%t", d, t, w)
}

func parseFlags() (int, int, bool) {
	d := flag.Int("d", -1, "Set Depth limit. Defaults to -1 (infinite)")
	t := flag.Int("t", -1, "Set time limit in milliseconds. Defaults to -1 (infinite)")
	w := flag.Bool("w", false, "Limit search to only subdirectory")
	flag.Parse()

	return *d, *t, *w
}

func customUsage() {
	fmt.Println("Way is a CLI tool to locate a named folder on a local file system.")
}
