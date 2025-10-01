package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	amount, from, to := getUserInput()

	rates := map[string]map[string]float64{
		"EUR": {"USD": 1.3, "RUB": 116.35},
		"USD": {"EUR": 0.76, "RUB": 89.5},
		"RUB": {"EUR": 1.3, "USD": 0.011},
	}

	result := calculateRate(amount, from, to, &rates)
	fmt.Println(amount, from, " = ", result, to)
}

var currencies = []string{"USD", "EUR", "RUB"}

func getUserInput() (amount float64, from string, to string) {
	from = getCurrency("Введите исходную валюту", currencies)
	amount = getAmount("Введите сумму: ")
	allowCurrencies := slices.DeleteFunc(currencies, func(i string) bool { return i == from })
	to = getCurrency("Введите целевую валюту", allowCurrencies)

	return
}

func getCurrency(text string, currencies []string) string {
	var currency string

	for {
		fmt.Print(text, "(", strings.Join(currencies, ", "), "): ")
		_, err := fmt.Scan(&currency)
		if err == nil && slices.Contains(currencies, currency) {
			break
		}
	}

	return currency
}

func getAmount(text string) float64 {
	var amount float64

	for {
		fmt.Print(text)
		_, err := fmt.Scan(&amount)
		if err == nil {
			break
		}
	}

	return amount
}

func calculateRate(amount float64, from string, to string, rates *map[string]map[string]float64) float64 {
	return amount * (*rates)[from][to]
}
