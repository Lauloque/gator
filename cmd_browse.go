// SPDX-License-Identifier: GPL-3.0-only

package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Lauloque/gator/internal/database"
)

func handlerBrowse(s *state, cmd command) error {
	limit := 2
	var err error
	if len(cmd.arguments) > 0 {
		limit, err = strconv.Atoi(cmd.arguments[0])
		if err != nil {
			return err
		}
	}

	currentUser, err := s.dbPtr.GetUser(context.Background(), s.configPtr.CurrentUserName)
	if err != nil {
		return err
	}

	params := database.GetPostsForUserParams{
		UserID: currentUser.ID,
		Limit:  int32(limit),
	}

	followedPosts, err := s.dbPtr.GetPostsForUser(context.Background(), params)
	if err != nil {
		return err
	}

	for _, post := range followedPosts {
		fmt.Printf("Title: %s\nURL: %s\n\n", post.Title, post.Url)
	}

	return nil
}
