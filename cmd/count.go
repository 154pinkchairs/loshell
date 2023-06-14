package cmd

import (
	"fmt"

	"github.com/samber/lo"
	"github.com/urfave/cli/v3"

	"github.com/154pinkchairs/loshell/internal/iokit"
)

// Count counts the number of lines that match a predicate.
func Count(c *cli.Context) error {
	// Get the file path from the command line
	filePath := c.Args().Get(0)
	// Open the file
	lines, err := iokit.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	result := lo.CountValues(lines)

	for k, v := range result {
		fmt.Printf("%s: %d\n", k, v)
	}

	return nil
}
