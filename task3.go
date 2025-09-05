package main

import (
	"fmt"
)

/* Скласти програму з перемикачем */
/* За номером дня тижня (натуральному числу від 1 до 7) вивести в якості результату кількість занять у Вашій групі в цей день */
func main() {
	var day int
	var err error

	fmt.Print("Enter a day number: ")
	_, err = fmt.Scan(&day)
	if err != nil {
		panic(err)
	}

	// ??? TODO: figure out if values should be just hardcoded
	switch day {
	case 1:
		fmt.Println("2")
		break
	case 2:
		fmt.Println("4")
		break
	case 3:
		fmt.Println("4")
		break
	case 4:
		fmt.Println("4")
		break
	case 5:
		fmt.Println("2")
		break
	case 6:
		fmt.Println("0")
		break
	case 7:
		fmt.Println("0")
		break
	default:
		panic("Invalid day")
	}
}
