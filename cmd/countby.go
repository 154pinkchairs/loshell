package cmd

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
	"github.com/urfave/cli/v3"

	"github.com/154pinkchairs/loshell/internal/iokit"
)

// CountBy counts the number of lines that match a predicate.
func CountBy(c *cli.Context) error {
	filePath := c.Args().Get(0)
	predicateName := c.Args().Get(1)
	lines, err := iokit.ReadFile(filePath)
	if err != nil {
		return err
	}

	var mapper func(string) string
	switch predicateName {
	case "empty":
		mapper = func(line string) string {
			if line == "" {
				return "empty"
			} else {
				return "not empty"
			}
		}
	case "contains":
		mapper = func(line string) string {
			if strings.Contains(line, c.Args().Get(2)) {
				return fmt.Sprintf("lines that contain %s", c.Args().Get(2))
			} else {
				return fmt.Sprintf("lines that do not contain %s", c.Args().Get(2))
			}
		}
	default:
		return fmt.Errorf("unknown predicate: %s", predicateName)
	}

	result := lo.CountValuesBy(lines, mapper)

	for k, v := range result {
		fmt.Printf("%s: %d\n", k, v)
	}

	return nil
}
