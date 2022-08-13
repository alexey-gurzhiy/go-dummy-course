// 1. Познакомьтесь с алгоритмом сортировки вставками. Напишите приложение, которое принимает на вход набор целых чисел и отдаёт на выходе его же в отсортированном виде. Сохраните код, он понадобится нам в дальнейших уроках.

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// var maxNum, searchNum, verify int
	//  принимаем на вход набор целых чисел
	var numbers string
	var arr []int
	fmt.Println("Введите числа через запятую (Без пробелов): ")
	fmt.Scanln(&numbers)
	splitNumbers := strings.Split(numbers, ",")

	for i, element := range splitNumbers {
		intVar, err := strconv.Atoi(element)
		// не понимаю, какой тут может быть эрор и как его применять.
		if i < 0 {
			fmt.Print(err)
		}
		arr = append(arr, intVar)
	}

	// начинаем сортировку
	for i, element := range arr {
		arr = insertionSort(i, element, arr)
	}
	fmt.Println(arr)
}

func insertionSort(i, element int, arr []int) []int {
	if i != 0 {
		if arr[i-1] > element {
			arr[i], arr[i-1] = arr[i-1], element
			arr = insertionSort(i-1, arr[i-1], arr)
		}
	}
	return arr
}
