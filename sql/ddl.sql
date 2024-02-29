CREATE TABLE IF NOT EXISTS transacoes (
	id serial,
	idCliente int not null,
	valor int not null,
	tipo VARCHAR(1) not null,
	descricao VARCHAR(100) not null,
	realizada_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (id, idCliente), 
    FOREIGN KEY (id) REFERENCES clinte(id)
);
CREATE TABLE IF NOT EXISTS clientes (
	id int not null,
	saldo int not null,
	limite int not null,
	PRIMARY KEY (id)
	);
	