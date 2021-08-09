package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/jasontconnell/extract/process"
)

func main() {
	dir := flag.String("d", "", "directory to search")
	ext := flag.String("e", "txt", "extensions to search")
	flag.Parse()

	unassigned := unassignedFlags(os.Args[1:])
	if len(unassigned) < 2 {
		log.Fatal("need regular expressions and output format")
	}

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

	regs := []*regexp.Regexp{}
	out := unassigned[len(unassigned)-1]

	for _, rs := range unassigned[:len(unassigned)-1] {
		rx, err := regexp.Compile("(?m:" + rs + ")")
		if err != nil {
			log.Fatal(err)
			break
		}
		regs = append(regs, rx)
	}

	s, err := process.Run(fullPath, *ext, regs, out)
	if err != nil {
		log.Fatal(err)
	}

	for _, ss := range s {
		fmt.Println(ss)
	}
}

func unassignedFlags(args []string) []string {
	flgs := []string{}
	for i := 0; i < len(args); i++ {
		if strings.HasPrefix(args[i], "-") {
			continue
		}

		if i > 0 && strings.HasPrefix(args[i-1], "-") {
			continue
		}

		flgs = append(flgs, args[i])
	}

	return flgs
}
