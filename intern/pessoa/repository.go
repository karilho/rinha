package pessoa

import (
	"context"
	"fmt"
	"github.com/vingarcia/ksql"
	"rinha/intern/domain"
)

type Repository interface {
	InsertPessoa(pessoa domain.Pessoa) error
	GetPessoaByID(id string) (domain.Pessoa, error)
}

type DatabaseRepository struct {
	db ksql.DB
}

func NewDatabaseRepository(db ksql.DB) *DatabaseRepository {
	return &DatabaseRepository{
		db: db,
	}
}

func (r *DatabaseRepository) InsertPessoa(pessoa domain.Pessoa) error {
	query := fmt.Sprintf("INSERT INTO public.pessoas (id, nome, cpfcnpj, nascimento, seguros) VALUES ('%s', '%s', '%s', '%s', '%s')",
		pessoa.ID, pessoa.Name, pessoa.CpfCnpj, pessoa.Nascimento, pessoa.Seguros)

	_, err := r.db.Exec(context.Background(), query)
	if err != nil {
		return err
	}

	return nil
}

func (r *DatabaseRepository) GetPessoaByID(id string) (domain.Pessoa, error) {
	query := fmt.Sprintf("SELECT id, nome, cpfcnpj, nascimento, seguros FROM public.pessoas WHERE id = '%s'", id)

	var pessoaID struct {
		ID string `ksql:"id"`
	}

	err := r.db.QueryOne(context.Background(), &pessoaID, query)
	if err != nil {
		return domain.Pessoa{}, err
	}

	return domain.Pessoa{}, nil
}
