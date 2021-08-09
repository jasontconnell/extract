package process

import (
	"regexp"
	"strings"
)

func Run(dir string, ext string, reg []*regexp.Regexp, out string) ([]string, error) {

	c, err := read(dir, ext)
	if err != nil {
		return nil, err
	}

	r, err := match(c, reg, out)
	if err != nil {
		return nil, err
	}

	sep := strings.Repeat("-", 20)
	s := []string{sep}
	for _, res := range r {
		s = append(s, res.filename)
		s = append(s, res.value)
		s = append(s, sep)
	}

	return s, nil
}
