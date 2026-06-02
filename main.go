package main

import (
	"fmt"
)

func main() {
	var file File
	var path = file.createFile("test")
	fmt.Println(path)
}
