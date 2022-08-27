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
	input = 600000
	fmt.Println(fibbonachi(input))
	fmt.Println(time.Since(start).Milliseconds())

	start = time.Now()
	var fiboMap = make(map[uint]uint, input)
	var i uint = 1
	input, fiboMap = fibbonachi2(input, i, fiboMap)
	fmt.Printf("Запрошенное число это: %d\n", fiboMap[input])
	fmt.Println(time.Since(start).Milliseconds())

	start = time.Now()
	var fiboMap2 = make(map[uint]uint, input)
	fiboMap2[0], fiboMap2[1] = 0, 1
	input, fiboMap2 = fibbonachi3(input-1, fiboMap2)
	fmt.Printf("Запрошенное число это: %d\n", input)
	fmt.Println(time.Since(start).Milliseconds())

	start = time.Now()
	input = 600000 // переуказал, по тому, что в предыдущей переменной была логика перезаписи значения инпут. Не удалил - что бы сравнивать.
	var fiboMap3 = make(map[uint]uint, input)
	fiboMap3[0], fiboMap3[1] = 0, 1
	input = fibbonachi4(input-1, fiboMap3)
	fmt.Printf("Запрошенное число это: %d\n", input)
	fmt.Println(time.Since(start).Milliseconds())
}

func fibbonachi(n uint) uint {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	return fibbonachi(n-1) + fibbonachi(n-2)
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

func fibbonachi3(n uint, fiboMap2 map[uint]uint) (uint, map[uint]uint) {
	var fib uint
	if _, ok := fiboMap2[n-1]; ok {
		fib = fiboMap2[n-1]
	} else {
		fib, fiboMap2 = fibbonachi3(n-1, fiboMap2)
	}
	fiboMap2[n] = fib + fiboMap2[n-2]
	return fiboMap2[n], fiboMap2
}


func fibbonachi4(n uint, fiboMap3 map[uint]uint) uint {
	if _, ok := fiboMap3[n-1]; ok {
		fiboMap3[n]= fiboMap3[n-1]+ fiboMap3[n-2]
		return fiboMap3[n]
	} 
	fiboMap3[n]= fibbonachi4(n-1, fiboMap3) + fiboMap3[n-2]
	return fiboMap3[n]
}