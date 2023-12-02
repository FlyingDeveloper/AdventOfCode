package internal

import (
	"fmt"
	"os"
	"strings"
)

func GetInputAsStringArray(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	stringData := fmt.Sprintf("%s", data)
	return strings.Split(stringData, "\n"), nil
}
