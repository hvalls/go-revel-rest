package repository

import (
	"go-revel-rest/app/models"
    "errors"
)

type UserRepository interface {
	GetUserById(id string) (*models.User, error)
	SaveUser(user *models.User) error
}

type DBUserRepository struct {
    users []*models.User
}

func New() *DBUserRepository {
    return &DBUserRepository {
        users : []*models.User {
            &models.User{ "1", "Hector" },
            &models.User{ "2", "Carlos" },
            &models.User{ "3", "Javi" },
            &models.User{ "4", "Dani" },
        },
    }
}

func (r *DBUserRepository) GetUserById(id string) (*models.User, error) {
    for _,user := range r.users {
        if user.Id == id {
            return user, nil
        }
    }
    return nil, errors.New("user not found")
}

func (r *DBUserRepository) SaveUser(user *models.User) error {
    r.users = append(r.users, user)
    return nil
}

var userRepository *DBUserRepository

func GetUserRepository() (r UserRepository) {
	if userRepository == nil {
		userRepository = New()
	}
	return userRepository
}