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
MoveFile("/pathA","/pathDesc"})
MoveFile("/pathA","/pathB","/pathDesc"})
PathDesc string, pathSrc ...string
*/
//manager
func MoveFile(args ...string) (result string) {
	lenOfargs := len(args)
	if lenOfargs >= 2 { // 1 pathSrc has 1 Desc
		fmt.Println()
		paths := new(PathArgs)
		paths.PathSrc = args[0 : lenOfargs-1]
		paths.PathDesc = args[lenOfargs-1]
		// check path are Vaild
		for i := range paths.PathSrc {
			pathName := paths.PathSrc[i]
			if !IsExist(pathName) {
				return initError(pathName + " not found")
			}
		}
		pathDecs := paths.PathDesc
		if !IsExist(pathDecs) { // Desc not found
			if !CreateFolder(pathDecs) {
				return initError(pathDecs + " can't make folder or path invaild")
			}
		}

		moveFile(&paths.PathSrc, pathDecs)
	} else {
		result = initError("Args not match in format")
	}
	return result
}

// worker
func moveFile(pathSrc *[]string, pathDec string) (msg string) {
	q := NewQueue(len(*pathSrc))
	q.Push(&Node{getValueSlice(*pathSrc, 0)})

	return ""
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
		initError(err.Error())
	}
	fmt.Println(result)
	return result
}

func initError(s string) string {
	msg := "ERROR : " + s
	//fmt.Println(msg)
	return msg
}

func getValueSlice(datas []string, index int) string {
	for i := range datas {
		if i == index {
			return datas[i]
		}
	}
	return ""
}
