// SPDX-License-Identifier: GPL-3.0-only

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Lauloque/gator/internal/database"
	"github.com/google/uuid"
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

	fmt.Println(rssFeed.Channel.Title)

	for i := range rssFeed.Channel.Item {
		description := sql.NullString{
			String: rssFeed.Channel.Item[i].Description,
			Valid:  true,
		}

		rawDate := rssFeed.Channel.Item[i].PubDate
		pubdate := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, rawDate); err == nil {
			pubdate = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		} else if t, err := time.Parse(time.RFC1123, rawDate); err == nil {
			pubdate = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}

		currentTime := time.Now()
		params := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   currentTime,
			UpdatedAt:   currentTime,
			Title:       rssFeed.Channel.Item[i].Title,
			Url:         rssFeed.Channel.Item[i].Link,
			Description: description,
			PublishedAt: pubdate,
			FeedID:      nextFeed.ID,
		}

		_, err := s.dbPtr.CreatePost(context.Background(), params)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("Couldn't create post: %v\n", err)
			continue
		}
		log.Printf("Created post: %v\n", rssFeed.Channel.Item[i].Title)
	}

	return nil
}
