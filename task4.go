package main

import (
	"fmt"
)

/* Скласти програму з використання циклічної структури */
/* Дано натуральне число n. Знайти суму першої та останньої цифри цього числа */
func main() {
	var n int
	var err error

	fmt.Print("Enter a number N: ")
	_, err = fmt.Scan(&n)
	if err != nil {
		panic(err)
	}

	var first, last int
	last = n % 10

	for n > 0 {
		first = n % 10
		n /= 10
	}

	result := first + last

	fmt.Printf("Sum of first and last digits of number: %d\n", result)
}
