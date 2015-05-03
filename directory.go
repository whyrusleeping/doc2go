package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func ConvertDirectory(in, out, pkg string, writer Writer) error {
	entries, err := ioutil.ReadDir(in)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		p := filepath.Join(in, entry.Name())	
		if entry.IsDir() {
			err = ConvertDirectory(p, filepath.Join(out, entry.Name()))
			if err != nil {
				return err
			}
		} else {
			variable, err := ConvertFile(
				p, filepath.Join(out, fmt.Sprintf("%s.go", entry.Name())), pkg)
			if err != nil {
				return err
			}
			_, err = fmt.Fprintf(ofi, "\t\"%s\": %s,\n", entry.Name(), variable)
			if err != nil {
				return err
			}
		}
	}
	_, err = ofi.Write([]byte("}\n"))
	return err
}
