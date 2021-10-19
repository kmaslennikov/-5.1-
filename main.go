package main

import "fmt"

func main() {
	fmt.Println("Введите координату X:")
	var x int
	fmt.Scan(&x)

	fmt.Println("Введите координату Y:")
	var y int
	fmt.Scan(&y)

	if x == 0 && y == 0 {
		fmt.Println("Точка лежит в начале координат")
	} else if x == 0 || y == 0 {
		fmt.Println("Точка находится на координатной оси")
	} else if x < 0 && y > 0 {
		fmt.Println("Точка находится в первой четверти")
	} else if x > 0 && y > 0 {
		fmt.Println("Точка находится во второй четверти")
	} else if x > 0 && y < 0 {
		fmt.Println("Точка находится в третьей четверти")
	} else {
		fmt.Println("Точка находится в четвертой четверти")
	}
}
