package main

import "fmt"

func Reverse(str string) string {
	var new_str string
	for i := len(str) - 1; i >= 0; i-- {
		new_str += string(str[i])
	}
	return new_str
}

func main() {
	var str string
	fmt.Scan(&str)
	fmt.Println(Reverse(str))
}
