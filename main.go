package main

import (
	"./plusmerge"
	"fmt"
)

func main() {
	//fmt.Println(plusmerge.GetChild(`./plusmerge/tmp`))
	fmt.Println(plusmerge.MoveFile("F:/", "./plusmerge/tmp/dir2", "./plusmerge/tmp/dir3"))
}
