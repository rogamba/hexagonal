package services

import (
	"github.com/rogamba/hexagonal/domain/models"
	"github.com/rogamba/hexagonal/domain/ports"
)

type UserService interface {
	FetchUserDetails(userId string) (*models.User, error)
	StoreUser(user *models.User) error
}

type userService struct {
	userStorage  ports.UserRepository
	tweetStorage ports.TweetRepository
}

func NewUserService(
	userStorage ports.UserRepository,
	tweetStorage ports.TweetRepository,
) UserService {
	return &userService{
		userStorage:  userStorage,
		tweetStorage: tweetStorage,
	}
}

// Gets the user details and append it's tweets
func (u *userService) FetchUserDetails(userId string) (*models.User, error) {
	user := u.userStorage.Get(userId)
	tweets := u.tweetStorage.ListByUser(userId)
	user.Tweets = tweets
	return user, nil
}

// Store user to wherever storage we have injected
func (u *userService) StoreUser(user *models.User) error {
	return nil
}
