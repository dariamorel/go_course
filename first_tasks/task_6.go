package main

import (
	"fmt"
)

func IsVowel(symb string) bool {
	vowels := []string{"a", "e", "i", "o", "u", "а", "е", "ё", "и", "о", "у", "ы", "э", "ю", "я"}
	for _, vowel := range vowels {
		if symb == vowel {
			return true
		}
	}
	return false
}

func main() {
	var symb string
	fmt.Scan(&symb)
	fmt.Println(IsVowel(symb))
}
