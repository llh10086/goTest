package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Each word consists of lowercase and uppercase letters only"
	lst := strings.Split(str, " ")
	str2 := ""

	for i, v := range lst {
		x := v[0]
		ind := i + 1
		var str1 string
		switch x {
		case 'a', 'e', 'i', 'o', 'u':
			for i := 1; i <= ind; i++ {
				str1 = str1 + "a"
			}
			v = v + "ma" + str1
		case 'A', 'E', 'I', 'O', 'U':
			for i := 1; i <= ind; i++ {
				str1 = str1 + "a"
			}
			v = v + "ma" + str1
		default:
			for i := 1; i <= ind; i++ {
				str1 = str1 + "a"
			}
			v = v[1:] + strings.ToLower(string(x)) + "ma" + str1

		}

		if i+1 == len(lst) {
			str2 = str2 + v
		} else {
			str2 = str2 + v + " "

		}

	}
	fmt.Println(str2)

}
