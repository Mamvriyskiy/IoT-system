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

func (s *HomeService) CreateHome(home pkg.Home) (int, error) {
	return s.repo.CreateHome(home)
}

func (s *HomeService) DeleteHome(homeID int, homeName string) error {
	return s.repo.DeleteHome(homeID, homeName)
}

func (s *HomeService) UpdateHome(home pkg.UpdateNameHome) error {
	return s.repo.UpdateHome(home)
}

func (s *HomeService) GetHomeByID(homeID int) (pkg.Home, error) {
	return s.repo.GetHomeByID(homeID)
}

func (s *HomeService) ListUserHome(userID int) ([]pkg.Home, error) {
	return s.repo.ListUserHome(userID)
}
