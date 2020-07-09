package handlers

import (
	"github.com/gleyconxavier/todolist/backend/router"
)

// HandleRoutes instance
func HandleRoutes() {
	r := router.New()

	// rotas
	r.GET("/listar-atividades", GetActivities)

	r.POST("/criar-atividade", CreateActivitie)

	r.PUT("/atualizar-atividade/:id", EditActivitie)

	r.DELETE("/remover-atividade/:id", RemoveActivitie)

	// log
	r.Logger.Fatal(r.Start(":1323"))
}
