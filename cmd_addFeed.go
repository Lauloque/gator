// SPDX-License-Identifier: GPL-3.0-only

package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Lauloque/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 2 {
		return fmt.Errorf("'addfeed' expects a feed name and a url")
	}

	feedName := cmd.arguments[0]
	feedUrl := cmd.arguments[1]

	currentTime := time.Now()
	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Name:      feedName,
		Url:       feedUrl,
		UserID:    user.ID,
	}

	feed, err := s.dbPtr.CreateFeed(context.Background(), params)
	if err != nil {
		return err
	}

	params2 := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	_, err = s.dbPtr.CreateFeedFollow(context.Background(), params2)
	if err != nil {
		return err
	}

	fmt.Printf(" Created feed and followed it successfully\n")
	fmt.Printf(" * ID         : %v\n", feed.ID)
	fmt.Printf(" * Name       : %v\n", feed.Name)
	fmt.Printf(" * Created At : %v\n", feed.CreatedAt)
	fmt.Printf(" * Updated At : %v\n", feed.UpdatedAt)
	fmt.Printf(" * URL        : %v\n", feed.Url)
	fmt.Printf(" * User ID    : %v\n", feed.UserID)

	return nil
}
