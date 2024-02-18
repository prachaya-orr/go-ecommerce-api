package middlewaresRepositories

import "github.com/jmoiron/sqlx"

type IMiddlewareRepository interface {
}

type middlewaresRepository struct {
	db *sqlx.DB
}

func MiddlewaresRepository(db *sqlx.DB) IMiddlewareRepository {
	return &middlewaresRepository{
		db: db,
	}
}
