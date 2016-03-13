package plusmerge

import (
	"fmt"
	"os"
)

func main() {
	IsExist(`F:\DropboxData\Dropbox\golangBook`)
}

//IsExist is a public func for check File or Folder is Exist
func IsExist(path string) bool {
	// path/to/whatever exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true

}

//CreateFolder is a public func for Create a new Folder
func CreateFolder(path string) bool {
	// path := "/folder/to/create"
	err := os.MkdirAll(path, 0711)

	if err != nil {
		fmt.Println("Error creating directory")
		fmt.Println(err)
		return false
	}
	return true
}

//MoveFile is a public func for move mutil file
func MoveFile(pathDest string, pathSrc ...string) bool {
	return true
}

func getChild(path string) []string {
	result := []string{}

	return result
}
