// SPDX-License-Identifier: GPL-3.0-only

package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Lauloque/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("'follow' expects a feed url")
	}

	feedUrl := cmd.arguments[0]
	feed, err := s.dbPtr.GetFeedFromUrl(context.Background(), feedUrl)
	if err != nil {
		return err
	}

	currentTime := time.Now()
	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		UserID:    user.ID,
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
