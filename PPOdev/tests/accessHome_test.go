package tests_test

import (
	"testing"

	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	mocks_service "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg/repository/mocks"
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg/service"
	"github.com/golang/mock/gomock"
)

func TestGetListUserHome(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks_service.NewMockIAccessHomeRepo(ctrl)

	mockRepo.EXPECT().GetListUserHome(10).Return(nil, nil)

	accessService := service.NewAccessHomeService(mockRepo)

	list, err := accessService.GetListUserHome(10)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if list != nil {
		t.Errorf("Expected nil, got %v", list)
	}
}

func TestUpdateStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks_service.NewMockIAccessHomeRepo(ctrl)

	accessHome := pkg.AccessHome{
		AccessStatus: "active",
		AccessLevel:  2,
	}

	mockRepo.EXPECT().UpdateStatus(10, accessHome).Return(nil)

	accessService := service.NewAccessHomeService(mockRepo)

	err := accessService.UpdateStatus(10, accessHome)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestUpdateLevel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks_service.NewMockIAccessHomeRepo(ctrl)

	accessHome := pkg.AddUserHome{
		Email:       "active",
		AccessLevel: 2,
	}

	mockRepo.EXPECT().UpdateLevel(10, accessHome).Return(nil)

	accessService := service.NewAccessHomeService(mockRepo)

	err := accessService.UpdateLevel(10, accessHome)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks_service.NewMockIAccessHomeRepo(ctrl)

	mockRepo.EXPECT().DeleteUser(10, "dev1").Return(nil)

	accessService := service.NewAccessHomeService(mockRepo)

	err := accessService.DeleteUser(10, "dev1")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestAddUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks_service.NewMockIAccessHomeRepo(ctrl)

	userID := 1

	mockRepo.EXPECT().AddUser(userID, 2, "mail").Return(5, nil)

	accessService := service.NewAccessHomeService(mockRepo)

	accessID, err := accessService.AddUser(userID, 2, "mail")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if accessID != 5 {
		t.Errorf("Expected userID 5, got %d", accessID)
	}
}
