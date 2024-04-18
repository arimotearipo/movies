package service

import (
	"github.com/arimotearipo/movies/internal/database"
	"github.com/arimotearipo/movies/internal/types"
)

// movies
func GetAllMovies() ([]types.Movie, error) {
	queryString := `---sql
	SELECT * FROM movies;
	`

	var result []types.Movie
	err := database.DB.Select(&result, queryString)
	if err != nil {
		return []types.Movie{}, err
	}

	return result, nil
}

func GetMovieById(id string) ([]types.Movie, error) {
	queryString := `---sql
	SELECT * FROM movies WHERE movie_id = $1;
	`

	var result []types.Movie
	err := database.DB.Select(&result, queryString, id)
	if err != nil {
		return []types.Movie{}, err
	}

	return result, nil
}

func PostMovie(m *types.MovieParams) error {
	queryString := `---sql
	INSERT INTO movies (title, director_id, year)
	VALUES ($1, $2, $3);
	`

	_, err := database.DB.Exec(queryString, m.Title, m.DirectorID, m.Year)
	if err != nil {
		return err
	}

	return nil
}

// directors
func GetAllDirectors() ([]types.Director, error) {
	queryString := `---sql
	SELECT * FROM directors;`

	var result []types.Director
	err := database.DB.Select(&result, queryString)
	if err != nil {
		return []types.Director{}, err
	}

	return result, nil
}

func GetDirectorById(id string) ([]types.Director, error) {
	queryString := `---sql
	SELECT * FROM directors WHERE director_id = $1;`

	var result []types.Director
	err := database.DB.Select(&result, queryString, id)
	if err != nil {
		return []types.Director{}, err
	}

	return result, nil
}

func PostDirector(d *types.DirectorParams) error {
	queryString := `---sql
	INSERT INTO directors (name, date_of_birth, gender, nationality)
	VALUES ($1, $2, $3, $4);`

	_, err := database.DB.Exec(queryString, d.Name, d.DOB, d.Gender, d.Nationality)
	if err != nil {
		return err
	}

	return nil
}
