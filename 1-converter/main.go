package main

import "fmt"

func main() {
	const usdToEur = 0.92
	const usdToRub = 97.0
	const eurToRub = usdToEur * usdToRub

	fmt.Println("EUR в RUB:", eurToRub)

}
