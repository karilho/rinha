package controller

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"rinha/intern/domain"
	"rinha/intern/pessoa"
)

var (
	validate       = validator.New()
	ErrInvalidJson = errors.New("invalid json")
	ErrNotFound    = errors.New("pessoa not found")
	InvalidDtoErr  = errors.New("invalid request")
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

	var pessoareq struct {
		Name       string   `json:"nome" validate:"required,max=32"`
		CpfCnpj    string   `json:"cpfcnpj" validate:"required,max=14"`
		Nascimento string   `json:"nascimento" validate:"required,datetime=2006-01-02"`
		Seguros    []string `json:"seguros" validate:"dive,max=32"`
	}

	if err := ctx.BodyParser(&pessoareq); err != nil {
		return fiber.NewError(http.StatusBadRequest, fmt.Sprintf("invalid json: %s", err.Error()))
	}

	if err := validate.Struct(&pessoareq); err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	PessoaSucess, err := c.service.CreatePessoa(&domain.Pessoa{
		Name:       pessoareq.Name,
		CpfCnpj:    pessoareq.CpfCnpj,
		Nascimento: pessoareq.Nascimento,
		Seguros:    pessoareq.Seguros,
	})
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(map[string]any{
		"id":         PessoaSucess.ID,
		"nome":       PessoaSucess.Name,
		"cpfcnpj":    PessoaSucess.CpfCnpj,
		"nascimento": PessoaSucess.Nascimento,
		"seguros":    PessoaSucess.Seguros,
	})
}

func (c *PessoaController) Get(ctx *fiber.Ctx) error {
	pessoaID := ctx.Params("pessoaId")

	pessoa, err := c.service.GetPessoaByID(pessoaID)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(pessoa)
}

func (c *PessoaController) GetByTerm(ctx *fiber.Ctx) error {
	pessoaTerm := ctx.Query("t")

	if pessoaTerm == "" {
		return fiber.NewError(http.StatusBadRequest, "invalid query param")
	}

	pessoas, err := c.service.GetPessoaByTerm(pessoaTerm)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, ErrInvalidJson.Error())
	}

	if pessoas == nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(ErrNotFound.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(pessoas)
}
