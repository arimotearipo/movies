package server

import (
	"log"

	"github.com/arimotearipo/movies/internal/handlers"
	"github.com/arimotearipo/movies/internal/psqlstorage"
	"github.com/gin-gonic/gin"
)

type Server struct {
	addr  string
	store *psqlstorage.Storage
}

func NewServer(addr string, store *psqlstorage.Storage) *Server {
	return &Server{addr, store}
}

func (s *Server) Serve() {
	router := gin.Default()

	handler := handlers.NewHandler(s.store)

	router.GET("/healthcheck", handler.HealthCheck)

	movies := router.Group("/movies")
	{
		movies.GET("/", handler.GetAllMovies)
		movies.GET("/:id", handler.GetMovieById)
		movies.POST("/", handler.PostMovie)
		movies.PUT("/:id", handler.UpdateMovie)
		movies.DELETE("/:id", handler.DeleteMovie)
	}

	directors := router.Group("/directors")
	{
		directors.GET("/", handler.GetAllDirectors)
		directors.GET("/:id", handler.GetDirectorById)
		directors.POST("/", handler.PostDirector)
		directors.PUT("/:id", handler.UpdateDirector)
		directors.DELETE("/:id", handler.DeleteDirector)
	}

	err := router.Run(s.addr)
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
