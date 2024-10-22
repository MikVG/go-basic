package main

import "fmt"

const usdToEur = 0.92
const usdToRub = 97.0
const eurToRub = usdToEur * usdToRub

func main() {
	userName := getUserName()
	fmt.Println("Привет,", userName)
	var originalAmount float64
	var originalCurrency string
	var targetCurrency string
	fmt.Println("Введите исходную сумму: ")
	fmt.Scan(&originalAmount)
	fmt.Println("Введите исходную валюту (USD или EUR): ")
	fmt.Scan(&originalCurrency)
	fmt.Println("Введите целевую валюту (USD или EUR): ")
	fmt.Scan(&targetCurrency)
	targetAmount := calculateCurrency(originalAmount, originalCurrency, targetCurrency)
	fmt.Println(targetAmount)

	fmt.Println("EUR в RUB:", eurToRub)
}

func getUserName() string {
	var userName string
	fmt.Println("Как вас зовут?")
	fmt.Scan(&userName)
	return userName
}

func calculateCurrency(originalAmount float64, originalCurrency string, targetCurrency string) float64 {
	return originalAmount
}
