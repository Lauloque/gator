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
