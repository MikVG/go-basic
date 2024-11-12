package main

import (
	"fmt"
)

// const usdToEur = 0.92
// const usdToRub = 97.0
// const eurToRub = usdToEur * usdToRub

type converter map[string]float64

func main() {
	convCurrency := converter{}
	convCurrency["usdToEur"] = 0.92
	convCurrency["usdToRub"] = 97.0

	userName := getUserName()
	fmt.Println("Привет,", userName)
	originalCurrency, originalAmount, targetCurrency := getUserInput()

	// targetAmount := calculateCurrency(originalCurrency, originalAmount, targetCurrency, &convCurrency)
	calculateCurrency(originalCurrency, originalAmount, targetCurrency, &convCurrency)

	// fmt.Printf("Целевая сумма: %.2f %s\n", *targetAmount, targetCurrency)
}

func getUserName() string {
	var userName string
	fmt.Println("Как вас зовут?")
	fmt.Scan(&userName)
	return userName
}

func getUserInput() (string, float64, string) {
	var originalCurrency string
	var originalAmount float64
	var targetCurrency string
	for {
		fmt.Println("Введите исходную валюту (USD или EUR): ")
		fmt.Scan(&originalCurrency)
		if originalCurrency == "USD" || originalCurrency == "EUR" {
			break
		} else {
			fmt.Println("Некорректная исходная валюта")
		}
	}
	for {
		fmt.Println("Введите исходную сумму: ")
		_, err := fmt.Scan(&originalAmount)
		if err != nil {
			fmt.Println("Ошибка ввода, попробуйте еще раз")
			continue
		}
		if originalAmount > 0 {
			break
		} else {
			fmt.Println("Некорректная сумма")
		}
	}
	for {
		var helpCurrency string
		if originalCurrency == "USD" {
			helpCurrency = "EUR"
		} else {
			helpCurrency = "USD"
		}
		fmt.Printf("Введите целевую валюту (%v): \n", helpCurrency)
		fmt.Scan(&targetCurrency)
		if (targetCurrency == "USD" || targetCurrency == "EUR") && targetCurrency != originalCurrency {
			break
		} else {
			fmt.Println("Некорректная целевая валюта")
		}
	}

	return originalCurrency, originalAmount, targetCurrency
}

func calculateCurrency(originalCurrency string, originalAmount float64, targetCurrency string, convCurrency *converter) {
	var targetAmount float64
	if originalCurrency == "USD" {
		targetAmount = originalAmount * (*convCurrency)["usdToEur"] //usdToEur
	} else {
		targetAmount = originalAmount / (*convCurrency)["usdToEur"] //usdToEur
	}
	fmt.Printf("Целевая сумма: %.2f %s\n", targetAmount, targetCurrency)
}
