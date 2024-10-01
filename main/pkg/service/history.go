package service

import (
	pkg "github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/Mamvriyskiy/database_course/main/pkg/repository"
	"crypto/rand"
	"errors"
	"math/big"

	"github.com/Mamvriyskiy/database_course/main/logger"
)

var ErrNoFloat64Interface = errors.New("отсутствует интерфейс {} для float64")

type HistoryDeviceService struct {
	repo repository.IHistoryDeviceRepo
}

func NewHistoryDeviceService(repo repository.IHistoryDeviceRepo) *HistoryDeviceService {
	return &HistoryDeviceService{repo: repo}
}

func generateRandomInt(max int) int {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		logger.Log("Error", "userID.(float64)", "Error:", ErrNoFloat64Interface, "")
		return 0
	}
	return int(n.Int64())
}

func generateRandomFloat(max float64) float64 {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max*100)))
	if err != nil {
		logger.Log("Error", "userID.(float64)", "Error:", ErrNoFloat64Interface, "")
		return 0.0
	}
	return float64(n.Int64()) / 100.0
}

func (s *HistoryDeviceService) CreateDeviceHistory(deviceID int,
	history pkg.AddHistory,
) (int, error) {
	history.TimeWork = generateRandomInt(101)
	history.AverageIndicator = generateRandomFloat(100)
	history.EnergyConsumed = generateRandomInt(101)

	return s.repo.CreateDeviceHistory(deviceID, history)
}

func (s *HistoryDeviceService) GetDeviceHistory(userID int,
	name, home string) ([]pkg.DevicesHistory, error) {
	return s.repo.GetDeviceHistory(userID, name, home)
}
