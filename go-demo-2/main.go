package main

import "fmt"

func main() {
	// В цикле спрашиваем ввод транзакций: -10, 10, 40.5
	// Добавлять каждую в массив транзакций
	// Вывести массив
	// !Вывести сумму баланса в консоль

	t := []int{0, 1, 2, 3, 4, 5}
	s := t[1:3]

	fmt.Println(s)

	fmt.Println(s[1])

	s[2] = 4
	fmt.Println(s)

	tr := make([]string, 0, 2)
	fmt.Println(len(tr), cap(tr))
	tr = append(tr, "1")
	fmt.Println(len(tr), cap(tr))
	tr = append(tr, "2")
	fmt.Println(len(tr), cap(tr))
	tr = append(tr, "3")
	fmt.Println(len(tr), cap(tr))
	tr = append(tr, "4")
	fmt.Println(len(tr), cap(tr))
	tr = append(tr, "5")
	fmt.Println(len(tr), cap(tr))
	fmt.Println(tr)

	transactions := []float64{}
	for {
		transaction := scanTransaction()
		if transaction == 0 {
			break
		}
		transactions = append(transactions, transaction)
	}
	fmt.Println(transactions)

	balance := calculateBalace(transactions)
	fmt.Printf("Ваш баланс: %.2f", balance)
}

func scanTransaction() float64 {
	var transaction float64
	fmt.Print("Введите транзакцию (n для выхода): ")
	fmt.Scan(&transaction)
	return transaction
}

func calculateBalace(transactions []float64) float64 {
	balance := 0.0
	for _, value := range transactions {
		balance += value
	}
	return balance
}
