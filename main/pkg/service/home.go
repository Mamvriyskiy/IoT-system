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

func (s *HomeService) CreateHome(home pkg.HomeHandler) (pkg.HomeData, error) {
	homeService := pkg.HomeService{
		Home: home.Home,
	}

	homeID, err := s.repo.CreateHome(homeService)
	if err != nil {
		return pkg.HomeData{}, err
	}

	return s.repo.GetHomeByID(homeID)
}

func (s *HomeService) DeleteHome(homeID string) error {
	return s.repo.DeleteHome(homeID)
}

func (s *HomeService) UpdateHome(homeID, name string) error {
	return s.repo.UpdateHome(homeID, name)
}

func (s *HomeService) GetHomeByID(homeID string) (pkg.HomeData, error) {
	return s.repo.GetHomeByID(homeID)
}

func (s *HomeService) ListUserHome(userID string) ([]pkg.HomeData, error) {
	return s.repo.ListUserHome(userID)
}
