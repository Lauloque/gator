// SPDX-License-Identifier: GPL-3.0-only

package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Lauloque/gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("Missing argument: 'register' expects a username")
	}

	inputName := cmd.arguments[0]

	_, err := s.dbPtr.GetUser(context.Background(), inputName)
	if err == nil {
		fmt.Printf("User '%s' already exists!", inputName)
		os.Exit(1)
	}

	currentTime := time.Now()
	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Name:      cmd.arguments[0],
	}
	user, err := s.dbPtr.CreateUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("Couldn't create user: %v", err)
	}

	s.configPtr.SetUser(user.Name)

	fmt.Printf("Added user: '%s'\n", s.configPtr.CurrentUserName)
	fmt.Printf("UUID:        %s\n", user.ID)
	fmt.Printf("Created at:  %v\n", user.CreatedAt)
	fmt.Printf("Updated at:  %v\n", user.UpdatedAt)

	return nil
}
