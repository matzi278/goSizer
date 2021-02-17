package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	executeScan()
}

func executeScan() {
	if len(os.Args) == 2 {
		size := getDirSize(os.Args[1])
		fmt.Printf("the size is:  %.3f mb", size)
	} else {
		println("only add one argument which has to be the path of the item to scan. Try enclousing the path in quotation marks")
	}
}

func getDirSize(path string) float64 {
	var size int64 = 0

	readPath := func(path string, file os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !file.IsDir() {
			size += file.Size()
		}
		return err
	}

	filepath.Walk(path, readPath)
	finalSize := float64(size) / 1024.0 / 1024.0

	if finalSize == 0 {
		println("path not found")
	}
	return finalSize
}
