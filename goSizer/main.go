package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var unit string = "MB"

func main() {
	executeScan()
}

func executeScan() {
	if len(os.Args) == 2 {
		size := getDirSize(os.Args[1])
		fmt.Printf("the size is: %.2f", size)
		print(" ", unit)
	} else {
		println("only add one argument which has to be the path of the item to scan. Try enclousing the path in quotation marks")
	}
}

func getDirSize(path string) float64 {
	var size int64 = 0
	var fileCount int = 0
	var dirCount int = 0

	readPath := func(_ string, file os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !file.IsDir() {
			size += file.Size()
			fileCount++
		} else {
			dirCount++
		}
		return err
	}

	filepath.Walk(path, readPath)
	finalSize := float64(size) / 1024.0 / 1024.0

	if finalSize >= 1024 {
		finalSize /= 1024
		unit = "GB"
	} else {
		if finalSize < 1 {
			finalSize *= 1024
			unit = "KB"
		}
	}

	if finalSize == 0 {
		println("path not found")
	}
	println(dirCount, "folders - ", fileCount, "files")
	return finalSize
}
