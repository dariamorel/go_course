package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)
	if number%2 == 0 {
		fmt.Println("четное")
	} else {
		fmt.Println("нечетное")
	}
}
