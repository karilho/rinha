CREATE TABLE IF NOT EXISTS pessoas(
                                     id TEXT PRIMARY KEY not null,
                                     nome varchar(32) NOT NULL,
                                     cpfcnpj varchar(11) NOT NULL,
                                     nascimento char(10) NOT NULL,
                                     seguros TEXT NULL
    );
