package main

import (
	"errors"
	"fmt"

	"github.com/jonryanedge/bootdev-gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("login command require arguments")
	}

	if err := s.cfg.SetUser(cmd.args[0]); err != nil {
		return errors.New("an error occured")
	}
	fmt.Println("User has been set")
	return nil
}
