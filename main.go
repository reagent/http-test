package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type App struct {
	Router http.Handler
}

type Router struct {
	*gin.Engine

	App *App
}

func newRouter(app *App) *Router {
	router := &Router{App: app}

	gin.SetMode(gin.ReleaseMode)

	router.Engine = gin.New()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hi2u\n")
	})

	return router
}

func newApp() *App {
	app := &App{}
	app.Router = newRouter(app)

	return app
}

func main() {
	app := newApp()

	server := &http.Server{
		Addr:    ":8080",
		Handler: app.Router,
	}

	server.ListenAndServe()
}
