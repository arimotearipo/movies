package handlers

import (
	"net/http"

	"github.com/arimotearipo/movies/internal/service"
	"github.com/arimotearipo/movies/internal/types"
	"github.com/gin-gonic/gin"
)

// movies
func GetAllMovies(c *gin.Context) {
	result, err := service.GetAllMovies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"movies": result,
		})
	}
}

func GetMovieById(c *gin.Context) {
	id := c.Params[0].Value

	result, err := service.GetMovieById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"movies": result,
		})
	}
}

func PostMovie(c *gin.Context) {
	var m types.MovieParams

	err := c.ShouldBind(&m)
	if err != nil {
		panic(err)
	}

	// TODO: verify if director_id exists
	_, err = service.GetDirectorById(m.DirectorID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = service.PostMovie(&m)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Data updated",
		})
	}
}

func UpdateMovie(c *gin.Context) {
	id := c.Params[0].Value

	var m types.MovieParams
	err := c.ShouldBind(&m)
	if err != nil {
		panic(err)
	}

	// TODO: verify if director_id exists
	_, err = service.GetDirectorById(m.DirectorID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = service.UpdateMovie(id, &m)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Successfully updated",
		})
	}

}

// directors
func GetAllDirectors(c *gin.Context) {
	result, err := service.GetAllDirectors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"data": result,
		})
	}
}

func GetDirectorById(c *gin.Context) {
	id := c.Params[0].Value

	result, err := service.GetDirectorById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"data": result,
		})
	}
}

func PostDirector(c *gin.Context) {
	var d types.DirectorParams

	err := c.ShouldBind(&d)
	if err != nil {
		panic(err)
	}

	err = service.PostDirector(&d)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Data updated",
		})
	}
}

func UpdateDirector(c *gin.Context) {
	id := c.Params[0].Value
	var d types.DirectorParams

	err := c.ShouldBind(&d)
	if err != nil {
		panic(err)
	}

	err = service.UpdateDirector(id, &d)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Successfully updated",
		})
	}

}
