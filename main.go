package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func App() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	app := gin.New()

	app.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hi2u\n")
	})

	return app
}

func main() {
	app := App()

	server := &http.Server{
		Addr:    ":8080",
		Handler: app,
	}

	server.ListenAndServe()
}
