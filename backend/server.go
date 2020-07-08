package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// rotas
	e.GET("/listar-atividades", func(c echo.Context) error {
		return c.String(http.StatusOK, "Recebi um GET request.")
	})

	e.POST("/criar-atividade", func(c echo.Context) error {
		return c.String(http.StatusOK, "Recebi um POST request.")
	})

	e.PUT("/atualizar-atividade", func(c echo.Context) error {
		return c.String(http.StatusOK, "Recebi um PUT request.")
	})

	e.DELETE("/remover-atividade", func(c echo.Context) error {
		return c.String(http.StatusOK, "Recebi um Delete request.")
	})

	// log
	e.Logger.Fatal(e.Start(":1323"))
}
