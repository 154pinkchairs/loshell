package iokit

import (
	"os"
	"strings"
)

func ReadFile(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	return lines, nil
}
