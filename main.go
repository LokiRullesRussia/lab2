package main

import (
	"fmt"
	"lab1/task1"
	"lab1/task2"
	"lab1/task3"
	"lab1/task4"
	"lab1/task5"
	"lab1/task6"
)

func main() {
	fmt.Println(task1.IsEven(3))
	fmt.Println(task2.IsPositive(-9))
	fmt.Println("Цикл с 1 до 10:")
	task3.OutputNumbers()
	fmt.Println("Длинна строки:")
	fmt.Println(task4.СalculateSum("Написать функцию, которая принимает строку и возвращает ее длину."))
	r := task5.Rectangle{Width: 5, Height: 3}
	fmt.Println("Площадь прямоугольника:")
	fmt.Println(task5.SquareRectangle(r))
	fmt.Println("Среднее значение:")
	fmt.Println(task6.Average(4, 7))
}
