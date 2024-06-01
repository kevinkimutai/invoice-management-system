package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kevinkimutai/invoice-management-system/internal/adapters/auth"
	"github.com/kevinkimutai/invoice-management-system/internal/adapters/db"
	handler "github.com/kevinkimutai/invoice-management-system/internal/adapters/handlers"
	"github.com/kevinkimutai/invoice-management-system/internal/adapters/pdf"
	"github.com/kevinkimutai/invoice-management-system/internal/adapters/server"
	application "github.com/kevinkimutai/invoice-management-system/internal/app/core/api"
)

func main() {

	//Get Env variables
	// Init Environment Variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env files")
	}

	//Env Variables
	APP_PORT := os.Getenv("APP_PORT")
	POSTGRES_USERNAME := os.Getenv("POSTGRES_USERNAME")
	POSTGRES_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	DATABASE_PORT := os.Getenv("DB_PORT")
	DATABASE_HOST := os.Getenv("DB_HOST")
	DATABASE_NAME := os.Getenv("DB_NAME")

	DBURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		POSTGRES_USERNAME,
		POSTGRES_PASSWORD,
		DATABASE_HOST,
		DATABASE_PORT,
		DATABASE_NAME)

	//RABBITMQSERVER := os.Getenv("RABBITMQ_SERVER")

	//Dependency injection

	//Database
	//Connect To DB
	dbAdapter := db.NewDB(DBURL)

	//connect to RabbitMQ
	//msgQueue := queue.NewRabbitMQServer(RABBITMQSERVER)

	//PDF Service
	PDFservice := pdf.New()

	//Repositories
	userRepo := application.NewUserRepo(dbAdapter)
	companyRepo := application.NewCompanyRepo(dbAdapter)
	invoiceRepo := application.NewInvoiceRepo(dbAdapter, PDFservice)

	//Services
	//App Services
	userService := handler.NewUserService(userRepo)
	companyService := handler.NewCompanyService(companyRepo)
	invoiceService := handler.NewInvoiceService(invoiceRepo)

	//Auth
	authService, err := auth.New(dbAdapter)
	if err != nil {
		log.Fatal(err)
	}

	//Server
	server := server.New(
		APP_PORT,
		authService,
		userService,
		companyService,
		invoiceService,
	)
	server.Run()

}
