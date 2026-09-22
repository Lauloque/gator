// SPDX-License-Identifier: GPL-3.0-only

package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("Missing argument: 'login' expects a username")
	}

	err := s.configPtr.SetUser(cmd.arguments[0])
	if err != nil {
		return err
	}

	fmt.Printf("Username set to '%s'", s.configPtr.CurrentUserName)

	return nil
}
