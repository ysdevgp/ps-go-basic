package main

import "fmt"

func main() {

}

const usdEurRate = 1.3
const usdRubRate = 89.5
const eurRubRate = usdEurRate * usdRubRate

func getUserInput() (amount float64, from string, to string) {
	fmt.Print("Введите сумму: ")
	fmt.Scan(&amount)

	fmt.Print("Введите исходную валюту: ")
	fmt.Scan(&from)

	fmt.Print("Введите целевую валюту: ")
	fmt.Scan(&to)

	return
}

func calculateRate(amount float64, from string, to string) float64 {

}
