package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var templateStart = `package %s
var %s = ` + "`"

func convertName(name string) string {
	name = strings.Replace(name, "/", "_", -1)
	name = strings.Replace(name, ".", "_", -1)
	name = strings.Replace(name, "-", "_", -1)
	return strings.Title(name)
}

func ConvertFile(in, out, pkg string) error {
	infi, err := os.Open(in)
	if err != nil {
		return err
	}

	ofi, err := os.Create(out)
	if err != nil {
		return err
	}

	defer ofi.Close()
	defer infi.Close()

	fmt.Fprintf(ofi, templateStart, pkg, convertName(in))
	io.Copy(ofi, infi)
	ofi.Write([]byte("`\n"))
	return nil
}

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
