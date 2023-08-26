package domain

type Pessoa struct {
	ID         string `json:"id"`
	Name       string `json:"nome"`
	CpfCnpj    string `json:"cpfcnpj"`
	Nascimento string `json:"nascimento"`
	Seguros    string `json:"seguros"`
}
