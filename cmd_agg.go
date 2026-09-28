// SPDX-License-Identifier: GPL-3.0-only

package main

import (
	"context"
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("'agg' expects a duration string, eg: '1s', '1m', '1h' etc")
	}

	frequency, err := time.ParseDuration(cmd.arguments[0])
	if err != nil {
		return err
	}

	fmt.Printf("Collectiing feeds every %v", frequency)

	ticker := time.NewTicker(frequency)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}

	return nil
}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.dbPtr.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	err = s.dbPtr.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		return err
	}

	rssFeed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return err
	}

	printRssFeedTitles(rssFeed)

	return nil
}

func printRssFeedTitles(rssFeed *RSSFeed) {
	fmt.Println(rssFeed.Channel.Title)
	for i := range rssFeed.Channel.Item {
		fmt.Println(rssFeed.Channel.Item[i].Title)
	}
}
