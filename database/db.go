package database

import (
	"callme/models"
	"gorm.io/gorm"
)

type postgresDB struct {
	databaseInterface
	db *gorm.DB
}

func (p *postgresDB) CreateUser(user *models.User) error {
	return p.db.Create(user).Error
}

func (p *postgresDB) GetUserByID(userID string) (*models.User, error) {
	user := new(models.User)
	err := p.db.Where("id = ?", userID).First(&user).Error
	return user, err
}

func (p *postgresDB) GetUserByEmail(email string) (*models.User, error) {
	user := new(models.User)
	err := p.db.Where("email = ?", email).First(&user).Error
	return user, err
}