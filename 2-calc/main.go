package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func getData() []float64 {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите числа через запятую: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	input = strings.ReplaceAll(input, " ", "")
	parts := strings.Split(input, ",")

	var numbers []float64
	for _, value := range parts {
		// Преобразуем каждую часть в целое число
		// fmt.Printf("value: %T", value)
		number, err := strconv.ParseFloat(value, 64)
		if err != nil {
			fmt.Printf("Ошибка преобразования '%s' к числу\n", value)
			continue
		}
		numbers = append(numbers, number)
	}
	return numbers
}

func calculateResult(operation string) float64 {
	var result float64
	data := getData()
	switch operation {
	case "AVG":
		sum := 0.0
		for _, value := range data {
			sum += value
		}
		result = sum / float64(len(data))
		fmt.Println("Среднее значение введенных чисел равно:", result)
	case "SUM":
		result := 0.0
		for _, value := range data {
			result += value
		}
		fmt.Println("Сумма введенных чисел равна:", result)
	case "MED":
		sort.Float64s(data)
		len_data := len(data)
		if len_data == 0 {
			return 0
		}
		mid := len_data / 2
		fmt.Println(mid)
		if len_data%2 == 1 {
			result = data[mid]
		} else {
			result = (data[mid-1] + data[mid]) / 2
		}
		fmt.Println("Медианное значение введенных чисел равно: ", result)
	}
	return result
}
