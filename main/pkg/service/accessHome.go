package service

import (
	pkg "github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/Mamvriyskiy/database_course/main/pkg/repository"
)

type AccessHomeService struct {
	repo repository.IAccessHomeRepo
}

func NewAccessHomeService(repo repository.IAccessHomeRepo) *AccessHomeService {
	return &AccessHomeService{repo: repo}
}

func (s *AccessHomeService) AddOwner(userID, homeID string) (string, error) {
	return s.repo.AddOwner(userID, homeID)
}

func (s *AccessHomeService) AddUser(homeID string, access pkg.Access) (string, error) {
	return s.repo.AddUser(homeID, access)
}

func (s *AccessHomeService) DeleteUser(accessID string) error {
	return s.repo.DeleteUser(accessID)
}

func (s *AccessHomeService) UpdateLevel(accessID string, updateAccess pkg.Access) error {
	return s.repo.UpdateLevel(accessID, updateAccess)
}

func (s *AccessHomeService) UpdateStatus(userID string, access pkg.AccessHome) error {
	return s.repo.UpdateStatus(userID, access)
}

func (s *AccessHomeService) GetListUserHome(homeID string) ([]pkg.ClientHome, error) {
	return s.repo.GetListUserHome(homeID)
}

func (s *AccessHomeService) GetInfoAccessByID(accessID string) (pkg.Access, error) {
	return s.repo.GetInfoAccessByID(accessID)
}
