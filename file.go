package main

import (
	"fmt"
	"io"
	"os"
)

var templateStart = `package %s
var %s = ` + "`"

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
