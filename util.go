package main

import (
	"strings"
)

func convertName(name string) string {
	name = strings.Replace(name, "/", "_", -1)
	name = strings.Replace(name, ".", "_", -1)
	name = strings.Replace(name, "-", "_", -1)
	return strings.Title(name)
}
