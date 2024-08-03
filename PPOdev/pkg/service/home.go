package service

import (
	pkg "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg/repository"
)

type HomeService struct {
	repo repository.IHomeRepo
}

func NewHomeService(repo repository.IHomeRepo) *HomeService {
	return &HomeService{repo: repo}
}

func (s *HomeService) CreateHome(idUser int, home pkg.Home) (int, error) {
	return s.repo.CreateHome(idUser, home)
}

func (s *HomeService) DeleteHome(homeID int) error {
	return s.repo.DeleteHome(homeID)
}

func (s *HomeService) UpdateHome(home pkg.Home) error {
	return s.repo.UpdateHome(home)
}

func (s *HomeService) GetHomeByID(homeID int) (pkg.Home, error) {
	return s.repo.GetHomeByID(homeID)
}

func (s *HomeService) ListUserHome(userID int) ([]pkg.Home, error) {
	return s.repo.ListUserHome(userID)
}
