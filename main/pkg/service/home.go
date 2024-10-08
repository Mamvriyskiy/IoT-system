package service

import (
	pkg "github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/Mamvriyskiy/database_course/main/pkg/repository"
)

type HomeService struct {
	repo repository.IHomeRepo
}

func NewHomeService(repo repository.IHomeRepo) *HomeService {
	return &HomeService{repo: repo}
}

func (s *HomeService) CreateHome(home pkg.Home) (string, error) {
	return s.repo.CreateHome(home)
}

func (s *HomeService) DeleteHome(homeID string) error {
	return s.repo.DeleteHome(homeID)
}

func (s *HomeService) UpdateHome(homeID, name string) error {
	return s.repo.UpdateHome(homeID, name)
}

func (s *HomeService) GetHomeByID(homeID string) (pkg.Home, error) {
	return s.repo.GetHomeByID(homeID)
}

func (s *HomeService) ListUserHome(userID string) ([]pkg.Home, error) {
	return s.repo.ListUserHome(userID)
}
