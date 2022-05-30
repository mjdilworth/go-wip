package main

import (
	"fmt"
)

func main() {

	fmt.Println("hi")

	bRet := IsUnique("abcd")

	fmt.Println(bRet)
}

// Using map of runes for duplicate detection.
func IsUnique(input string) bool {
	seen := make(map[rune]struct{})
	for _, r := range input {
		if _, ok := seen[r]; ok {
			return false
		} else {
			seen[r] = struct{}{}
		}
	}
	return true
}
