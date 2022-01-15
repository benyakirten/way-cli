package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path"
)

func main() {
	flag.Usage = customHelp
	l, w, r := parseFlags()

	f := flag.Arg(0)

	res := collectResults(f, l, w, r)
	for _, v := range res {
		fmt.Println(v)
	}
}

func collectResults(f string, l int, w string, r bool) []string {
	var res []string
	fileSystem := os.DirFS(w)

	fs.WalkDir(fileSystem, ".", func(fp string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if l > 0 && len(res) >= l {
			// To exit early
			return errors.New("max results found")
		}

		if d.IsDir() && d.Name() == f {
			if r {
				res = append(res, fp)
			} else {
				res = append(res, path.Join(w, fp))
			}
		}
		return nil
	})
	return res
}

func parseFlags() (int, string, bool) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	l := flag.Int("l", -1, "Set maximum number of results collected. Defaults to -1 (infinite)")
	w := flag.String("w", cwd, "Start search from a particular absolute path. Defaults to current working directory")
	r := flag.Bool("r", false, "Return relative paths. Defaults to false")
	flag.Parse()

	return *l, *w, *r
}

func customHelp() {
	fmt.Println("Way is a CLI tool to locate a named folder on a local file system. The idea is to run this CLI to find a desired folder then just copy the absolute path then run cd <path>")
	fmt.Println("Arguments:")
	fmt.Println("<target folder> -- NOTE: This must be passed after all flags")
	fmt.Println("Ex: way -l 1 invoices")
	fmt.Println("Copy results and then cd into them directly")
	fmt.Println("Possible flags:")
	fmt.Println("-l (int) - Maximum amount of results collected. Default: -1 (infinite)")
	fmt.Println("-w (str) - Starts search in a specific (absolute) path. Default: current working directory")
	fmt.Println("-r (bool) - Gives relative paths instead of absolute paths")
}
