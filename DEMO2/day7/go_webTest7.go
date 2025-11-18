package main

import (
	"fmt"
	"strings"
)

func main() {
	command := "G()(al)"
	str1 := strings.Replace(command, "(al)", "al", -1)
	str2 := strings.Replace(str1, "()", "o", -1)
	fmt.Println(str2)

	return str2
}
