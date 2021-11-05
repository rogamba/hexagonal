package ports

import (
	"github.com/rogamba/hexagonal/domain/models"
)

type UserRepository interface {
	List() []models.User
	Get(userId string) *models.User
}
