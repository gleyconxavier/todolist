package main

import (
	"net/http"

	"github.com/gleyconxavier/todolist/backend/router"
	"github.com/labstack/echo"
)

func main() {
	r := router.New()

	// rotas
	r.GET("/listar-atividades", func(c echo.Context) error {
		return c.String(http.StatusOK, "Recebi um GET request.")
	})

	r.POST("/criar-atividade", func(c echo.Context) error {
		return c.String(http.StatusOK, "Recebi um POST request.")
	})

	r.PUT("/atualizar-atividade", func(c echo.Context) error {
		return c.String(http.StatusOK, "Recebi um PUT request.")
	})

	r.DELETE("/remover-atividade", func(c echo.Context) error {
		return c.String(http.StatusOK, "Recebi um Delete request.")
	})

	// log
	r.Logger.Fatal(r.Start(":1323"))
}
