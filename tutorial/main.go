// пакет
package main

// импорт пакета
import "fmt"

// переменные
var textArray = [3]string{2: "Я по прежнему здесь!", 1: "Я тоже здесь!", 0: "Я здесь!"}

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

func printText(answer uint) {
	// краткое объявление переменной внутри функции
	text3 := "А я в переменной .."
	// цикл по массиву
	for index, value := range textArray {
		fmt.Println(index, value)
	}

	fmt.Println(text3, answer)
}

func printDigit(count uint) {
	if count < 1 {
		fmt.Println(count, "Счётчик меньше 1!")
		return
	}

	// цикл с переключателем
	for i := 0; i < int(count); i++ {
		switch i {
		case 0, 1, 2, 3, 4:
			fmt.Println(i, "меньше 5")
		case 5:
			fmt.Println("Это пять!")
		default:
			fmt.Println(i, "Больше чем 5")
		}
	}
}

func sum(digits ...int) (sum int) {
	for _, digit := range digits {
		sum += digit
	}
	return
}
