package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, n := range numbers {
		if n%2 == 0 {
			fmt.Printf("%d is even\n", n)
		} else {
			fmt.Printf("%d is odd\n", n)
		}
	}

	evenOddRecursive(numbers)
}

func evenOddRecursive(numbers []int) {
	if len(numbers) == 0 {
		return
	}

	if numbers[0]%2 == 0 {
		fmt.Printf("%d is even\n", numbers[0])
	} else {
		fmt.Printf("%d is odd\n", numbers[0])
	}

	evenOddRecursive(numbers[1:])
}
