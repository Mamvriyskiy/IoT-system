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

func (s *AccessHomeService) AddUser(homeID string, access pkg.AccessHandler) (string, error) {
	accessService := pkg.AccessService{
		Access: access.Access,
	}
	return s.repo.AddUser(homeID, accessService)
}

func (s *AccessHomeService) DeleteUser(accessID string) error {
	return s.repo.DeleteUser(accessID)
}

func (s *AccessHomeService) UpdateLevel(accessID string, access pkg.AccessHandler) error {
	updateAccess := pkg.AccessService{
		Access: access.Access,
	}
	return s.repo.UpdateLevel(accessID, updateAccess)
}

func (s *AccessHomeService) UpdateStatus(userID string, access pkg.AccessHandler) error {
	updateAccess := pkg.AccessService{
		Access: access.Access,
	}
	return s.repo.UpdateStatus(userID, updateAccess)
}

func (s *AccessHomeService) GetListUserHome(homeID string) ([]pkg.AccessInfoData, error) {
	return s.repo.GetListUserHome(homeID)
}

func (s *AccessHomeService) GetInfoAccessByID(accessID string) (pkg.AccessInfoData, error) {
	return s.repo.GetInfoAccessByID(accessID)
}
