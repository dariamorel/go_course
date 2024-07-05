package main

import "fmt"

func SumOfArray(arr []int) int64 {
	var sum int64 = 0
	for _, el := range arr {
		sum += int64(el)
	}
	return sum
}

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println(SumOfArray(arr))
}
