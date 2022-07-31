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

	FOREIGN KEY (produto_base_id) REFERENCES produtos_base(id),
	FOREIGN KEY (unidade_de_medida_id) REFERENCES unidades_de_medida(id)
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

	FOREIGN KEY (funcao_usuario_id) REFERENCES funcoes_usuarios(id)
);

CREATE TABLE inventario (
	id SERIAL PRIMARY KEY,
	quantidade INTEGER NOT NULL,

	FOREIGN KEY (usuario_id) REFERENCES usuarios(id)
);