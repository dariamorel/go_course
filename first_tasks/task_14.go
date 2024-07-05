package main

import "fmt"

func Find(arr []int, element int) bool {
	for _, el := range arr {
		if el == element {
			return true
		}
	}
	return false
}

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	var element int
	fmt.Scan(&element)
	if Find(arr, element) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
