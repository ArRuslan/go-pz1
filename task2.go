package main

import (
	"fmt"
	"math"
)

/* Скласти програму з розгалуженням, використовуючи оператор умовного переходу */
/* Дані три дійсні числа. Звести в квадрат ті з них, значення яких додатні, і в четверту ступінь – від’ємні */
func main() {
	var num1, num2, num3 float64
	var err error

	fmt.Print("Enter a first number: ")
	_, err = fmt.Scan(&num1)
	if err != nil {
		panic(err)
	}

	fmt.Print("Enter a second number: ")
	_, err = fmt.Scan(&num2)
	if err != nil {
		panic(err)
	}

	fmt.Print("Enter a third number: ")
	_, err = fmt.Scan(&num3)
	if err != nil {
		panic(err)
	}

	if num1 > 0 {
		num1 = math.Pow(num1, 2)
	} else {
		num1 = math.Pow(num1, 4)
	}

	if num2 > 0 {
		num2 = math.Pow(num2, 2)
	} else {
		num2 = math.Pow(num2, 4)
	}

	if num3 > 0 {
		num3 = math.Pow(num3, 2)
	} else {
		num3 = math.Pow(num3, 4)
	}

	fmt.Printf("Results: num1=%.2f, num2=%.2f, num3=%.2f\n", num1, num2, num3)

	// Or

	/*nums := [3]float64{}
	  var err error

	  for idx := range nums {
	  	fmt.Printf("Enter a number #%d: ", idx+1)
	  	_, err = fmt.Scan(&nums[idx])
	  	if err != nil {
	  		panic(err)
	  	}
	  }

	  for idx, num := range nums {
	  	if num > 0 {
	  		nums[idx] = math.Pow(num, 2)
	  	} else {
	  		nums[idx] = math.Pow(num, 4)
	  	}
	  }

	  fmt.Printf("Results: num1=%.2f, num2=%.2f, num3=%.2f\n", nums[0], nums[1], nums[2])*/
}
