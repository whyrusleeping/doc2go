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

	_, err = fmt.Fprintf(ofi, templateStart, pkg, convertName(in))
	if err != nil {
		return err
	}
	_, err = io.Copy(ofi, infi)
	if err != nil {
		return err
	}
	_, err = ofi.Write([]byte("`\n"))
	return err
}
