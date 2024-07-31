package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/wlucjan/rssagg/internal/database"
)

func startScraping(db *database.Queries, concurrency int, timeBetweenReqiest time.Duration) {
	log.Printf("Scraping on %v goroutines with %v between requests", concurrency, timeBetweenReqiest)

	ticker := time.NewTicker(timeBetweenReqiest)
	for ; ; <-ticker.C {
		log.Printf("Scraping...")
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Printf("Error fetching feeds: %v", err)
			continue
		}

		wg := &sync.WaitGroup{}

		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()

	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()
	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Error marking feed as fetched: %v", err)
		return
	}

	rssFeed, err := URLToFeed(feed.Url)
	if err != nil {
		log.Printf("Error fetching feed: %v", err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		log.Println("Found post", item.Title, "on feed", feed.Name)
	}

	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}
