package main

import "fmt"

func main() {
	outputText()
}

func outputText() {

	var revenue float64
	var expenses float64
	var taxRate float64

	revenue = getUserInput("Revenue: ")
	expenses = getUserInput("Expenses: ")
	taxRate = getUserInput("Tax Rate: ")

	ebt, profit, ratio := calculateEBTProfitRatio(revenue, expenses, taxRate)

	fmt.Println("EBT :", ebt)
	fmt.Println("Profit:", profit)
	fmt.Println("Ratio:", ratio)
}

func getUserInput(printString string) float64 {
	var userInput float64
	fmt.Print(printString)
	fmt.Scan(&userInput)
	return userInput
}

func calculateEBTProfitRatio(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}
