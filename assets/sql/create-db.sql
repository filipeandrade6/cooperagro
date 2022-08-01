INSERT INTO funcoes_usuarios (funcao) VALUES ('administrador');

INSERT INTO usuarios (funcao_usuario_id, primeiro_nome, ultimo_nome, email, telefone, endereco, latitude, longitude, criado_em, atualizado_em)
VALUES (1, 'filipe', 'andrade', 'filipe@email.com', '5511555554444', 'rua tal conjunto tal', -12.123123, -13.234234, now(), now());

-- ----------------------------------------------------------

CREATE TABLE unidades_de_medida (
	id SERIAL PRIMARY KEY,
	tipo VARCHAR(20) NOT NULL
);

CREATE TABLE produtos_base (
	id SERIAL PRIMARY KEY,
	nome VARCHAR(30) NOT NULL,
	criado_em TIMESTAMP NOT NULL,
	atualizado_em TIMESTAMP NOT NULL
);

CREATE TABLE produtos (
	id SERIAL PRIMARY KEY,
	nome VARCHAR(30) NOT NULL,
	criado_em TIMESTAMP NOT NULL,
	atualizado_em TIMESTAMP NOT NULL,
	produto_base_id INTEGER NOT NULL,
	unidade_de_medida_id INTEGER NOT NULL,
	CONSTRAINT produto_base_id FOREIGN KEY (produto_base_id) REFERENCES produtos_base(id),
	CONSTRAINT unidade_de_medida_id FOREIGN KEY (unidade_de_medida_id) REFERENCES unidades_de_medida(id)
);

CREATE TABLE funcoes_usuarios (
	id SERIAL PRIMARY KEY,
	funcao VARCHAR(15)
);

CREATE TABLE usuarios (
	id SERIAL PRIMARY KEY,
	primeiro_nome VARCHAR(30) NOT NULL,
	ultimo_nome VARCHAR(30) NOT NULL,
	email VARCHAR(50) NOT NULL,
	telefone VARCHAR(14) NOT NULL,
	endereco VARCHAR(100),
	latitude DECIMAL(8,6),
	longitude DECIMAL(9,6),
	criado_em TIMESTAMP NOT NULL,
	atualizado_em TIMESTAMP NOT NULL,
	funcao_usuario_id INT NOT NULL,
	CONSTRAINT funcao_usuario_id FOREIGN KEY (funcao_usuario_id) REFERENCES funcoes_usuarios(id)
);

CREATE TABLE inventario (
	id SERIAL PRIMARY KEY,
	quantidade INTEGER NOT NULL,
	usuario_id INTEGER NOT NULL,
	CONSTRAINT usuario_id FOREIGN KEY (usuario_id) REFERENCES usuarios(id)
);