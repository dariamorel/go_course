package main

import "fmt"

func MeanValue(arr []int) float32 {
	sum := 0
	for _, el := range arr {
		sum += el
	}
	return float32(sum) / float32(len(arr))
}

func main() {
	var n int
	fmt.Scan(&n)
	var arr []int = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println(MeanValue(arr))
}
