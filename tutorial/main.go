// пакет
package main

// импорт пакета
import "fmt"

// главная функция - точка входа
func main() {
	// вызов функций
	fmt.Println(">>>", 1)
	printText(42)
	printDigit(0)
	printDigit(7)
	fmt.Println("Сумма чисел =", sum(2, 5, 6, 7, 3, 2, 45))

	fmt.Println(">>>", 2)
	// переменная как тип функции
	f1 := printText
	f1(123)
	f1 = printDigit
	f1(6)
}
