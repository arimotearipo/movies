package service

import (
	"errors"
	"log"

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

func UpdateMovie(id string, m *types.MovieParams) error {
	queryString := `---sql
	UPDATE movies
	SET title = $2, director_id = $3, year = $4
	WHERE movie_id = $1;
	`
	rowsAffected, err := database.DB.Exec(queryString, id, m.Title, m.DirectorID, m.Year)
	if err != nil {
		return err
	}

	n, err := rowsAffected.RowsAffected()
	if err != nil || n == 0 {
		return errors.New("no rows affected")
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

func UpdateDirector(id string, d *types.DirectorParams) error {
	queryString := `---sql
	UPDATE directors
	SET name = $2, date_of_birth = $3, gender = $4, nationality = $5
	WHERE director_id = $1;
	`

	rowsAffected, err := database.DB.Exec(queryString, id, d.Name, d.DOB, d.Gender, d.Nationality)
	if err != nil {
		log.Print(err)
		return err
	}

	n, err := rowsAffected.RowsAffected()
	if err != nil || n == 0 {
		return errors.New("no rows affected")
	}

	return nil
}
