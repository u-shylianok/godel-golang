package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("not enough arguments to run")
		return
	}

	filename := os.Args[1]

	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("failed to open file `%s`: %s\n", filename, err)
		return
	}

	defer f.Close()

	fmt.Printf("file `%s` found\n", filename)
}
