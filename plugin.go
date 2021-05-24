package main

import "strings"

func Replace(text string) string {
	return strings.ReplaceAll(text, "world", "planet")
}
