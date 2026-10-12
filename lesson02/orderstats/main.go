package main

import "fmt"

func stats(amounts []int64) (int64, int64, int64, float64, bool) {
	if len(amounts) == 0 {
		return 0, 0, 0, 0, false
	}
	minAmount := amounts[0]
	maxAmount := amounts[0]
	var sum int64
	for _, amount := range amounts {
		if amount < minAmount {
			minAmount = amount
		}
		if amount > maxAmount {
			maxAmount = amount
		}
		sum += amount
	}
	return minAmount, maxAmount, sum, float64(sum) / float64(len(amounts)), true
}

func report(amounts []int64) {
	minAmount, maxAmount, sum, avg, ok := stats(amounts)
	if !ok {
		fmt.Println("no orders")
		return
	}
	fmt.Println(minAmount)
	fmt.Println(maxAmount)
	fmt.Println(sum)
	fmt.Printf("%.2f\n", avg)
}

func main() {
	report([]int64{12000, 7500, 31000, 4500})
	report([]int64{})
}
