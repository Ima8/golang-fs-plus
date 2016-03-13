package main

import (
	"fmt"
	"os"
)

func main() {
	isFileExist(`F:\DropboxData\Dropbox\golangBook\go.pdf`)
}

func isFileExist(path string) bool {
	// path/to/whatever exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	} else {
		fmt.Println("Hay")
		return true
	}
}

func isFolderExist(path string) bool {
	return true
}

func createFolder(path string) bool {
	return true
}

func moveFile(pathDest string, pathSrc ...string) bool {
	return true
}

func getChild(path string) []string {
	result := []string{}

	return result
}
