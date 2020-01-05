package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Программа суммирует введенные числа рекурсивно. Введите пять чисел через пробел")

	// Чтение ввода с консоли и запись в переменную
	myscanner := bufio.NewScanner(os.Stdin)
	myscanner.Scan()
	line := myscanner.Text()

	//преобразование стрингового массива в интовый
	var split = strings.Split(line, " ")
	var intsplit = make([]int, 6) //искусственное ограничение длинны среза т.к. нельзя добавлять элементы в массивы в Go
	var err error
	for i := 0; i < len(split); i++ {
		intsplit[i], err = strconv.Atoi(split[i])
		fmt.Println(err)
	}

	//вызов рекурсивной функции сложения
	recur(intsplit, 0, 0)
}

func recur(arr []int, sum, pos int) {
	if pos < len(arr) {
		sum = sum + pos
		fmt.Println(pos, sum)
		pos++
	}
	recur(arr, sum, pos)
}
