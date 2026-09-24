// SPDX-License-Identifier: GPL-3.0-only

package main

import (
	"context"
	"fmt"
	"os"
)

func handlerUsers(s *state, cmd command) error {

	users, err := s.dbPtr.GetUsers(context.Background())
	if err != nil {
		fmt.Printf("Can't get user list: %v", err)
		os.Exit(1)
	}

	if len(users) == 0 {
		fmt.Println("User table is empty.")
		return nil
	}

	for _, user := range users {
		if user.Name == s.configPtr.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}

	return nil
}
