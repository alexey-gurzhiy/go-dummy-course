// The package is sorting BubleSort and then making a binary search

// ПУЗЫРЬКОВАЯ СОРТИРОВКА
// Бинарная сортировка

package Sort

import (
	"fmt"
	"math/rand"
	// "os"
	"strconv"
)

func StartSort(maxNum, searchNum int) string {
	var verify int
	var output string
	// Создадим массив значений
	var arr = []int{1}
	for k := 0; k < maxNum-1; k++ {
		verify = maxNum + rand.Intn(arr[k])
		verify = verifyFn(arr, verify)
		arr = append(arr, verify)
	}
	fmt.Println("\nИсходный порядок в массиве: ", arr)

	// Пузырьковая сортировка
	var left = 0
	var right = len(arr) - 1
	var min = right
	var max = left
	for i := 0; i < len(arr)-1; i++ {
		if max == right {
			right = right - 1
			max = left
		}
		if min == left {
			left = left + 1
			min = right
		}
		for j := left; j < right; j++ {
			if arr[j] > arr[j+1] {
				arr[j] = arr[j] + arr[j+1]
				arr[j+1] = arr[j] - arr[j+1]
				arr[j] = arr[j] - arr[j+1]
				if arr[j+1] > arr[max] {
					max = j + 1
				}
				if arr[j] < arr[min] {
					min = j
				}
			} else {
				if arr[j+1] > arr[max] {
					max = j + 1
				}
				if arr[j] < arr[min] {
					min = j
				}
			}
		}
	}
	fmt.Print("Порядок: ", arr, "\n")


	// Бинарный поиск
	left = 0
	right = maxNum - 1
	for z := 0; ; z++ {
		var center = int((right + left) / 2)
		if arr[center] == searchNum {
			fmt.Printf("\nПопытка поиска № %d\nЧисло: %d\n", z+1, arr[center])
			output = "Ваше число записано в слайсе под номером " + strconv.Itoa(center) 
			return (output)
		} else if arr[center] > searchNum {
			fmt.Printf("\nПопытка поиска № %d\nЧисло: %d", z+1, arr[center])
			right = center - 1
		} else {
			fmt.Printf("\nПопытка поиска № %d\nЧисло: %d", z+1, arr[center])
			left = center + 1
		}
		if left > right {
			return output
		}
	} 
}

func verifyFn(arr []int, verify int) int {
	for _, element := range arr {
		if verify == element {
			verify++
			verify = verifyFn(arr, verify)
		}
	}
	return verify
}
