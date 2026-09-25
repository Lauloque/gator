// SPDX-License-Identifier: GPL-3.0-only

package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, _ command) error {
	url := "https://www.wagslane.dev/index.xml"

	feed, err := fetchFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Couldn't fetch feed: %v", err)
	}

	fmt.Printf("%+v\n", feed)

	return nil
}
