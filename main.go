package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	result := 0
	if len(os.Args) > 1 && os.Args[1] != "" {
		result = len(strings.Fields(os.Args[1][1 : len(os.Args[1])-1]))
	}
	fmt.Println(result)
}
