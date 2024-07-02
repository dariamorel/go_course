package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)
	ans := 1
	for i := 2; i < number+1; i++ {
		ans *= i
	}
	fmt.Println(ans)
}
