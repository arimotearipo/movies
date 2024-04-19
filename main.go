package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/arimotearipo/movies/internal/database"
	"github.com/arimotearipo/movies/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	myEnv, err := godotenv.Read()
	if err != nil {
		log.Fatal(err)
	}

	database.InitDatabase(myEnv)
	defer database.CloseDB()

	router := gin.Default()

	v1 := router.Group("/movies")
	{
		v1.GET("/", handlers.GetAllMovies)
		v1.GET("/:id", handlers.GetMovieById)
		v1.POST("/", handlers.PostMovie)
		v1.PUT("/:id", handlers.UpdateMovie)
	}

	v2 := router.Group("/directors")
	{
		v2.GET("/", handlers.GetAllDirectors)
		v2.GET("/:id", handlers.GetDirectorById)
		v2.POST("/", handlers.PostDirector)
		v2.PUT("/:id", handlers.UpdateDirector)
	}

	srv := &http.Server{
		Addr:    "localhost:" + myEnv["PORT"],
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)
	<-quitChan

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}
