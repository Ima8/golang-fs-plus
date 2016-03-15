package main

import (
	"./plusmerge"
	"fmt"
)

func main() {
	//plusmerge.GetChild(`./plusmerge/tmp`)
	fmt.Println(plusmerge.MoveFile("./plusmerge/tmp/dir1", "./plusmerge/tmp/dir2", "./plusmerge/tmp/dir3"))
}
