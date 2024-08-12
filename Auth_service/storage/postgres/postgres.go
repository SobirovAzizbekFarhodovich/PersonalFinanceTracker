package postgres

import (
	"database/sql"
	"fmt"

	"auth/config"
	u "auth/storage"

	_ "github.com/lib/pq"
)

type Storage struct {
	Db    *sql.DB
	Users u.UserI
}

func NewPostgresStorage() (u.StorageI, error) {
	config := config.Load()
	con := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.PostgresUser, config.PostgresPassword,
		config.PostgresHost, config.PostgresPort,
		config.PostgresDatabase)
	db, err := sql.Open("postgres", con)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Storage{Db: db}, err
}

func (s *Storage) User() u.UserI {
	if s.Users == nil {
		s.Users = &UserStorage{s.Db}
	}

	return s.Users
}
