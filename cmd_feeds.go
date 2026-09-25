// SPDX-License-Identifier: GPL-3.0-only

package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, _ command) error {

	feeds, err := s.dbPtr.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	if len(feeds) == 0 {
		fmt.Println("Feeds table is empty.")
		return nil
	}

	for _, feed := range feeds {
		fmt.Printf("* Name    : %s\n", feed.Name)
		fmt.Printf("* URL     : %s\n", feed.Url)
		author, err := s.dbPtr.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return err
		}
		fmt.Printf("* Author  : %s\n", author.Name)
	}

	return nil
}
