package handlers

import (
	"net/http"

	"github.com/gleyconxavier/todolist/backend/database"
	"github.com/labstack/echo"
)

// Activitie user
type Activitie struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// GetActivities from db
func GetActivities(c echo.Context) error {
	db := database.ConnectDb()
	activities := []Activitie{}
	db.Select(&activities, "SELECT * FROM activities")

	return c.JSON(http.StatusOK, activities)
}

// CreateActivitie on db
func CreateActivitie(c echo.Context) error {

	m := new(Activitie)
	if err := c.Bind(m); err != nil {
		return err
	}

	db := database.ConnectDb()
	query := `INSERT INTO activities (title) VALUES ($1)`
	db.MustExec(query, m.Title)

	return c.JSON(http.StatusCreated, "Activitie created.")
}

// EditActivitie by id
func EditActivitie(c echo.Context) error {

	id := c.Param("id")

	m := new(Activitie)
	if err := c.Bind(m); err != nil {
		return err
	}

	db := database.ConnectDb()
	query := `UPDATE activities SET title = $1 WHERE id = $2;`
	db.MustExec(query, m.Title, id)

	return c.JSON(http.StatusOK, "Activitie updated.")
}

// RemoveActivitie by id
func RemoveActivitie(c echo.Context) error {

	id := c.Param("id")

	db := database.ConnectDb()
	query := `DELETE FROM activities WHERE id = $1;`
	db.MustExec(query, id)

	return c.JSON(http.StatusOK, "Activitie removed.")
}
