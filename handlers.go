package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jonryanedge/bootdev-gator/internal/config"
	"github.com/jonryanedge/bootdev-gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("login command require arguments")
	}

	user, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		return errors.New("user does not exist")
	}

	if err := s.cfg.SetUser(user.Name); err != nil {
		return errors.New("an error occured")
	}
	fmt.Println("User has been set")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("register command require arguments")
	}

	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	}

	user, err := s.db.CreateUser(context.Background(), userParams)
	if err != nil {
		return errors.New("could not create user")
	}

	if err := s.cfg.SetUser(user.Name); err != nil {
		return errors.New("an error occured setting user")
	}
	fmt.Println("User registered")
	fmt.Printf("User details: %v", user)

	return nil
}

func handlerResetUsers(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return errors.New("reset command requires no arguments")
	}

	if err := s.db.DeleteUsers(context.Background()); err != nil {
		return errors.New("could not delete users")
	}

	fmt.Println("All users deleted")

	return nil
}

func handlerGetUsers(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return errors.New("getUsers command requires no arguments")
	}

	currentUser := s.cfg.CurrentUserName

	registeredUsers, err := s.db.GetUsers(context.Background())
	if err != nil {
		return errors.New("could not retrieve user list")
	}

	for _, user := range registeredUsers {
		if user == currentUser {
			fmt.Printf("* %s (current)\n", user)
		}
		if user != currentUser {
			fmt.Printf("* %s\n", user)
		}
	}

	return nil
}
