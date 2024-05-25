package main

import (
	ginhandlers "my-learning-box-backend-v2/httpHandlers/gin"
)

func main() {
	dIContainer := BuildDIContainer()
	if err := dIContainer.Invoke(func(g *ginhandlers.GinHttpServer) {
		g.Run()

	}); err != nil {
		panic(err)
	}
}
