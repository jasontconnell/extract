package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/jasontconnell/extract/process"
)

func main() {
	dir := flag.String("d", "", "directory to search")
	ext := flag.String("e", "txt", "extensions to search")
	reg := flag.String("r", ".*", "regular expression")
	out := flag.String("o", "$1", "output format regex")
	flag.Parse()

	fmt.Println("got params", *dir, *ext, *reg, *out)

	fullPath := *dir
	if !filepath.IsAbs(fullPath) {
		wd, _ := os.Getwd()
		fullPath = filepath.Join(wd, fullPath)
	}
	d, err := os.Stat(fullPath)
	if err != nil {
		log.Fatal(err)
	}

	if !d.IsDir() {
		log.Fatal(fullPath, "is not a directory")
	}

	s, err := process.Run(fullPath, *ext, *reg, *out)
	if err != nil {
		log.Fatal(err)
	}

	for _, ss := range s {
		fmt.Println(ss)
	}
}
