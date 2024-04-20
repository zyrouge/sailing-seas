package core

import (
	"errors"
	"os"
	"strconv"
)

type Env struct {
	Host     string
	Port     int
	Username string
	Password string
}

func ParseEnv() (*Env, error) {
	host := os.Getenv("SS_HOST")
	port, err := strconv.Atoi(os.Getenv("SS_PORT"))
	if err != nil {
		return nil, errors.New("invalid env SS_PORT")
	}
	username := os.Getenv("SS_USERNAME")
	if username == "" {
		return nil, errors.New("invalid env SS_USERNAME")
	}
	passcode := os.Getenv("SS_PASSCODE")
	if passcode == "" {
		return nil, errors.New("invalid env SS_PASSCODE")
	}
	env := &Env{
		Host:     host,
		Port:     port,
		Username: username,
		Password: passcode,
	}
	return env, nil
}
