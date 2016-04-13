package plusmerge

import (
	"fmt"
	"hash/fnv"
	"os"

	"github.com/RincLiu/Go-Algorithm/data-structures/stack"
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
			//fmt.Println(pathName)
			if !IsExist(pathName) {
				//fmt.Println(pathName)
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
// change to Tree stu
// func moveFile(pathSrc *[]string, pathDec string) (msg string) {
// 	q := NewQueue(len(*pathSrc))
// 	fullPath := NewQueue(len(*pathSrc) * 2)
// 	// q.Push(&Node{getValueSlice(*pathSrc, 0)})
//
// 	//Scen all of directory
// 	for _, data := range *pathSrc {
// 		var fPath string
// 		q.Push(&Node{data})
// 		for {
// 			fPath = ""
// 			if q.count != 0 {
// 				v := q.Pop()
// 				fPath = fPath + v.Value + "/"
// 				fullPath.Push(&Node{fPath})
// 				//fmt.Println(fullPath.Pop())
// 				child := GetChild(v.Value)
// 				fmt.Println(child)
// 				for _, data := range child {
// 					q.Push(&Node{data})
// 				}
// 			} else {
// 				break
// 			}
// 		}
// 	}
//
// 	// just for test now
// 	count := 0
// 	for {
// 		if fullPath.count != 0 {
// 			fmt.Println(fullPath.Pop())
// 			fmt.Println(count)
// 			count++
// 		} else {
// 			break
// 		}
// 	}
//
// 	return "TRUE"
// }
func moveFile(pathSrc *[]string, pathDec string) (msg string) {
	e := new(entry)
	e.key = "/home"
	fmt.Print(e)
	//Scen all of directory
	for _, data := range *pathSrc {
		//var fPath string
		//q.Push(&Node{data})
		fmt.Println(data)
	}
	return "FAIL"
}

func Visit(path string, f os.FileInfo, err error) error {
	fmt.Printf("Visited: %s\n", path)
	return nil
}

//GetChild is a public func to get child in directory
func GetChild(path string) []string {
	var result = make([]string, 0)

	directory, _ := os.Open(path)

	objects, _ := directory.Readdir(-1)

	//fmt.Println(objects)
	if len(objects) > 0 {
		//fmt.Println("hey we have child")
		for _, obj := range objects {
			//sourcefilepointer := source + "/" + obj.Name()
			result = append(result, obj.Name())
			//fmt.Println(obj.Name())
		}
	} else {
		//initError(err.Error())
	}
	//fmt.Println(result)
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

//LinkedHashMap

const TABLE_SIZE = 128

type entry struct {
	key  string
	next *entry
}

type LinkedHashMap struct {
	table []*entry
}

func (hashMap *LinkedHashMap) Put(key string) {
	hashMap.checkTable()
	hash := getHashValue(key)
	e := hashMap.table[hash]
	if e == nil {
		hashMap.table[hash] = &entry{key, nil}
	} else {
		for e.next != nil {
			if e.key == key {
				return
			}
			e = e.next
		}
		if e.key == key {
		} else {
			e.next = &entry{key, nil}
		}
	}
}

func (hashMap *LinkedHashMap) Remove(key string) {
	hashMap.checkTable()
	hash := getHashValue(key)
	e := hashMap.table[hash]
	if e != nil {
		for e.next != nil {
			if e.key == key {
				e.key = e.next.key
				e.next = e.next.next
				return
			}
			e = e.next
		}
		if e.key == key {
			e = nil
		}
	}
}

func (hashMap *LinkedHashMap) Clear() {
	hashMap.checkTable()
	for _, e := range hashMap.table {
		if e != nil {
			stack := &stack.LinkedStack{}
			stack.Push(e)
			nextEntry := e.next
			for nextEntry != nil {
				stack.Push(nextEntry)
				nextEntry = nextEntry.next
			}
			for stack.Size() > 0 {
				e := convertToEntry(stack.Peek())
				e.key = ""
				e.next = nil
				e = nil
				stack.Pop()
			}
		}
	}
}

func getHashValue(key string) uint32 {
	return hashCode(key) % TABLE_SIZE
}

func hashCode(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
func convertToEntry(x interface{}) *entry {
	if v, ok := x.(*entry); ok {
		return v
	} else {
		panic("Entry convertion exception.")
	}
}

func (hashMap *LinkedHashMap) checkTable() {
	if hashMap.table == nil {
		hashMap.table = []*entry{}
		for i := 0; i < TABLE_SIZE; i++ {
			hashMap.table = append(hashMap.table, nil)
		}
	}
}
