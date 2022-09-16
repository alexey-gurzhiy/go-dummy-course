// // 1. Напишите приложение, рекурсивно вычисляющее заданное из стандартного ввода число Фибоначчи.
// // 2. Оптимизируйте приложение за счёт сохранения предыдущих результатов в мапе.** (Оно так же должно быть рекурсивным!)**

package Fibbonachi

import (
	"fmt"
)



func RunOrig(input uint) {
	fmt.Println(fibbonachi(input))
}

func RunOptimized1(input uint) {
	var fiboMap = make(map[uint]uint, input)
	var i uint = 1
	result, fiboMap := fibbonachi2(input, i, fiboMap)
	result = result
}

func RunOptimized2(input uint) {
	var fiboMap2 = make(map[uint]uint, input)
	fiboMap2[0], fiboMap2[1] = 0, 1
	result, fiboMap2 := fibbonachi3(input-1, fiboMap2)
	result = result
}

func RunOptimized3(input uint) {
	var fiboMap3 = make(map[uint]uint, input)
	fiboMap3[0], fiboMap3[1] = 0, 1
	fibbonachi4(input-1, fiboMap3)
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
		fiboMap3[n] = fiboMap3[n-1] + fiboMap3[n-2]
		return fiboMap3[n]
	}
	fiboMap3[n] = fibbonachi4(n-1, fiboMap3) + fiboMap3[n-2]
	return fiboMap3[n]
}
