package service

import (
	pkg "github.com/Mamvriyskiy/DBCourse/main/pkg"
	"github.com/Mamvriyskiy/DBCourse/main/pkg/repository"
)

type HistoryDeviceService struct {
	repo repository.IHistoryDeviceRepo
}

func NewHistoryDeviceService(repo repository.IHistoryDeviceRepo) *HistoryDeviceService {
	return &HistoryDeviceService{repo: repo}
}

func (s *HistoryDeviceService) CreateDeviceHistory(deviceID int,
	history pkg.AddHistory,
) (int, error) {
	return s.repo.CreateDeviceHistory(deviceID, history)
}

func (s *HistoryDeviceService) GetDeviceHistory(userID int,
	name string,
) ([]pkg.DevicesHistory, error) {
	return s.repo.GetDeviceHistory(userID, name)
}
