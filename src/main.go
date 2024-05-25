package main

import (
	ginhandlers "diary-app-backend/httpHandlers/gin"
)

func main() {
	dIContainer := BuildDIContainer()
	if err := dIContainer.Invoke(func(g *ginhandlers.GinHttpServer) {
		g.Run()

	}); err != nil {
		panic(err)
	}
}
