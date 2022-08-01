package domain

import "time"

type UnidadeDeMedida struct {
	ID   int
	Tipo string
}

type ProdutoBase struct {
	ID           int
	Nome         string
	CriadoEm     *time.Time
	AtualizadoEm *time.Time
}

type Produto struct {
	ID                int
	Nome              string
	ProdutoBaseID     int
	UnidadeDeMedidaID int
	CriadoEm          *time.Time
	AtualizadoEm      *time.Time
}

type FuncaoUsuario struct {
	ID     int
	Funcao string
}

type Usuario struct {
	ID              int
	PrimeiroNome    string
	UltimoNome      string
	Email           string
	Telefone        string
	Endereco        string
	Latitude        float32
	Longitude       float32
	FuncaoUsuarioID int
	CriadoEm        *time.Time
	AtualizadoEm    *time.Time
}
