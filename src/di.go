package main

import (
	"diary-app-backend/application/domain/repository"
	"diary-app-backend/application/infrastructure/postgres"
	"diary-app-backend/application/usecase"
	"diary-app-backend/config"
	"diary-app-backend/httpHandlers/gin"
	ginhandlers "diary-app-backend/httpHandlers/gin"

	"go.uber.org/dig"
)

func BuildDIContainer() *dig.Container {
	container := dig.New()

	container.Provide(config.NewConfig)
	container.Provide(postgres.ConnectDB)
	container.Provide(postgres.CreateGormInstance)
	container.Provide(postgres.NewUserRepository, []dig.ProvideOption{dig.As(new(repository.IUserRepository))}...)
	container.Provide(usecase.NewCreateUserUseCase)
	container.Provide(ginhandlers.NewPostUsers)
	container.Provide(ginhandlers.NewGinHanders)
	container.Provide(ginhandlers.NewGinHTTPServer)
	container.Provide(gin.NewGinHanders)

	return container
}
