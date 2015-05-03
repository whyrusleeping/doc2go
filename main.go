package main

import (
	"flag"
)

func main() {
	input := flag.String("in", "", "specify input directory")
	output := flag.String("out", ".", "specify output directory")
	flag.Parse()

	err := ConvertTree(input, output)
	if err != nil {
		panic(err)
	}
}
