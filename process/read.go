package process

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func read(dir, ext string) ([]contents, error) {
	c := []contents{}
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		fext := filepath.Ext(path)
		if fext != "."+ext {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("error opening file %s. %w", path, err)
		}

		b, err := ioutil.ReadAll(f)
		if err != nil {
			return fmt.Errorf("error reading file %s. %w", path, err)
		}

		lines := strings.Split(string(b), "\n")
		c = append(c, contents{filename: path, contents: string(b), lines: lines})
		return nil
	})

	return c, err
}
