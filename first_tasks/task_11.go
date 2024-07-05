package main

import "fmt"

func ToFahrenheit(temp float32) float32 {
	return temp*9/5 + 32
}

func main() {
	var temp float32
	fmt.Scan(&temp)
	fmt.Println(ToFahrenheit(temp))
}
