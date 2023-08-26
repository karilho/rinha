package routes

import (
	"github.com/gofiber/fiber/v2"
	"rinha/cmd/api/controller"
)

func InitRoutes(app *fiber.App, pessoaController *controller.PessoaController) {

	app.Post("/pessoa", pessoaController.Create)
	app.Get("/pessoa/:pessoaId", pessoaController.Get)
	app.Get("/pessoa", pessoaController.GetByTerm)

}
