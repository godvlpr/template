package usersservice

import "github.com/godvlpr/template/models"

func (s *Service) Create(user models.User) (models.User, error) {
	user = s.CreateFullName(user)

	err := s.db.UsersQ().Create(user)
	if err != nil {
		s.log.WithError(err).Error("failed to create user")
		return user, err
	}
	return user, nil
}

func (s *Service) IncrementUserAge(user models.User) models.User {
	user.Age++
	return user
}

func (s *Service) CreateFullName(user models.User) models.User {
	user.FullName = user.FirstName + user.LastName
	return user
}
