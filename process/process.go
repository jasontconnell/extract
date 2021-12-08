package process

import (
	"strings"
)

func Run(dir string, ext string, inputs []Input) ([]string, error) {

	c, err := read(dir, ext)
	if err != nil {
		return nil, err
	}

	r, err := match(c, inputs)
	if err != nil {
		return nil, err
	}

	sep := strings.Repeat("-", 20)
	s := []string{sep}
	for _, res := range r {
		s = append(s, res.filename)
		for _, val := range res.values {
			s = append(s, val)
		}
		s = append(s, sep)
	}

	return s, nil
}
