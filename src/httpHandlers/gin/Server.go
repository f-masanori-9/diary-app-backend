package gin

import (
	"fmt"
	"time"

	"diary-app-backend/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type POSTHandler struct {
	Users *PostUsers
}

type GinHanders struct {
	POST POSTHandler
}

func NewGinHanders(PostUsers *PostUsers) *GinHanders {
	return &GinHanders{
		POST: POSTHandler{
			Users: PostUsers,
		},
	}
}

type GinHttpServer struct {
	config     *config.Config
	ginHanders *GinHanders
	Run        func()
}

func NewGinHTTPServer(ginHanders *GinHanders, config *config.Config) *GinHttpServer {
	port := *config.APIPort

	r := gin.Default()
	r.Use(getCorsSetting())

	authGroup := r.Group("/") // TODO: 認証通す
	authGroup.POST("/users", ginHanders.POST.Users.CreateHandler())

	r.Run(fmt.Sprintf(":%s", port))
	return &GinHttpServer{
		ginHanders: ginHanders,
		Run: func() {
			r.Run(fmt.Sprintf(":%s", port))
		},
		config: config,
	}
}

func getCorsSetting() gin.HandlerFunc {
	return cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			"POST",
			"GET",
			"PUT",
			"DELETE",
			"OPTIONS",
			"PATCH",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Authorization",
			"Referer",
		},
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	})
}
