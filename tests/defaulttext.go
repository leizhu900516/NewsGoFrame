package  main

import (
	"strings"
	"fmt"
)

func main() {
	var text string = "ls -ll"
	result :=strings.Split(text," ")
	fmt.Println(result)
}