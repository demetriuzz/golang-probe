// пакет
package main

// импорт пакета
import "fmt"

// переменные
var textArray = [3]string{2: "Я по прежнему здесь!", 1: "Я тоже здесь!", 0: "Я здесь!"}
var answer uint = 42

// главная функция - точка входа
func main() {
	printText()
	printDigit()
}

func printText() {
	// краткое объявление переменной внутри функции
	text3 := "А я в переменной .."
	for index, value := range textArray {
		fmt.Println(index, value)
	}

	fmt.Println(text3, answer)
}

func printDigit() {
	for i := 0; i < 10; i++ {
		switch i {
		case 0, 1, 2, 3, 4:
			fmt.Println("меньше 5")
		default:
			fmt.Println("Больше или равно 5")
		}
	}
}
