package process

import (
	"fmt"
	"regexp"
	"strings"
)

func match(c []contents, regs []*regexp.Regexp, out string) ([]result, error) {
	results := []result{}

	for _, content := range c {
		sm := []interface{}{}
		for _, reg := range regs {
			matches := reg.FindAllStringSubmatch(content.contents, -1)
			if len(matches) == 0 {
				continue
			}

			s := ""
			sep := ""
			if len(matches[0]) > 2 {
				sep = "\n\t"
				s += sep
			}

			for _, rs := range matches[0][1:] {
				s = s + sep + strings.Trim(rs, "\n")
			}

			sm = append(sm, s)
		}

		val := fmt.Sprintf(out, sm...)
		results = append(results, result{filename: content.filename, value: val})
	}

	return results, nil
}
