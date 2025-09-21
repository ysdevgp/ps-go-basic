package main

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	operation, numbers := getUserInput()
	result := calculateResult(operation, numbers)
	fmt.Println("Результат: ", result)
}

var operations = []string{"AVG", "SUM", "MED"}

func getUserInput() (operation string, numbers []float64) {
	operation = getOperation()
	numbers = getNumbers()

	return operation, numbers
}

func getOperation() string {
	var operation string

	for {
		fmt.Print("Введите операцию ", "(", strings.Join(operations, ", "), "): ")
		_, err := fmt.Scan(&operation)
		operation = strings.ToUpper(operation)
		if err == nil && slices.Contains(operations, operation) {
			break
		}
	}

	return operation
}

func getNumbers() []float64 {
	var inputString string

	for {
		fmt.Print("Введите числа через запятую: ")
		if _, err := fmt.Scan(&inputString); err == nil {
			stringFields := strings.Split(inputString, ",")
			numbers := make([]float64, 0, len(stringFields))
			hasError := false

			for _, s := range stringFields {
				val, err := strconv.ParseFloat(strings.TrimSpace(s), 64)

				if err != nil {
					hasError = true
					break
				}

				numbers = append(numbers, val)
			}

			if !hasError {
				return numbers
			}
		}
	}
}

func calculateResult(operation string, numbers []float64) float64 {
	var result float64

	switch operation {
	case "AVG":
		result = calculateAvg(numbers)
	case "SUM":
		result = calculateSum(numbers)
	case "MED":
		result = calculateMed(numbers)
	}

	return result
}

func calculateAvg(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}

	return calculateSum(numbers) / float64(len(numbers))
}

func calculateSum(numbers []float64) float64 {
	result := 0.0

	for _, value := range numbers {
		result = result + value
	}

	return result
}

func calculateMed(numbers []float64) float64 {
	sort.Float64s(numbers)

	n := len(numbers)
	if n == 0 {
		return 0
	}

	if n%2 != 0 {
		return numbers[n/2]
	} else {
		mid1 := numbers[n/2-1]
		mid2 := numbers[n/2]
		return (mid1 + mid2) / 2.0
	}
}
