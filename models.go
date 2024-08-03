package main

import (
	"time"

	"github.com/google/uuid"

	"github.com/wlucjan/rssagg/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func databaseUserToAPIUser(user database.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		ApiKey:    user.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
}

func databaseFeedToAPIFeed(feed database.Feed) Feed {
	return Feed{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name:      feed.Name,
		Url:       feed.Url,
	}
}

func databaseFeedsToAPIFeeds(feeds []database.Feed) []Feed {
	apiFeeds := make([]Feed, len(feeds))
	for i, feed := range feeds {
		apiFeeds[i] = databaseFeedToAPIFeed(feed)
	}

	return apiFeeds
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func databaseFeedFollowToAPIFeedFollow(follow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        follow.ID,
		CreatedAt: follow.CreatedAt,
		UpdatedAt: follow.UpdatedAt,
		UserID:    follow.UserID,
		FeedID:    follow.FeedID,
	}
}

func databaseFeedFollowsToAPIFeedFollows(follows []database.FeedFollow) []FeedFollow {
	apiFollows := make([]FeedFollow, len(follows))
	for i, follow := range follows {
		apiFollows[i] = databaseFeedFollowToAPIFeedFollow(follow)
	}

	return apiFollows
}

type Post struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	URL         string    `json:"url"`
	FeedID      string    `json:"feed_id"`
}

func databasePostToAPIPost(post database.Post) Post {
	var description string
	if post.Description.Valid {
		description = post.Description.String
	}

	var publishedAt time.Time
	if post.PublishedAt.Valid {
		publishedAt = post.PublishedAt.Time
	}

	return Post{
		ID:          post.ID.String(),
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
		Title:       post.Title,
		Description: description,
		PublishedAt: publishedAt,
		URL:         post.Url,
		FeedID:      post.FeedID.String(),
	}
}

func databasePostsToAPIPosts(posts []database.Post) []Post {
	apiPosts := make([]Post, len(posts))
	for i, post := range posts {
		apiPosts[i] = databasePostToAPIPost(post)
	}

	return apiPosts
}
