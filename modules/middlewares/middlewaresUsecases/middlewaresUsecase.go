package middlewaresUsecases

import "github.com/prachaya-or/eCommerce-api/modules/middlewares/middlewaresRepositories"

type IMiddlewaresUsecase interface {
}

type middlewaresUsecase struct {
	middlewarerepositories middlewaresRepositories.IMiddlewareRepository
}

func MiddlewaresUsecase(middlewarerepositories middlewaresRepositories.IMiddlewareRepository) IMiddlewaresUsecase {
	return &middlewaresUsecase{
		middlewarerepositories: middlewarerepositories,
	}
}
