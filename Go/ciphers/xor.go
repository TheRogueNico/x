package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: xor <file> <key>")
		os.Exit(1)
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("reading file:", err)
		os.Exit(1)
	}
	key := byte(os.Args[2][0])

	result := make([]byte, len(data))
	for i, b := range data {
		result[i] = b ^ key
	}

	fmt.Println(string(result))
}
