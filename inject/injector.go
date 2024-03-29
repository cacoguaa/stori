package inject

import (
	"log"
	"stori/config"
	"stori/repositories"
	"stori/services"
)

type Injector struct {
	Controllers
	Repositories
	Services
}

type Controllers struct {
}

type Repositories struct {
}

type Services struct {
	FileReaderService    *services.FileReaderService
	TransactionsServices *services.TransactionsService
	EmailService         *services.EmailService
}

func NewInjector(config config.Config) *Injector {
	fileReaderService := services.NewFileReaderService()

	db, err := repositories.NewDatabase(config.DBConfig)
	if err != nil {
		log.Fatal("Can't create database connection")
	}

	transactionsService := services.NewTransactionsService(db)
	emailService := services.NewEmailService(config.EmailConfig)

	return &Injector{
		Controllers{},
		Repositories{},
		Services{
			EmailService:         emailService,
			FileReaderService:    fileReaderService,
			TransactionsServices: transactionsService,
		},
	}
}
