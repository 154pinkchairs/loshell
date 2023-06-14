package cmd

import (
	"fmt"
	"strconv"

	"github.com/samber/lo"
	"github.com/urfave/cli/v3"

	"github.com/154pinkchairs/loshell/internal/iokit"
)

// Subset prints a subset of lines from a file.
func Subset(c *cli.Context) error {
	filePath := c.Args().Get(0)
	offset, err := strconv.Atoi(c.Args().Get(1))
	if err != nil {
		return fmt.Errorf("error parsing offset: %w", err)
	}
	length, err := strconv.ParseUint(c.Args().Get(2), 10, 64)
	if err != nil {
		return fmt.Errorf("error parsing length: %w", err)
	}

	lines, err := iokit.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}
	result := lo.Subset(lines, offset, uint(length))

	for _, line := range result {
		fmt.Println(line)
	}

	return nil
}
