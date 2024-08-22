package service

import (
	pkg "github.com/Mamvriyskiy/DBCourse/main/pkg"
	"github.com/Mamvriyskiy/DBCourse/main/pkg/repository"
)

type AccessHomeService struct {
	repo repository.IAccessHomeRepo
}

func NewAccessHomeService(repo repository.IAccessHomeRepo) *AccessHomeService {
	return &AccessHomeService{repo: repo}
}

func (s *AccessHomeService) AddOwner(userID, homeID int) (int, error) {
	return s.repo.AddOwner(userID, homeID)
}

func (s *AccessHomeService) AddUser(userID int, access pkg.Access) (int, error) {
	return s.repo.AddUser(userID, access)
}

func (s *AccessHomeService) DeleteUser(idUser int, access pkg.Access) error {
	return s.repo.DeleteUser(idUser, access)
}

func (s *AccessHomeService) UpdateLevel(idUser int, updateAccess pkg.Access) error {
	return s.repo.UpdateLevel(idUser, updateAccess)
}

func (s *AccessHomeService) UpdateStatus(idUser int, access pkg.AccessHome) error {
	return s.repo.UpdateStatus(idUser, access)
}

func (s *AccessHomeService) GetListUserHome(idHome int) ([]pkg.ClientHome, error) {
	return s.repo.GetListUserHome(idHome)
}
