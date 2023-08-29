package pessoa

import "rinha/intern/domain"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreatePessoa(pessoa *domain.Pessoa) (domain.Pessoa, error) {
	return s.repo.InsertPessoa(pessoa)
}

func (s *Service) GetPessoaByID(id string) (domain.Pessoa, error) {
	return s.repo.GetPessoaByID(id)
}

func (s *Service) GetPessoaByTerm(term string) ([]domain.Pessoa, error) {
	return s.repo.GetPessoaByTerm(term)
}
