package usersservice

import (
	"github.com/godvlpr/template/config"
	"github.com/godvlpr/template/models"
	"github.com/godvlpr/template/storage"
	"github.com/godvlpr/template/storage/db"
	"github.com/sirupsen/logrus"
)

const (
	serviceName = "USER"
)

type UserService interface {
	ServiceName() string
	Create(user models.User) (models.User, error)
	IncrementUserAge(user models.User) models.User
	CreateFullName(user models.User) models.User
}

type Service struct {
	db  db.QInterface
	log *logrus.Entry
}

func New(storageProvider *storage.StorageProvider) UserService {
	return &Service{
		db:  storageProvider.DB,
		log: config.GetConfig().GetLogCopy().WithField("SERVICE", serviceName),
	}
}

func (s Service) ServiceName() string {
	return serviceName
}
