package types

type MovieParams struct {
	Title      string `json:"title"`
	DirectorID string `json:"director_id"`
	Year       string `json:"year"`
}

type Movie struct {
	MovieID    string `db:"movie_id"`
	Title      string `db:"title"`
	DirectorID string `db:"director_id"`
	Year       string `db:"year"`
}

type DirectorParams struct {
	Name        string `json:"name"`
	DOB         string `json:"date_of_birth"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
}

type Director struct {
	DirectorID  string `db:"director_id"`
	Name        string `db:"name"`
	DOB         string `db:"date_of_birth"`
	Gender      string `db:"gender"`
	Nationality string `db:"nationality"`
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type SchemaExists struct {
	T1 bool `db:"table1_exists"`
	T2 bool `db:"table2_exists"`
}
