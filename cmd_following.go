// SPDX-License-Identifier: GPL-3.0-only

package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {

	currentUser, err := s.dbPtr.GetUser(context.Background(), s.configPtr.CurrentUserName)
	if err != nil {
		return err
	}

	feedFollows, err := s.dbPtr.GetFeedFollowsForUserId(context.Background(), currentUser.ID)
	if err != nil {
		return err
	}

	if len(feedFollows) == 0 {
		fmt.Printf("No feeds followed by user '%v', yet...\n", currentUser.Name)
		return nil
	}

	fmt.Printf("Feeds followed by user '%v':\n", currentUser.Name)
	for _, feedFollow := range feedFollows {
		fmt.Printf("* Feed Name: %v\n", feedFollow.FeedName)
	}

	return nil
}
