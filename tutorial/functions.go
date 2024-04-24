package main

import "fmt"

// переменные
var textArray = [3]string{2: "Я по прежнему здесь!", 1: "Я тоже здесь!", 0: "Я здесь!"}

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
