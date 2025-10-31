package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type AggregationResult struct {
	CategoryTotals         map[string]float64
	DailyMaximums          map[string]float64
	TotalTransactions      int
	SuccessfulTransactions int
	FailedTransactions     int
}

type TransactionData struct {
	date     string
	amount   float64
	category string
}

type Transaction struct {
	data TransactionData
}

func (t *Transaction) GetDate() string {
	time.Sleep(10 * time.Millisecond)
	return t.data.date
}

func (t *Transaction) GetAmount() float64 {
	time.Sleep(10 * time.Millisecond)
	return t.data.amount
}

func (t *Transaction) GetCategory() string {
	time.Sleep(10 * time.Millisecond)
	return t.data.category
}

// processTransactions processes transactions concurrently and performs aggregation
func processTransactions(transactionChan <-chan Transaction, workerCount int) AggregationResult {
	// TODO
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	workerCountStr, _ := reader.ReadString('\n')
	workerCountStr = strings.TrimSpace(workerCountStr)
	workerCount, _ := strconv.Atoi(workerCountStr)

	transactionCountStr, _ := reader.ReadString('\n')
	transactionCountStr = strings.TrimSpace(transactionCountStr)
	transactionCount, _ := strconv.Atoi(transactionCountStr)

	transactionChan := make(chan Transaction)

	// Start the transaction generator
	go func() {
		for i := 0; i < transactionCount; i++ {
			transactionInput, _ := reader.ReadString('\n')
			transactionInput = strings.TrimSpace(transactionInput)

			transactionParts := strings.Fields(transactionInput)

			var t TransactionData
			if len(transactionParts) == 3 {
				amount, err := strconv.ParseFloat(transactionParts[1], 64)
				if err != nil {
					amount = -1
				}
				t = TransactionData{
					date:     transactionParts[0],
					amount:   amount,
					category: transactionParts[2],
				}
			} else {
				t = TransactionData{
					date:     "",
					amount:   -1,
					category: "",
				}
			}
			transactionChan <- Transaction{data: t}
		}
		close(transactionChan)
	}()

	// Start time measurement
	startTime := time.Now()

	// Call the student's function
	result := processTransactions(transactionChan, workerCount)

	// Measure elapsed time
	elapsedTime := time.Since(startTime)

	// Set the time limit
	timeLimit := 10 * time.Second

	if elapsedTime > timeLimit {
		fmt.Println("Processing took too long. Ensure you are processing transactions concurrently.")
		return
	}

	// Format and print the results
	formattedCategoryTotals := make(map[string]string)
	for category, total := range result.CategoryTotals {
		formattedCategoryTotals[category] = fmt.Sprintf("%.2f", total)
	}

	formattedDailyMaximums := make(map[string]string)
	for date, maxAmount := range result.DailyMaximums {
		formattedDailyMaximums[date] = fmt.Sprintf("%.2f", maxAmount)
	}

	fmt.Println("Category Totals:", formattedCategoryTotals)
	fmt.Println("Daily Maximums:", formattedDailyMaximums)
	fmt.Println("Total Transactions:", result.TotalTransactions)
	fmt.Println("Successful Transactions:", result.SuccessfulTransactions)
	fmt.Println("Failed Transactions:", result.FailedTransactions)
}
