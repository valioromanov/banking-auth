package domain

import (
	"banking-auth/errs"

	"github.com/jmoiron/sqlx"
)

type AuthRepository interface {
	FindBy(username string, password string) (*Login, *errs.AppError)
}

type AuthRepositoryDb struct {
	client *sqlx.DB
}

func NewAuthRepository(client *sqlx.DB) AuthRepositoryDb {
	return AuthRepositoryDb{client}
}
