package main

import (
	"strings"
)

func CleanInput(text string) []string {
	text = strings.ToLower(text)
	text = strings.TrimSpace(text)

	return strings.Fields(text)
}
