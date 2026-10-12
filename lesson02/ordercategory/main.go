package main

import "fmt"

func category(amount int64) (string, bool) {
	switch {
	case amount < 0:
		return "", false
	case amount < 10_000:
		return "small", true
	case amount < 50_000:
		return "regular", true
	default:
		return "large", true
	}
}

func main() {
	for _, amount := range []int64{7_500, 12_000, 50_000, -300, 70_000} {
		name, ok := category(amount)
		if !ok {
			fmt.Println("negative order amount")
			return
		}
		fmt.Println(name)
	}
}
