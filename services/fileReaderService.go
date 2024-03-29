package services

import (
	"encoding/csv"
	"os"
	"stori/models"
	"strconv"
)

type FileReaderService struct {
}

func NewFileReaderService() *FileReaderService {
	return &FileReaderService{}
}

func (frs *FileReaderService) ReadFile(filePath string) ([]models.Transaction, error) {
	transactions, err := frs.readFile(filePath)

	return transactions, err
}

func (frs *FileReaderService) readFile(filePath string) ([]models.Transaction, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var transactions []models.Transaction

	reader := csv.NewReader(file)

	lines := 0
	for {
		record, readerErr := reader.Read()
		if readerErr != nil {
			break
		}

		// Skip the csv headers
		if lines == 0 {
			lines = lines + 1
			continue
		}

		transactions = append(transactions, frs.parseTransaction(record))
		lines = lines + 1
	}

	return transactions, err
}

func (frs *FileReaderService) parseTransaction(record []string) models.Transaction {
	id, _ := strconv.ParseInt(record[0], 10, 64)
	sign := record[2][0:1]
	value, _ := strconv.ParseFloat(record[2][1:], 64)

	if sign == "-" {
		value = -1 * value
	}

	return models.Transaction{
		ID:    id,
		DATE:  record[1],
		VALUE: value,
	}
}
