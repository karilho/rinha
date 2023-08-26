package domain

type Pessoa struct {
	ID         string   `ksql:"id" json:"id"`
	Name       string   `ksql:"nome" json:"nome"`
	CpfCnpj    string   `ksql:"cpfcnpj" json:"cpfcnpj"`
	Nascimento string   `ksql:"nascimento" json:"nascimento"`
	Seguros    []string `ksql:"seguros,json" json:"seguros"`
}

// Modifier do ksql, qndo colocamos ,json -> ele vai converter e fazer o procedimento do marshal automaticamente.
