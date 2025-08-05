package application

import "gin/internal/domain/repository"

type UserService struct {
	repository repository.UserRepository
}

func (s *UserService) GetAllUsers() (..., error) {
	users, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}