// Напишите программу, вычисляющую диаметр и длину окружности по заданной площади круга. Площадь круга должна вводиться пользователем с клавиатуры.
// s = pr^2

package main

import ("fmt"
		"math")
func main() {
	var s, d float64
	fmt.Print("Введите площадь круга: ")
	fmt.Scanln(&s)
	d = 2 * math.Sqrt(s / math.Pi)
	fmt.Printf("Диаметр круга = %f \nДлина окружности = %f", d, d *math.Pi)
}
