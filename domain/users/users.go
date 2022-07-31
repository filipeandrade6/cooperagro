package users

import (
	"github.com/filipeandrade6/cooperagro-go/adapters/log"
	"github.com/filipeandrade6/cooperagro-go/adapters/repo"
)

// Usually its here where the business logic complexity builds up,
// but since this is just an example both these functions are actually
// very simple, but in real world scenarios you would want to make
// these contain all your business logic.

type Service struct {
	logger    log.Provider
	usersRepo repo.Users
}

func NewService(
	logger log.Provider,
	usersRepo repo.Users,
) Service {
	return Service{
		logger:    logger,
		usersRepo: usersRepo,
	}
}
