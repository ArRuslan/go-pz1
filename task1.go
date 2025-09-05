package main

import (
	"fmt"
)

func panic_(v any) bool {
	panic(v)
	return true
}

/* Скласти програму для визначення відповіді при розв’язанні логічної задачі */
/* (використовувати оператори розгалуження ЗАБОРОНЕНО!!!, програма повинна бути ЛІНІЙНОЮ!!!) */
/* Сума двох перших цифр заданого чотиризначного числа дорівнює сумі двох його останніх цифр */
func main() {
	var num int

	fmt.Print("Enter a number: ")
	_, err := fmt.Scan(&num)
	_ = err != nil && panic_(err)

	_ = (num < 1000 || num > 9999) && panic_("Number must be between 1000 and 9999")

	sum1 := num/1000 + (num/100)%10
	sum2 := (num/10)%10 + num%10

	fmt.Printf("Sum of first two numbers equal to sum of last two numbers: %t\n", sum1 == sum2)
}
