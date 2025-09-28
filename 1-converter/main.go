package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	amount, from, to := getUserInput()
	result := calculateRate(amount, from, to)
	fmt.Println(amount, from, " = ", result, to)
}

var currencies = []string{"USD", "EUR", "RUB"}

var usdRates = map[string]float64{
	"EUR": 1.3,
	"RUB": 89.5,
}

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

func calculateRate(amount float64, from string, to string) float64 {
	var result float64

	switch {
	case from == "USD" && to == "EUR":
		result = amount * usdRates["EUR"]
	case from == "EUR" && to == "USD":
		result = amount / usdRates["EUR"]
	case from == "USD" && to == "RUB":
		result = amount * usdRates["RUB"]
	case from == "RUB" && to == "USD":
		result = amount / usdRates["RUB"]
	case from == "EUR" && to == "RUB":
		result = amount * usdRates["EUR"] * usdRates["RUB"]
	case from == "RUB" && to == "EUR":
		result = amount / (usdRates["EUR"] * usdRates["RUB"])
	}

	return result
}
