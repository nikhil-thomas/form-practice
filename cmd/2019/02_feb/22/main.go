package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	dirname := "." + string(filepath.Separator)

	d, err := os.Open(dirname)
	if err != nil {
		log.Fatalln(err)
	}

	defer d.Close()

	// If n > 0, Readdir returns at most n FileInfo structures. In this case, if
	// Readdir returns an empty slice, it will return a non-nil error
	// explaining why. At the end of a directory, the error is io.EOF.
	//
	// If n <= 0, Readdir returns all the FileInfo from the directory in
	// a single slice.
	files, err := d.Readdir(-1)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Processing:", dirname)

	for _, f := range files {
		if f.Mode().IsRegular() {
			ext := filepath.Ext(f.Name())
			ext = strings.ToLower(ext)
			if ext == ".png" {
				os.Remove(f.Name())
				fmt.Println("Deleted", f.Name())
			}
		}
	}
}
