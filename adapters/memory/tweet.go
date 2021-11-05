package memory

import (
	"github.com/rogamba/hexagonal/domain/models"
	"github.com/rogamba/hexagonal/domain/ports"
)

// Implementation of the tweet port (repository)

var tweetData []models.Tweet

type memoryTweetRepository struct {
	logLevel string
}

func NewMemoryTweetRepository() (ports.TweetRepository, error) {
	tweetRepository := &memoryTweetRepository{
		logLevel: "INFO",
	}
	return tweetRepository, nil
}

func (u *memoryTweetRepository) List() []models.Tweet {
	return []models.Tweet{
		{
			UserId: "1",
			Title:  "ExampleTweet",
			Body:   "This is an example tweet!",
		},
		{
			UserId: "2",
			Title:  "Another tweet",
			Body:   "Lorem ipum domain ergo subito",
		},
	}
}

func (u *memoryTweetRepository) ListByUser(userId string) []models.Tweet {
	return []models.Tweet{
		{
			UserId: "1",
			Title:  "ExampleTweet",
			Body:   "This is an example tweet!",
		},
	}
}
