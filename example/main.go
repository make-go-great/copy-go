package main

import (
	"fmt"

	"github.com/make-go-great/copy-go"
)

func main() {
	// Copy file
	if err := copy.Copy("copy.go", "copy1.go"); err != nil {
		fmt.Println(err)
		return
	}

	// Copy dir
	if err := copy.Copy("example", "example1"); err != nil {
		fmt.Println(err)
		return
	}

	// Replace file
	if err := copy.Replace("copy1.go", "copy2.go"); err != nil {
		fmt.Println(err)
		return
	}

	// Replace dir
	if err := copy.Replace("example1", "example2"); err != nil {
		fmt.Println(err)
		return
	}
}
