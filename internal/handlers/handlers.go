package handlers

import (
	"net/http"

	"github.com/arimotearipo/movies/internal/psqlstorage"
	"github.com/arimotearipo/movies/internal/types"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	storage *psqlstorage.Storage
}

type HandlerFuncs interface {
	GetAllMovies(*gin.Context)
	GetMovieById(*gin.Context)
	PostMovie(*gin.Context)
	UpdateMovie(*gin.Context)
	DeleteMovie(*gin.Context)

	GetAllDirectors(*gin.Context)
	GetDirectorById(*gin.Context)
	PostDirector(*gin.Context)
	UpdateDirector(*gin.Context)
	DeleteDirector(*gin.Context)

	HealthCheck(*gin.Context)
}

func NewHandler(s *psqlstorage.Storage) *Handler {
	return &Handler{s}
}

// movies
func (h *Handler) GetAllMovies(c *gin.Context) {
	result, err := h.storage.GetAllMovies()
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

func (h *Handler) GetMovieById(c *gin.Context) {
	id := c.Params[0].Value

	result, err := h.storage.GetMovieById(id)
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

func (h *Handler) PostMovie(c *gin.Context) {
	var m types.MovieParams

	err := c.ShouldBind(&m)
	if err != nil {
		panic(err)
	}

	// TODO: verify if director_id exists
	_, err = h.storage.GetDirectorById(m.DirectorID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = h.storage.PostMovie(&m)
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

func (h *Handler) UpdateMovie(c *gin.Context) {
	movieId := c.Params[0].Value

	var m types.MovieParams
	err := c.ShouldBind(&m)
	if err != nil {
		panic(err)
	}

	// TODO: verify if movie_id exists

	_, err = h.storage.GetDirectorById(m.DirectorID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = h.storage.UpdateMovie(movieId, &m)
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

func (h *Handler) DeleteMovie(c *gin.Context) {
	movieId := c.Params[0].Value

	err := h.storage.DeleteMovie(movieId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Successfully deleted",
		})
	}

}

// directors
func (h *Handler) GetAllDirectors(c *gin.Context) {
	result, err := h.storage.GetAllDirectors()
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

func (h *Handler) GetDirectorById(c *gin.Context) {
	id := c.Params[0].Value

	result, err := h.storage.GetDirectorById(id)
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

func (h *Handler) PostDirector(c *gin.Context) {
	var d types.DirectorParams

	err := c.ShouldBind(&d)
	if err != nil {
		panic(err)
	}

	err = h.storage.PostDirector(&d)
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

func (h *Handler) UpdateDirector(c *gin.Context) {
	id := c.Params[0].Value
	var d types.DirectorParams

	err := c.ShouldBind(&d)
	if err != nil {
		panic(err)
	}

	err = h.storage.UpdateDirector(id, &d)
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

func (h *Handler) DeleteDirector(c *gin.Context) {
	directorId := c.Params[0].Value

	err := h.storage.DeleteDirector(directorId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Sucessfully deleted",
		})
	}
}

// healthcheck
func (h *Handler) HealthCheck(c *gin.Context) {
	err := h.storage.HealthCheck()

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "connection healthy",
		})
	}
}
