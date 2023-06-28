package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func main() {
	configureApi()
}

func configureApi() {
	e := echo.New()

	e.GET("/people", func(c echo.Context) error {
		return c.JSON(http.StatusOK, findAll())
	})

	e.GET("/people/:name", func(c echo.Context) error {
		name := c.Param("name")
		person := findByName(name)
		if person != nil {
			return c.JSON(http.StatusOK, findByName(name))
		}
		return c.NoContent(http.StatusNotFound)
	})

	e.POST("/people", func(c echo.Context) error {
		person := new(Person)
		c.Bind(person)
		save(*person)
		return c.NoContent(http.StatusAccepted)
	})

	e.PUT("/people/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		person := new(Person)
		c.Bind(person)
		update(id, *person)
		return c.NoContent(http.StatusAccepted)
	})

	e.DELETE("/people/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		delete(id)
		return c.NoContent(http.StatusAccepted)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
