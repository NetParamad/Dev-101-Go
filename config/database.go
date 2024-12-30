package config

import (

)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func DBConfig() *Config {
	return &Config{
		Host:     "localhost",
		Port:     "3306",
		User:     "myuser",
		Password: "mypassword",
		DBName:   "mydatabase",
	}
}
 