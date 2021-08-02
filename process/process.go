package process

func Run(dir string, ext, reg, out string) ([]string, error) {

	c, err := read(dir, ext)
	if err != nil {
		return nil, err
	}

	r, err := match(c, reg, out)
	if err != nil {
		return nil, err
	}

	s := []string{}
	for _, res := range r {
		s = append(s, res.value)
	}

	return s, nil
}
