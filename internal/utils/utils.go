package utils

import (
	"fmt"
	"strings"
)

func GetCmdFromPrompt(prompt string) ([]string, error) {
	if prompt == "" {
		return nil, fmt.Errorf("prompt cannot be empty")
	}
	return strings.Fields(strings.ToLower(prompt)), nil
}
