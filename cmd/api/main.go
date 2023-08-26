package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/vingarcia/ksql"
	"github.com/vingarcia/ksql/adapters/kpgx"
	"log"
	"rinha/cmd/api/controller"
	"rinha/cmd/api/routes"
	"rinha/intern/pessoa"
)

var DB_URL = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

func main() {
	ctx := context.Background()
	db, err := kpgx.New(ctx, DB_URL, ksql.Config{})
	if err != nil {
		log.Fatalf("unable connect to database: %s", err)
	}
	defer db.Close()

	fmt.Println("Db connected")

	app := fiber.New()
	repo := pessoa.NewDatabaseRepository(db)
	service := pessoa.NewService(repo)
	pessoaController := controller.NewPessoaController(*service)

	routes.InitRoutes(app, pessoaController)
	fmt.Println("Routes initialized")
	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}

}

/*

	_, err = db.Exec(ctx,
		"INSERT INTO public.pessoas (id, nome, cpfcnpj, nascimento, seguros) VALUES (4, 'lucas', 'eokwoekwo', '293823', '[vida,rural]')")
	if err != nil {
		log.Fatalf("unable to insert pessoa: %s", err)
	}

	xx, err := db.Exec(ctx, "SELECT * FROM public.pessoas")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(xx)

*/
