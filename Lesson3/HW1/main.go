// 1. Доработать калькулятор из методички: больше операций и валидации данных (проверка деления на 0, возведение в степень, факториал) + форматирование строки при выводе дробного числа ( 2 знака после точки)

package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, calc float64 
	var symbol string
	a, b, symbol = inputAndValidation()
	if symbol != "error" {
		calc = calcfn(a, b, calc, symbol)
		calc = round(calc)
	}
}

func inputAndValidation() (float64, float64, string) {
	var a, b float64
	var symbol, yn string

	//Check for operators
	fmt.Print("Введите оператор (доступные: +, -, *, /, n!, ^ и V): ")
	fmt.Scanln(&symbol)
	if symbol != "/" && symbol != "+" && symbol != "-" && symbol != "*" && symbol != "n!" && symbol != "^" && symbol != "V" {
		fmt.Println("Неправильный ввод, хотите начать сначала? \ny - да, n - нет: ")
		fmt.Scanln(&yn)
		if yn == "y" {
			symbol = "error"
			a, b, symbol = inputAndValidation()
			return a, b, symbol
		} else {
			fmt.Print("Спасибо, что вопспользовались нашим калькулятором. Приходите еще")
			yn = "n"
			symbol = "error"
			return a, b, symbol
		}
	}

	//Enter the Operands
	fmt.Print("Введите число: ")
	fmt.Scanln(&a)
	if symbol != "n!" {
		fmt.Print("Введите второе число: ")
		fmt.Scanln(&b)
	}

	//Check for validity
	if symbol == "/" && b == 0 {
		fmt.Print("Деление на 0 невозможно, хотите начать сначала? \ny - да, n - нет: ")
	}
	return a, b, symbol
}

func calcfn(a, b, calc float64, symbol string) float64 {
	switch symbol {
	case "*":
		calc = a * b
	case "/":
		calc = a / b
	case "-":
		calc = a - b
	case "+":
		calc = a + b
	case "n!":
		// Сперва посчитал, через цикл, но он отрабатывает только для целых чисел. Долго не мог найти как, но все же нашел.
		calc = math.Gamma(a+1)
	case "^":
		calc = math.Pow(a, b)
	case "V":
		calc = math.Pow(a, 1/b)
	}
	return calc
}

func round(calc float64) float64 {
	calc = math.Round(calc*100) / 100
	var intCalc int = int(calc * 100)

	if calc == float64(intCalc/100) {
		fmt.Printf("Результат: %.f", calc)
	} else if calc*10 == float64(intCalc/10) {
		fmt.Printf("Результат: %.1f", calc)
	} else {
		fmt.Printf("Результат: %.2f", calc)
	}

	return calc
}