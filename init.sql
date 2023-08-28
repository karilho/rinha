CREATE TABLE IF NOT EXISTS pessoas(
                                     id TEXT PRIMARY KEY NOT NULL,
                                     nome varchar(32) NOT NULL,
                                     cpfcnpj varchar(11) UNIQUE NOT NULL,
                                     nascimento char(10) NOT NULL,
                                     seguros TEXT NULL
    );
