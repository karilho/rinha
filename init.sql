CREATE TABLE IF NOT EXISTS pessoas(
                                     id INT PRIMARY KEY NOT NULL,
                                     nome varchar(32) NOT NULL,
                                     cpfcnpj varchar(11) NOT NULL,
                                     nascimento varchar(9) NOT NULL,
                                     seguros TEXT NULL
    );
