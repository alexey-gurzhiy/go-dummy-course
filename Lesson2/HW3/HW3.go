// 3. С клавиатуры вводится трехзначное число. Выведите цифры, соответствующие количество сотен, десятков и единиц в этом числе.
package HW3

import (
	"fmt"
)

func HW3() {
	var number int
	fmt.Print("Введите число: ")
	fmt.Scanln(&number)
	fmt.Printf("В вашем тексте \nСотен: %d \nДесятков: %d \nЕдиниц: %d", number/100, (number-number/100*100)/10, number-number/10*10)
}
