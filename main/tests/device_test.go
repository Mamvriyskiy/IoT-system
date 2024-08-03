package tests_test

import (
	"testing"

	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	mocks_service "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg/repository/mocks"
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg/service"
	"github.com/golang/mock/gomock"
)

func TestCreateDevice(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks_service.NewMockIDeviceRepo(ctrl)

	device := pkg.Devices{
		Name:             "tea",
		TypeDevice:       "kettle",
		Status:           "free",
		Brand:            "Samsung",
		PowerConsumption: 1500,
		MinParameter:     50, // temperature
		MaxParameter:     120,
	}

	homeID := 1
	mockRepo.EXPECT().CreateDevice(homeID, &device).Return(10, nil)

	deviceService := service.NewDeviceService(mockRepo)

	deviceID, err := deviceService.CreateDevice(homeID, &device)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if deviceID != 10 {
		t.Errorf("Expected userID 10, got %d", deviceID)
	}
}

func TestDeleteDevice(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks_service.NewMockIDeviceRepo(ctrl)

	mockRepo.EXPECT().DeleteDevice(10, "dev1").Return(nil)

	deviceService := service.NewDeviceService(mockRepo)

	err := deviceService.DeleteDevice(10, "dev1")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestDeleteHomeDevice(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks_service.NewMockIDeviceRepo(ctrl)

	mockRepo.EXPECT().DeleteDevice(10, "dev1").Return(nil)

	deviceService := service.NewDeviceService(mockRepo)

	err := deviceService.DeleteDevice(10, "dev1")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
