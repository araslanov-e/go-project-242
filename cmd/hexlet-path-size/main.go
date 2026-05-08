package main

import (
	"code"
	"fmt"
	"os"
)

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	res, err := code.GetPathSize(path)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s\t%s\n", res, path)
}
