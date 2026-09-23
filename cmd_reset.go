// SPDX-License-Identifier: GPL-3.0-only

package main

import (
	"context"
	"fmt"
	"os"
)

func handlerReset(s *state, _ command) error {
	err := s.dbPtr.NukeUsers(context.Background())
	if err != nil {
		fmt.Printf("Couldn't delete all users: %v", err)
		os.Exit(1)
	}

	fmt.Println("Now I am become Death, the destroyer of users.")
	return nil
}
