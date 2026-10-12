package main

import "fmt"

func fizzBuzz(n int) string {
	switch {
	case n%15 == 0:
		return "FizzBuzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		return ""
	}
}

func main() {
	for i := range 30 {
		n := i + 1
		word := fizzBuzz(n)
		if word == "" {
			fmt.Println(n)
		} else {
			fmt.Println(word)
		}
	}
}
