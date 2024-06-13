package interfaces

import "github.com/imjowend/patron-dao/models"

type UserDAO interface {
	Create(u *models.User) error
	GetAll() ([]models.User, error)
}
