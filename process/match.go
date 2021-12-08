package process

import (
	"fmt"
)

func match(c []contents, inputs []Input) ([]result, error) {
	results := []result{}
	values := []string{}
	for _, content := range c {
		for _, line := range content.lines {
			sm := []interface{}{}
			for _, input := range inputs {
				matches := input.Reg.FindAllStringSubmatch(line, -1)
				if len(matches) == 0 {
					continue
				}

				for _, rs := range matches[0][1:] {
					sm = append(sm, rs)
				}

				if len(sm) > 0 {
					val := fmt.Sprintf(input.Out, sm...)
					values = append(values, val)
				}
			}
		}

		results = append(results, result{filename: content.filename, values: values})
	}

	return results, nil
}
