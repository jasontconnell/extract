package process

import (
	"fmt"
	"regexp"
)

func match(c []contents, reg, out string) ([]result, error) {
	results := []result{}
	re, err := regexp.Compile("(?m:" + reg + ")")
	if err != nil {
		return nil, fmt.Errorf("couldn't compile regex: %s. %w", reg, err)
	}

	for _, content := range c {
		matches := re.FindAllStringSubmatch(content.contents, -1)
		if len(matches) == 0 {
			continue
		}
		sm := []interface{}{}
		for _, rs := range matches[0][1:] {
			sm = append(sm, rs)
		}

		val := fmt.Sprintf(out, sm...)
		results = append(results, result{filename: content.filename, value: val})
	}

	return results, nil
}
