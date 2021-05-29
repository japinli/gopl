// Modify the echo program to print the index and value of each of its
// arguments, one per line.
package main

import (
	"fmt"
	"os"
)

func main() {
	for index, arg := range os.Args[1:] {
		fmt.Println(index, arg)
	}
}
