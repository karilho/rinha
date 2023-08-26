package controller

import (
	"github.com/gofiber/fiber/v2"
	"rinha/intern/domain"
	"rinha/intern/pessoa"
)

type PessoaController struct {
	service pessoa.Service
}

func NewPessoaController(service pessoa.Service) *PessoaController {
	return &PessoaController{
		service: service,
	}
}

func (c *PessoaController) Create(ctx *fiber.Ctx) error {
	var pessoa *domain.Pessoa

	if err := ctx.BodyParser(&pessoa); err != nil {
		return err
	}

	err := c.service.CreatePessoa(pessoa)
	if err != nil {
		return err
	}

	return ctx.JSON(&pessoa)
}

func (c *PessoaController) Get(ctx *fiber.Ctx) error {
	pessoaID := ctx.Params("pessoaId")

	pessoa, err := c.service.GetPessoaByID(pessoaID)
	if err != nil {
		return err
	}

	return ctx.JSON(pessoa)
}

func (c *PessoaController) GetByTerm(ctx *fiber.Ctx) error {
	pessoaTerm := ctx.Query("pessoaTerm")

	pessoas, err := c.service.GetPessoaByTerm(pessoaTerm)
	if err != nil {
		return err
	}

	return ctx.JSON(pessoas)
}
