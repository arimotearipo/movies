package handlers

import (
	"fmt"
	"net/http"

	"github.com/arimotearipo/movies/internal/service"
	"github.com/arimotearipo/movies/internal/types"
	"github.com/gin-gonic/gin"
)

// movies
func GetAllMovies(c *gin.Context) {
	result, err := service.GetAllMovies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}

	c.JSON(http.StatusOK, gin.H{
		"movies": result,
	})
}

func GetMovieById(c *gin.Context) {
	id := c.Params[0].Value

	result, err := service.GetMovieById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}

	c.JSON(200, gin.H{
		"movies": result,
	})
}

func PostMovie(c *gin.Context) {
	var m types.MovieParams

	err := c.ShouldBind(&m)
	if err != nil {
		panic(err)
	}

	err = service.PostMovie(&m)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data updated",
	})
}

// directors
func GetAllDirectors(c *gin.Context) {
	result, err := service.GetAllDirectors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}

	c.JSON(200, gin.H{
		"data": result,
	})
}

func GetDirectorById(c *gin.Context) {
	id := c.Params[0].Value

	result, err := service.GetDirectorById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}

	c.JSON(200, gin.H{
		"data": result,
	})
}

func PostDirector(c *gin.Context) {
	var d types.DirectorParams

	err := c.ShouldBind(&d)
	if err != nil {
		panic(err)
	}

	fmt.Println(d)
	err = service.PostDirector(&d)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data updated",
	})
}
