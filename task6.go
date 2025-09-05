package main

import (
	"fmt"
	"sort"
	"strings"
)

type Student struct {
	name  string
	group int8
	ses   [5]float64
}

/* Створити програму обробки масиву структур */
/*
1. Описати структуру з ім'ям STUDENT, що містить наступні поля:
 - NAME – прізвище та ініціали;
 - GROUP – номер групи;
 - SES – успішність (масив з п'яти елементів).
2. Написати програму, що виконує наступні дії:
 - введення з клавіатури даних в масиві STUD1, що складається з десяти структур типу STUDENT; записи повинні бути впорядковані за зростанням вмісту поля GROUP;
 - вивести прізвища і номери груп для всіх студентів, включених в масив, якщо середній бал студента більше 4.0;
 - якщо таких студентів немає вивести відповідне повідомлення
*/
func main() {
	var stud1 [10]Student
	var err error

	for i := range stud1 {
		var name string
		var i1, i2 rune

		fmt.Printf("Student #%d name: ", i)
		_, err = fmt.Scanf("%s %c. %c.\n", &name, &i1, &i2)
		if err != nil {
			panic(err)
		}

		stud1[i].name = fmt.Sprintf("%s %c. %c.", name, i1, i2)

		fmt.Printf("Student #%d group: ", i)
		_, err = fmt.Scanf("%d", &stud1[i].group)
		if err != nil {
			panic(err)
		}

		for j := range stud1[i].ses {
			fmt.Printf("Student #%d ses #%d: ", i, j)
			_, err = fmt.Scanf("%f", &stud1[i].ses[j])
			if err != nil {
				panic(err)
			}

			// TODO: validate that number is in range 1..5 ?
		}
	}

	sort.Slice(stud1[:], func(a, b int) bool {
		return stud1[a].group < stud1[b].group
	})

	found := false

	for i, student := range stud1 {
		avg := 0.
		for _, ses := range student.ses {
			avg += ses
		}
		avg /= float64(len(student.ses))

		if avg <= 4. {
			continue
		}

		found = true

		fmt.Printf("Student #%d: %s in group %d\n", i, strings.Split(student.name, " ")[0], student.group)
	}

	if !found {
		fmt.Println("No students with grade >4.0 were found!")
	}
}

/*
Student0 a. b.
6
3
4
4
4
4
Student1 c. d.
6
4
3
2
3
3
Student2 e. f.
4
4
5
2
2
2
Student3 g. h.
8
2
5
5
4
2
Student4 i. j.
7
5
3
5
5
3
Student5 k. l.
7
5
3
5
2
5
Student6 m. n.
9
3
3
3
2
2
Student7 o. p.
3
3
4
4
5
5
Student8 q. r.
10
4
2
2
5
3
Student9 s. t.
7
4
4
3
5
3
*/
