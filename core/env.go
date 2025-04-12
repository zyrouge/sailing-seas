package core

import (
	"errors"
	"os"
	"strconv"
)

type Env struct {
	Host string
	Port int
}

func ParseEnv() (*Env, error) {
	host := os.Getenv("SS_HOST")
	port, err := strconv.Atoi(os.Getenv("SS_PORT"))
	if err != nil {
		return nil, errors.New("invalid env SS_PORT")
	}
	env := &Env{
		Host: host,
		Port: port,
	}
	return env, nil
}
