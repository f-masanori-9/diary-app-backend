package main

import (
	"my-learning-box-backend-v2/application/domain/repository"
	"my-learning-box-backend-v2/application/infrastructure/postgres"
	"my-learning-box-backend-v2/application/usecase"
	"my-learning-box-backend-v2/config"
	"my-learning-box-backend-v2/httpHandlers/gin"
	ginhandlers "my-learning-box-backend-v2/httpHandlers/gin"

	"go.uber.org/dig"
)

func BuildDIContainer() *dig.Container {
	container := dig.New()

	container.Provide(config.NewConfig)
	container.Provide(postgres.ConnectDB)
	container.Provide(postgres.CreateGormInstanceByConnection)
	container.Provide(postgres.NewUserRepository, []dig.ProvideOption{dig.As(new(repository.IUserRepository))}...)
	container.Provide(usecase.NewCreateUserUseCase)
	container.Provide(ginhandlers.NewPostUsers)
	container.Provide(ginhandlers.NewGinHanders)
	container.Provide(ginhandlers.NewGinHTTPServer)
	container.Provide(gin.NewGinHanders)

	return container
}
