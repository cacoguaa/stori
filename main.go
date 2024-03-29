package main

import (
	"log"
	"stori/config"
	"stori/inject"
)

func main() {
	conf := config.Environments()
	injector := inject.NewInjector(conf)

	transactions, err := injector.FileReaderService.ReadFile(conf.Filepath)
	if err != nil {
		log.Fatal("ERROR |  Can't process the file")
	}

	emailBody := injector.TransactionsServices.ProcessTransactions(transactions)

	injector.EmailService.Send("from@mail.com", "password", "target@mail.com", emailBody)
}
