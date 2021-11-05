package memory

import (
	"github.com/rogamba/hexagonal/domain/models"
	"github.com/rogamba/hexagonal/domain/ports"
)

// Implementation of the user port (repository)

var usersData []models.User

type memoryUserRepository struct {
	logLevel string
}

func NewMemoryUserRepository() (ports.UserRepository, error) {
	userRepository := &memoryUserRepository{
		logLevel: "INFO",
	}
	return userRepository, nil
}

func (u *memoryUserRepository) Get(userId string) *models.User {
	// Get user from memory
	return &models.User{
		UserId: "1",
		Name:   "Rodrigo",
	}
}

func (u *memoryUserRepository) List() []models.User {
	return []models.User{}
}
