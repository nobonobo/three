package main

import (
	"strings"
)

var nameStack []string

func nsPush(s string) {
	nameStack = append(nameStack, s)
}

func nsPop() {
	nameStack = nameStack[:len(nameStack)-1]
}

func nsPath() string {
	return strings.Join(nameStack, ".")
}
