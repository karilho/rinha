package pessoa

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/vingarcia/ksql"
	"rinha/intern/domain"
)

// Como segundo parametro eu passo o nome da coluna que será usada como id
var pessoasTable = ksql.NewTable("pessoas", "id")

type Repository interface {
	InsertPessoa(pessoa *domain.Pessoa) error
	GetPessoaByID(id string) (domain.Pessoa, error)
	GetPessoaByTerm(term string) ([]domain.Pessoa, error)
}

type DatabaseRepository struct {
	db ksql.DB
}

func NewDatabaseRepository(db ksql.DB) *DatabaseRepository {
	return &DatabaseRepository{
		db: db,
	}
}

func (r *DatabaseRepository) InsertPessoa(pessoa *domain.Pessoa) error {
	id := uuid.New().String()
	fmt.Println(id)
	pessoa.ID = id
	fmt.Println(pessoa)
	err := r.db.Insert(context.Background(), pessoasTable, pessoa)

	if err != nil {
		return fmt.Errorf("unable to insert pessoa: %w", err)
	}

	return nil
}

func (r *DatabaseRepository) GetPessoaByID(id string) (p domain.Pessoa, _ error) {
	query := "FROM pessoas WHERE id = $1"

	err := r.db.QueryOne(context.Background(), &p, query, id)
	if err != nil {
		return p, fmt.Errorf("unable to query pessoa: %w", err)
	}

	return p, nil
}

func (r *DatabaseRepository) GetPessoaByTerm(term string) (p []domain.Pessoa, _ error) {
	query := "FROM pessoas WHERE nome LIKE $1 OR seguros LIKE $1"

	err := r.db.Query(context.Background(), &p, query, "%"+term+"%")
	if err != nil {
		return p, fmt.Errorf("unable to query pessoa: %w", err)
	}

	return p, nil
}
