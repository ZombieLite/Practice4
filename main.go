package main

import "fmt"

func tableMultiple(id int) {
	sizeTable := make([][]int, id)

	// Заполнение массива произведениями элементов
	for a := 0; a < id; a++ {
		// Инициализация
		sizeTable[a] = make([]int, id)
		for b := 0; b < id; b++ {
			sizeTable[a][b] = (a + 1) * (b + 1)
		}
	}

	// Вывод значений массива с использованием range
	for _, row := range sizeTable {
		fmt.Println(row)
	}
}

func main() {
	tableMultiple(5)
}
