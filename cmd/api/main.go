package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/vingarcia/ksql"
	"github.com/vingarcia/ksql/adapters/kpgx"
	"log"
	"os"
	"rinha/cmd/api/controller"
	"rinha/cmd/api/routes"
	"rinha/intern/pessoa"
)

var DB_URL = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

func main() {
	ctx := context.Background()
	db, err := kpgx.New(ctx, os.Getenv("DATABASE_URL"), ksql.Config{})
	if err != nil {
		log.Fatalf("unable connect to database on url '%s': %s", os.Getenv("DATABASE_URL"), err)
	}
	defer db.Close()

	fmt.Println("Db connected")

	app := fiber.New()
	repo := pessoa.NewDatabaseRepository(db)
	service := pessoa.NewService(repo)
	pessoaController := controller.NewPessoaController(*service)
	routes.InitRoutes(app, pessoaController)
	fmt.Println("Routes initialized")
	if err := app.Listen(":80"); err != nil {
		log.Fatal(err)
	}

}
