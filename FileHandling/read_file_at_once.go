package main

import (
	"fmt"
	"os"
)

func main() {
	b, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b)

	fmt.Println(str)
}