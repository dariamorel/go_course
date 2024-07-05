package main

import "fmt"

func SimpleNumbers(n int) []int {
	var simple []int = make([]int, n)
	k := 2
	for ; k*k < n; k++ {
		i := k + k
		for ; i < n; i += k {
			simple[i] = 1
		}
	}
	return simple
}

func main() {
	var n int
	fmt.Scan(&n)
	simple := SimpleNumbers(n + 1)
	for i := 2; i < len(simple); i++ {
		if simple[i] == 0 {
			fmt.Println(i)
		}
	}
}
