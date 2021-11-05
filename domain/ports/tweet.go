package ports

import (
	"github.com/rogamba/hexagonal/domain/models"
)

type TweetRepository interface {
	List() []models.Tweet
	ListByUser(userId string) []models.Tweet
}
