// 2. Задание для продвинутых (необязательное). Написать приложение, которое ищет все простые числа от 0 до N включительно. Число N должно быть задано из стандартного потока ввода.

package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a uint
	fmt.Print("Введите число: ")
	fmt.Scanln(&a)

	for i := 0; i < int(a)+1; i++ {		
		if big.NewInt(int64(i)).ProbablyPrime(20) {
			fmt.Printf("%d - простое число\n", i)
		}
	}
}
