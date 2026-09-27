// SPDX-License-Identifier: GPL-3.0-only

package main

import (
	"context"
	"fmt"

	"github.com/Lauloque/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	feedFollows, err := s.dbPtr.GetFeedFollowsForUserId(context.Background(), user.ID)
	if err != nil {
		return err
	}

	if len(feedFollows) == 0 {
		fmt.Printf("No feeds followed by user '%v', yet...\n", user.Name)
		return nil
	}

	fmt.Printf("Feeds followed by user '%v':\n", user.Name)
	for _, feedFollow := range feedFollows {
		fmt.Printf("* Feed Name: %v\n", feedFollow.FeedName)
	}

	return nil
}
