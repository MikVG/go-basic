package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	operation := getOperation()
	calculateResult(operation)
}

func getOperation() string {
	var operation string
	fmt.Println("Выберите операцию: AVG - среднее, SUM - сумма, MED - медиана")
	for {
		fmt.Scan(&operation)
		switch operation {
		case "AVG", "SUM", "MED":
			return operation
		default:
			fmt.Println("Выберите правильную операцию")
		}
	}
}

func getData() []int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите числа через запятую: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	input = strings.ReplaceAll(input, " ", "")
	parts := strings.Split(input, ",")

	var numbers []int
	for _, value := range parts {
		// Преобразуем каждую часть в целое число
		number, err := strconv.Atoi(value)
		if err != nil {
			fmt.Printf("Ошибка преобразования '%s' к числу\n", value)
			continue
		}
		numbers = append(numbers, number)
	}
	return numbers
}

func calculateResult(operation string) int {
	var result int
	data := getData()
	switch operation {
	case "AVG":
		sum := 0
		for _, value := range data {
			sum += value
		}
		result = sum / len(data)
		fmt.Println("Среднее значение введенных чисел равно:", result)
	case "SUM":
		result := 0
		for _, value := range data {
			result += value
		}
		fmt.Println("Сумма введенных чисел равна:", result)
	case "MED":
		// fmt.Printf("Медианное значение введенных чисел равно: %d", result)
	}
	return result
}
