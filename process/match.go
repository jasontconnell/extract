package process

import (
	"fmt"
	"regexp"
)

func match(c []contents, reg, out string) ([]result, error) {
	results := []result{}
	re, err := regexp.Compile(reg)
	if err != nil {
		return nil, fmt.Errorf("couldn't compile regex: %s. %w", reg, err)
	}

	for _, content := range c {
		val := re.ReplaceAllString(content.contents, out)
		res := result{filename: content.filename, value: val}
		results = append(results, res)
	}

	return results, nil
}
