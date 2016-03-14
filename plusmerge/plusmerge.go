package plusmerge

import (
	"fmt"
	"os"
)

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

//DeleteFolder and sub directory
func DeleteFolder(path string) bool {
	err := os.RemoveAll(path)
	if err != nil {
		fmt.Println("Error can't Remove path: " + path)
		return false
	}
	return true

}

//MoveFile is a public func for move mutil file
func MoveFile(pathDest string, pathSrc ...string) bool {
	return true
}

//GetChild is a public func to get child in directory
func GetChild(path string) []string {
	result := []string{}

	directory, _ := os.Open(path)

	objects, err := directory.Readdir(-1)

	//fmt.Println(objects)
	if len(objects) > 0 {
		fmt.Println("hey")
		for _, obj := range objects {
			//sourcefilepointer := source + "/" + obj.Name()
			result = append(result, obj.Name())
			fmt.Println(obj.Name())
		}
	} else {
		printError(err.Error())
	}
	fmt.Println(result)
	return result
}

func printError(s string) {
	fmt.Println("ERROR :", s)
}
