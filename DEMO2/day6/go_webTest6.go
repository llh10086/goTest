package main

import "fmt"

func main() {
	arr := []int{2, 1}
	k := 1
	sum := 0

	lst := make([]int, 0)
	for i, v := range arr {
		isGood := true

		if i-k >= 0 {

			if arr[i-k] > v {
				isGood = false
			}
		}
		if i+k < len(arr) {
			if arr[i+k] > v {
				isGood = false
			}
		}
		if isGood {
			lst = append(lst, v)
		}

	}
	for _, v := range lst {
		sum += v
	}
	fmt.Println(sum)

}
