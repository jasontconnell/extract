package process

import "regexp"

type contents struct {
	filename string
	contents string
	lines    []string
}

type result struct {
	filename string
	values   []string
}

type Input struct {
	Reg *regexp.Regexp
	Out string
}
