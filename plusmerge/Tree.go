package plusmerge

import (
	"fmt"
)

// Tree
type Tree struct {
	Child *Tree
	Value *[]string
}

// Walk traverses a tree depth-first,
// sending each Value on a channel.
func Walk(t *Tree, ch chan int) {
	if t == nil {
		return
	}
	for _, value := range *t.Value {
		fmt.Println(value)
		// Walk(value, ch)
		// ch <- t.Value
	}
}

// func main() {

//t1 := New(100, 1)
// fmt.Println(Compare(t1, New(100, 1)), "Same Contents")
// fmt.Println(Compare(t1, New(99, 1)), "Differing Sizes")
// fmt.Println(Compare(t1, New(100, 2)), "Differing Values")
// fmt.Println(Compare(t1, New(101, 2)), "Dissimilar")
// }
