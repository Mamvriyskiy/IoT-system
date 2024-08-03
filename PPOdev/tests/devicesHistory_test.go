package tests_test

import (
	"testing"

	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	mocks_service "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg/repository/mocks"
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg/service"
	"github.com/golang/mock/gomock"
)

func TestCreateDeviceHistory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks_service.NewMockIHistoryDeviceRepo(ctrl)

	history := pkg.AddHistory{
		Name:             "dev1",
		TimeWork:         200,
		AverageIndicator: 66,
		EnergyConsumed:   100,
	}

	mockRepo.EXPECT().CreateDeviceHistory(10, history).Return(50, nil)

	historyService := service.NewHistoryDeviceService(mockRepo)

	historyID, err := historyService.CreateDeviceHistory(10, history)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if historyID != 50 {
		t.Errorf("Expected userID 50, got %d", historyID)
	}
}

// GetDeviceHistory(idDevice int) ([]pkg.DevicesHistory, error).
func TestGetDeviceHistory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks_service.NewMockIHistoryDeviceRepo(ctrl)

	mockRepo.EXPECT().GetDeviceHistory(10, "dev1").Return(nil, nil)

	historyService := service.NewHistoryDeviceService(mockRepo)

	list, err := historyService.GetDeviceHistory(10, "dev1")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if list != nil {
		t.Errorf("Expected no nil, got %v", err)
	}
}
