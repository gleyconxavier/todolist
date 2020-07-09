package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// Activitie user
type Activitie struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

var (
	activities = map[int]*Activitie{}
	seq        = 1
)

// GetActivities from db
func GetActivities(c echo.Context) error {
	return c.JSON(http.StatusOK, activities)
}

// CreateActivitie on db
func CreateActivitie(c echo.Context) error {

	m := new(Activitie)
	if err := c.Bind(m); err != nil {
		return err
	}
	m.ID = seq
	activities[seq] = m
	seq++

	return c.JSON(http.StatusCreated, m.Title)
}

// EditActivitie by id
func EditActivitie(c echo.Context) error {

	id := c.Param("id")
	idPars, err := strconv.Atoi(id)
	if err != nil {
		// handle error
		fmt.Println(err)
	}
	edited := activities[idPars]
	return c.JSON(http.StatusOK, edited)
}

// RemoveActivitie by id
func RemoveActivitie(c echo.Context) error {

	id := c.Param("id")
	idPars, err := strconv.Atoi(id)
	if err != nil {
		// handle error
		fmt.Println(err)
	}
	removed := activities[idPars]
	return c.JSON(http.StatusOK, removed)
}
