// // 1. Напишите приложение, рекурсивно вычисляющее заданное из стандартного ввода число Фибоначчи.
// // 2. Оптимизируйте приложение за счёт сохранения предыдущих результатов в мапе.** (Оно так же должно быть рекурсивным!)**

package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	var input uint
	fmt.Println("Введите номер числа в порядке Фибоначи: ")
	// fmt.Scanln(&input)
	input = 50
	fmt.Println(fibbonachi(input))
	fmt.Println(time.Since(start).Milliseconds())

	start = time.Now()
	var fiboMap = make(map[uint]uint, input)
	var i uint = 1
	input, fiboMap = fibbonachi2(input, i, fiboMap)
	fmt.Printf("Запрошенное число это: %d\n", fiboMap[input])
	fmt.Println(time.Since(start).Milliseconds())

}

func fibbonachi(n uint) uint {
	if n == 1 {
		return 0
	} else if n == 2 {
		return 1
	} else {
		return fibbonachi(n-1) + fibbonachi(n-2)
	}
}

func fibbonachi2(input, i uint, fiboMap map[uint]uint) (uint, map[uint]uint) {
	var operand1, operand2 uint
	if i <= input {
		if i == 1 {
			operand1, operand2 = 0, 0
		} else if i == 2 {
			operand1, operand2 = 0, 1
		} else {
			operand1 = fiboMap[i-1]
			operand2 = fiboMap[i-2]
		}
		fiboMap[i] = operand1 + operand2
		i++
		fibbonachi2(input, i, fiboMap)
	}
	return input, fiboMap
}
