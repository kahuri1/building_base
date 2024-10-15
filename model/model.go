package model

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLmode  string
}

type Building struct {
	ID            int    `json:"id" db:"id"`
	Name          string `json:"name" db:"name"`
	City          string `json:"city" db:"city"`
	YearCompleted string `json:"year_completed" db:"year_completed"`
	Floors        string `json:"floors" db:"floors"`
}
