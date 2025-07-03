package main

import (
	"fmt"
	"math"
)

func getRevenue() float64 {
	var revenue float64
	fmt.Print("Enter the total revenue: ")
	fmt.Scan(&revenue)

	return revenue
}

func getExpenses() float64 {
	var expenses float64

	fmt.Print("Enter the total expenses: ")
	fmt.Scan(&expenses)

	return expenses
}

func getTaxRate() float64 {
	var taxRate float64

	fmt.Print("Enter the tax rate (as a percentage): ")
	fmt.Scan(&taxRate)

	return taxRate / 100.0
}

func calculateEarningsBeforeTax(revenue float64, expenses float64) float64 {
	return math.Round(revenue-expenses) / 100
}

func calculateProfit(ebt float64, taxRate float64) float64 {
	return math.Round(ebt*(1-taxRate)) / 100
}

func calculateRatioEbtToProfit(ebt float64, profit float64) float64 {
	return math.Round(ebt/profit) / 100
}

func printResults(ebt float64, profit float64, ratio float64) {
	fmt.Print("Earnings Before Tax: ")
	formattedEbt := fmt.Sprintf("%.2f", ebt)
	fmt.Println(formattedEbt)

	fmt.Print("Earnings After Tax: ")
	formattedProfit := fmt.Sprintf("%.2f", profit)
	fmt.Println(formattedProfit)

	fmt.Print("Ratio of EBT to Profit: ")
	fromattedRatio := fmt.Sprintf("%.2f", ratio)
	fmt.Println(fromattedRatio)
}

func main() {
	revenue, expenses, taxRate := getRevenue(), getExpenses(), getTaxRate()

	ebt := calculateEarningsBeforeTax(revenue, expenses)
	profit := calculateProfit(ebt, taxRate)
	ratio := calculateRatioEbtToProfit(ebt, profit)

	printResults(ebt, profit, ratio)
}
