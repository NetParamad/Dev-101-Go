package config

import (
	_ "github.com/go-sql-driver/mysql"

)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

func ConnectDB() (*Config, error) {
	return &Config{
		Host:     "localhost",
		Port:     "3306",
		User:     "root",
		Password: "",
		Dbname:   "test",
	}, nil
}