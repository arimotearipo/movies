package server

import (
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

	v1 := router.Group("/movies")
	{
		v1.GET("/", handler.GetAllMovies)
		v1.GET("/:id", handler.GetMovieById)
		v1.POST("/", handler.PostMovie)
		v1.PUT("/:id", handler.UpdateMovie)
	}

	v2 := router.Group("/directors")
	{
		v2.GET("/", handler.GetAllDirectors)
		v2.GET("/:id", handler.GetDirectorById)
		v2.POST("/", handler.PostDirector)
		v2.PUT("/:id", handler.UpdateDirector)
	}

	err := router.Run(s.addr)
	if err != nil {
		panic(err)
	}
}
