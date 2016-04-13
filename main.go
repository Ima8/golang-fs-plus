package main

import (
	"flag"
	"fmt"
	"path/filepath"

	"./plusmerge"
)

func main() {
	//fmt.Println(plusmerge.GetChild(`./plusmerge/tmp`))
	flag.Parse()
	root := flag.Arg(0)
	err := filepath.Walk(root, plusmerge.Visit)
	fmt.Printf("filepath.Walk() returned %v\n", err)
	fmt.Println(plusmerge.MoveFile("./plusmerge/tmp/a", "./plusmerge/tmp/dir1", "./plusmerge/tmp/dir3"))
	// e := new(linkedHashmap.entry)
	// e.key = "/home"
	// fmt.Print(e)
}
