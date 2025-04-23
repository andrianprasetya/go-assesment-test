package main

import (
	"flag"
	"fmt"
	"github.com/andrianprasetya/go-assesment-test/database"
	_ "github.com/andrianprasetya/go-assesment-test/database/dialect/postgres"
	_ "github.com/andrianprasetya/go-assesment-test/internal/config"
	"github.com/andrianprasetya/go-assesment-test/internal/repository"
	"github.com/andrianprasetya/go-assesment-test/internal/usecase"
	"github.com/andrianprasetya/go-assesment-test/routes"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
)

func main() {

	// Argument parser for non-sensitive configurations
	host := flag.String("host", os.Getenv("APP_HOST"), "API server host")
	port := flag.String("port", os.Getenv("APP_PORT"), "API server port")
	flag.Parse()

	//initialize fiber
	app := fiber.New()

	//connect database
	db := database.GetConnection()

	//migrate database table
	database.MigrateDatabase(db)

	userRepo := repository.NewUserRepository(db)
	userUC := usecase.NewUserUsecase(userRepo)

	transactionRepo := repository.NewTransactionRepository(db)
	transactionUC := usecase.NewTransactionUsecase(transactionRepo, userRepo)

	routes.SetupRoutes(app, userUC, transactionUC)

	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", *host, *port)))
}
