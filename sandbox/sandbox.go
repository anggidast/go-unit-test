package main

import "fmt"

func PalindromeChecker(word string) bool {
	for i := 0; i < len(word)/2; i++ {
		if word[i] != word[len(word)-i-1] {
			return false
		}
	}
	return true
}

func FibonnaciNumber(fibonnaciLength int) []int {
	results := []int{}

	for i := 0; i < fibonnaciLength; i++ {
		if len(results) < 2 {
			results = append(results, i)
		} else {
			results = append(results, results[i-2]+results[i-1])
		}
	}

	return results
}

func main() {
	fmt.Println(FibonnaciNumber(15))
}
