package tests_test

import (
	"testing"

	"github.com/Mamvriyskiy/DBCourse/main/pkg"
	mocks_service "github.com/Mamvriyskiy/DBCourse/main/pkg/repository/mocks"
	"github.com/Mamvriyskiy/DBCourse/main/pkg/service"
	"github.com/golang/mock/gomock"
)

func TestCreateHome(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks_service.NewMockIHomeRepo(ctrl)

	home := pkg.Home{
		Name:    "home",
		OwnerID: 20,
	}

	mockRepo.EXPECT().CreateHome(10, home).Return(5, nil)

	homeService := service.NewHomeService(mockRepo)

	homeID, err := homeService.CreateHome(10, home)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if homeID != 5 {
		t.Errorf("Expected userID 123, got %d", homeID)
	}
}

func TestDeleteHome(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks_service.NewMockIHomeRepo(ctrl)

	mockRepo.EXPECT().DeleteHome(10).Return(nil)

	homeService := service.NewHomeService(mockRepo)

	err := homeService.DeleteHome(10)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestUpdateHome(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks_service.NewMockIHomeRepo(ctrl)

	home := pkg.Home{
		ID:      1,
		Name:    "home",
		OwnerID: 20,
	}

	mockRepo.EXPECT().UpdateHome(home).Return(nil)

	homeService := service.NewHomeService(mockRepo)

	err := homeService.UpdateHome(home)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
