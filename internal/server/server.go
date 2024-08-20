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

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "false")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

        if c.Request.Method == "OPTIONS" {
            c.Writer.WriteHeader(204)
            return
        }

        c.Next()
    }
}

func NewServer(addr string, store *psqlstorage.Storage) *Server {
	return &Server{addr, store}
}

func (s *Server) Serve() {
	router := gin.Default()

	router.Use(CORSMiddleware())

	handler := handlers.NewHandler(s.store)

	router.GET("/healthcheck", handler.HealthCheck)

	movies := router.Group("/movies")
	{
		movies.GET("/:id", handler.GetMovieById)
		movies.POST("/", handler.PostMovie)
		movies.PUT("/:id", handler.UpdateMovie)
		movies.DELETE("/:id", handler.DeleteMovie)
		movies.GET("/", handler.GetAllMovies)
	}


	directors := router.Group("/directors")
	{
		directors.GET("/:id", handler.GetDirectorById)
		directors.POST("/", handler.PostDirector)
		directors.PUT("/:id", handler.UpdateDirector)
		directors.DELETE("/:id", handler.DeleteDirector)
		directors.GET("/", handler.GetAllDirectors)
	}

	err := router.Run(s.addr)
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
