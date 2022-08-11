package echo

import "github.com/labstack/echo/v4"

type Context struct {
	echo.Context
	UserID string
	Roles  []string
}

// TODO deletar se não for utilizar
