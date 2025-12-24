package input

import (
	"fmt"
	"os"
	"strings"
)

func ReadAll(path string) []string {
	dd, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("invalid path %v: %w", path, err))
	}

	return strings.Split(string(dd), "\n")
}
