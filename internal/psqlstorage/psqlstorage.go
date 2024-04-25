package psqlstorage

import (
	"errors"
	"log"

	"github.com/arimotearipo/movies/internal/types"
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	db *sqlx.DB
}

type StorageService interface {
	// movies
	GetAllMovies() ([]types.Movie, error)
	GetMovieById(string) ([]types.Movie, error)
	PostMovie(*types.MovieParams) error
	UpdateMovie(string, *types.MovieParams) error

	// directors
	GetAllDirectors() ([]types.Director, error)
	GetDirectorById(string) ([]types.Director, error)
	PostDirector(*types.DirectorParams) error
	UpdateDirector(string, *types.DirectorParams) error
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{db}
}

// movies
func (s *Storage) GetAllMovies() ([]types.Movie, error) {
	queryString := `---sql
	SELECT * FROM movies;
	`

	var result []types.Movie
	err := s.db.Select(&result, queryString)
	if err != nil {
		return []types.Movie{}, err
	}

	return result, nil
}

func (s *Storage) GetMovieById(id string) ([]types.Movie, error) {
	queryString := `---sql
	SELECT * FROM movies WHERE movie_id = $1;
	`

	var result []types.Movie
	err := s.db.Select(&result, queryString, id)
	if err != nil {
		return []types.Movie{}, err
	}

	if len(result) == 0 {
		return []types.Movie{}, errors.New("movie_id not found")
	}

	return result, nil
}

func (s *Storage) PostMovie(m *types.MovieParams) error {
	queryString := `---sql
	INSERT INTO movies (title, director_id, year)
	VALUES ($1, $2, $3);
	`

	_, err := s.db.Exec(queryString, m.Title, m.DirectorID, m.Year)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) UpdateMovie(id string, m *types.MovieParams) error {
	queryString := `---sql
	UPDATE movies
	SET title = $2, director_id = $3, year = $4
	WHERE movie_id = $1;
	`
	rowsAffected, err := s.db.Exec(queryString, id, m.Title, m.DirectorID, m.Year)
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
func (s *Storage) GetAllDirectors() ([]types.Director, error) {
	queryString := `---sql
	SELECT * FROM directors;`

	var result []types.Director
	err := s.db.Select(&result, queryString)
	if err != nil {
		return []types.Director{}, err
	}

	return result, nil
}

func (s *Storage) GetDirectorById(id string) ([]types.Director, error) {
	queryString := `---sql
	SELECT * FROM directors WHERE director_id = $1;`

	var result []types.Director
	err := s.db.Select(&result, queryString, id)
	if err != nil {
		return []types.Director{}, err
	}

	if len(result) == 0 {
		return []types.Director{}, errors.New("director_id not found")
	}

	return result, nil
}

func (s *Storage) PostDirector(d *types.DirectorParams) error {
	queryString := `---sql
	INSERT INTO directors (name, date_of_birth, gender, nationality)
	VALUES ($1, $2, $3, $4);`

	_, err := s.db.Exec(queryString, d.Name, d.DOB, d.Gender, d.Nationality)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) UpdateDirector(id string, d *types.DirectorParams) error {
	queryString := `---sql
	UPDATE directors
	SET name = $2, date_of_birth = $3, gender = $4, nationality = $5
	WHERE director_id = $1;
	`

	rowsAffected, err := s.db.Exec(queryString, id, d.Name, d.DOB, d.Gender, d.Nationality)
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
