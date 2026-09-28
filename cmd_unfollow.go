// SPDX-License-Identifier: GPL-3.0-only

package main

import (
	"context"
	"fmt"

	"github.com/Lauloque/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("'unfollow' expects a feed url")
	}

	feed, err := s.dbPtr.GetFeedFromUrl(context.Background(), cmd.arguments[0])
	if err != nil {
		return err
	}

	params := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	err = s.dbPtr.DeleteFeedFollow(context.Background(), params)
	if err != nil {
		return err
	}

	fmt.Printf(" * User '%v' unfollowed feed '%v'\n", user.Name, feed.Name)

	return nil
}
