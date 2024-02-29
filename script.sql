CREATE TABLE IF NOT EXISTS clientes (
	id int not null,
	saldo int not null,
	limite int not null,
	PRIMARY KEY (id)
	);
CREATE TABLE IF NOT EXISTS transacoes (
	id serial,
	idCliente int not null,
	valor int not null,
	tipo VARCHAR(1) not null,
	descricao VARCHAR(100) not null,
	realizada_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (id, idCliente), 
    FOREIGN KEY (idCliente) REFERENCES clientes(id)
);
INSERT INTO clientes (id, limite, saldo)
  VALUES
    (1,	100000,	0),
    (2,	80000,	0),
    (3,	1000000,	0),
    (4,	10000000,	0),
    (5,	500000,	0);
	