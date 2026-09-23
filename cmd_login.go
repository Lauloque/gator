// SPDX-License-Identifier: GPL-3.0-only

package main

import (
	"context"
	"fmt"
	"os"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("Missing argument: 'login' expects a username")
	}

	inputName := cmd.arguments[0]

	_, err := s.dbPtr.GetUser(context.Background(), inputName)
	if err != nil {
		fmt.Printf("Can't login as user '%s': %v", inputName, err)
		os.Exit(1)
	}

	err = s.configPtr.SetUser(inputName)
	if err != nil {
		return err
	}

	fmt.Printf("Username set to '%s'", s.configPtr.CurrentUserName)

	return nil
}
