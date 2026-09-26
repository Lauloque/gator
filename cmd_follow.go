// SPDX-License-Identifier: GPL-3.0-only

package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Lauloque/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("'follow' expects a feed url")
	}

	feedUrl := cmd.arguments[0]
	feed, err := s.dbPtr.GetFeedFromUrl(context.Background(), feedUrl)
	if err != nil {
		return err
	}

	currentUser, err := s.dbPtr.GetUser(context.Background(), s.configPtr.CurrentUserName)
	if err != nil {
		return err
	}

	currentTime := time.Now()
	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		UserID:    currentUser.ID,
		FeedID:    feed.ID,
	}

	feedFollow, err := s.dbPtr.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return err
	}

	fmt.Printf(" * Feed Name : %v\n", feedFollow.FeedName)
	fmt.Printf(" * User Name : %v\n", feedFollow.UserName)

	return nil
}
