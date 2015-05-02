package main

import (
	"flag"
)

func main() {
	input := flag.String("in", "", "specify input file")
	output := flag.String("out", "", "specify output filename")
	pkgname := flag.String("package", "", "specify package name for output")
	flag.Parse()

	err := ConvertFile(*input, *output, *pkgname)
	if err != nil {
		panic(err)
	}
}
