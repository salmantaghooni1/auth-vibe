package repository

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/salmantaghooni/golang-car-web-api/api/dto"
	"github.com/salmantaghooni/golang-car-web-api/config"
	"github.com/salmantaghooni/golang-car-web-api/constants"
	"github.com/salmantaghooni/golang-car-web-api/data/db"
	"github.com/salmantaghooni/golang-car-web-api/data/models"
	"github.com/salmantaghooni/golang-car-web-api/pkg/logging"
	"github.com/salmantaghooni/golang-car-web-api/pkg/service_errors"
)

type UserRepository struct {
	logger   logging.Logger
	cfg      *config.Config
	database *gorm.DB
}

func NewUserRepository(cfg *config.Config) *UserRepository {
	database := db.GetDb()
	logger := logging.NewLogger(cfg)
	return &UserRepository{logger: logger, cfg: config.GetConfig(), database: database}
}

// Register by username

func (s *UserRepository) RegisterByGoogle(req *dto.RegisterUserByUsernameRequest) error {
	u := models.User{Username: req.Username, FirstName: req.FirstName, LastName: req.LastName, Email: req.Email}

	exists, err := s.existsByEmail(req.Email)
	if err != nil {
		return err
	}
	if exists {
		return &service_errors.ServiceError{EndUserMessage: service_errors.EmailExists}
	}
	exists, err = s.existsByUsername(req.Username)
	if err != nil {
		return err
	}
	if exists {
		return &service_errors.ServiceError{EndUserMessage: service_errors.UsernameExists}
	}

	bp := []byte(req.Password)
	hp, err := bcrypt.GenerateFromPassword(bp, bcrypt.DefaultCost)
	if err != nil {
		s.logger.Error(logging.General, logging.HashPassword, err.Error(), nil)
		return err
	}
	u.Password = string(hp)
	roleId, err := s.getDefaultRole()
	if err != nil {
		s.logger.Error(logging.Postgres, logging.DefaultRoleNotFound, err.Error(), nil)
		return err
	}

	tx := s.database.Begin()
	err = tx.Create(&u).Error
	if err != nil {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Rollback, err.Error(), nil)
		return err
	}
	err = tx.Create(&models.UserRole{RoleId: roleId, UserId: u.Id}).Error
	if err != nil {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Rollback, err.Error(), nil)
		return err
	}
	tx.Commit()
	return nil

}

func (s *UserRepository) existsByEmail(email string) (bool, error) {
	var exists bool
	if err := s.database.Model(&models.User{}).
		Select("count(*) > 0").
		Where("email = ?", email).
		Find(&exists).
		Error; err != nil {
		s.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return false, err
	}
	return exists, nil
}

func (s *UserRepository) existsByUsername(username string) (bool, error) {
	var exists bool
	if err := s.database.Model(&models.User{}).
		Select("count(*) > 0").
		Where("username = ?", username).
		Find(&exists).
		Error; err != nil {
		s.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return false, err
	}
	return exists, nil
}

func (s *UserRepository) getDefaultRole() (roleId int, err error) {

	if err = s.database.Model(&models.Role{}).
		Select("id").
		Where("name = ?", constants.DefaultRoleName).
		First(&roleId).Error; err != nil {
		return 0, err
	}
	return roleId, nil
}
