package main

import (
	"fmt"
	"io"
	"os"
)

var fileTemplateStart = `package %s
var %s = ` + "`"

func ConvertFile(in, out, pkg string) (variable string, err error) {
	variable = convertName(in)

	infi, err := os.Open(in)
	if err != nil {
		return "", err
	}

	ofi, err := os.Create(out)
	if err != nil {
		return "", err
	}

	defer ofi.Close()
	defer infi.Close()

	_, err = fmt.Fprintf(ofi, fileTemplateStart, pkg, variable)
	if err != nil {
		return "", err
	}
	_, err = io.Copy(ofi, infi)
	if err != nil {
		return "", err
	}
	_, err = ofi.Write([]byte("`\n"))
	return variable, err
}
