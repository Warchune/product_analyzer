package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		err := fmt.Errorf("use: product_analyzer <file>")
		fmt.Println(err)
		return
	}
	fmt.Println(os.Args[1])
}
