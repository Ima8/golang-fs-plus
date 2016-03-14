package plusmerge

import (
	"fmt"
	"os"
)

//PathArgs is a type of Path
type PathArgs struct {
	PathSrc  []string
	PathDesc string
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

//DeleteFolder and sub directory
func DeleteFolder(path string) bool {
	err := os.RemoveAll(path)
	if err != nil {
		fmt.Println("Error can't Remove path: " + path)
		return false
	}
	return true

}

/*
MoveFile is a public func for move mutil file
MoveFile(pathArgs {PathDesc:"/PathDesc"})
MoveFile(PathArgs {PathSrc:[]string{"/pathA","/pathB"},PathDesc:"/pathDesc"})
PathDesc string, pathSrc ...string
*/
//manager
func MoveFile(args ...string) bool {
	if len(args) > 2 { // 1 pathSrc 1 Desc
		fmt.Println(args)
	} else { // N pathSrc 1 Desc

	}

	return false
}

// worker
func moveFile(args PathArgs) bool {
	fmt.Println(args)
	// file, err := os.Open()
	// fi,err :=file.Stat()
	return false
}

//GetChild is a public func to get child in directory
func GetChild(path string) []string {
	result := []string{}

	directory, _ := os.Open(path)

	objects, err := directory.Readdir(-1)

	//fmt.Println(objects)
	if len(objects) > 0 {
		//fmt.Println("hey we have child")
		for _, obj := range objects {
			//sourcefilepointer := source + "/" + obj.Name()
			result = append(result, obj.Name())
			//fmt.Println(obj.Name())
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
