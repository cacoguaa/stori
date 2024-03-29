package services

import (
	"fmt"
	"log"
	"stori/models"
	"stori/repositories"
	"stori/utils"
	"strconv"
	"strings"
	"time"
)

type TransactionsService struct {
	DBRepository *repositories.PostgresDatabase
}

func NewTransactionsService(database *repositories.PostgresDatabase) *TransactionsService {
	return &TransactionsService{
		DBRepository: database,
	}
}

func (ts *TransactionsService) saveTransaction(transaction models.Transaction) {
	rowsAffected, err := ts.DBRepository.Save(utils.InsertTransaction, transaction.ID, transaction.DATE, transaction.VALUE)
	if err != nil {
		log.Fatalf("Can't insert transactions: %e", err)
	}
	log.Printf("Total affected rows %v", rowsAffected)
}

func (ts *TransactionsService) ProcessTransactions(transactions []models.Transaction) string {
	totalBalance := float64(0)
	totalTransactionsPerMonth := make(map[string]int)
	totalDebitPerMonth := make(map[string]float64)
	totalCreditPerMonth := make(map[string]float64)

	for _, transaction := range transactions {
		ts.saveTransaction(transaction)

		dateMonth, _ := strconv.Atoi(strings.Split(transaction.DATE, "")[0])
		month := time.Month(dateMonth).String()

		if transaction.VALUE <= 0.0 {
			totalDebitPerMonth[month] = transaction.VALUE + totalDebitPerMonth[month]
		} else {
			totalCreditPerMonth[month] = transaction.VALUE + totalCreditPerMonth[month]
		}

		totalTransactionsPerMonth[month]++
		totalBalance += transaction.VALUE
	}

	result := fmt.Sprintf("Total balance is %v\n", totalBalance)

	totalDebit := 0.0
	totalCredit := 0.0
	for month, quantity := range totalTransactionsPerMonth {
		result = result + fmt.Sprintf("Number of transactions in %s: %v\n", month, quantity)
		totalDebit += totalDebitPerMonth[month]
		totalCredit += totalCreditPerMonth[month]
	}

	months := float64(len(totalTransactionsPerMonth))
	result = result + fmt.Sprintf("Average debit amount: %v\n", totalDebit/months)
	result = result + fmt.Sprintf("Average credit amount: %v\n", totalCredit/months)

	return result
}
