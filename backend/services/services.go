package services

import (
	"github.com/thisusami/thaibrevquiz/models"
	"github.com/thisusami/thaibrevquiz/repositories"
)

type Service struct {
	Repo *repositories.Repository
}

func (s *Service) RegisterService(user *models.User) (*models.User, error) {
	result, err := s.Repo.Insert(user)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (s *Service) LoginService(user *models.User) (*models.User, error) {
	result, err := s.Repo.Get(user)
	if result == nil {
		return nil, err
	}
	return result, nil
}
func NewService(repo *repositories.Repository) *Service {
	return &Service{repo}
}
