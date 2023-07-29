package main

import "fmt"

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}


// func SumAll(numbersToSum ...[]int) []int {
// 	lengtOfNumbers := len(numbersToSum)
// 	sums := make([]int, lengtOfNumbers)
//
// 	for i, numbers := range numbersToSum {
// 		sums[i] = Sum(numbers)
// 	}
//
// 	return sums
// }


func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}

func main() {
	fmt.Println(Sum([]int{2,3,4,5}))	
	fmt.Println(SumAll([]int{1,2}, []int{0,9}))
	fmt.Println(SumAllTails([]int{1,2}, []int{0,9}))
	fmt.Println(SumAllTails([]int{}, []int{0,9}))
}
