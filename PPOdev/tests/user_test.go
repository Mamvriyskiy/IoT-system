package tests_test

import (
	"testing"

	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	mocks_service "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg/repository/mocks"
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg/service"
	"github.com/golang/mock/gomock"
)

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks_service.NewMockIUserRepo(ctrl)

	user := pkg.User{
		Username: "username",
		Email:    "qwerty@mail.ru",
		Password: "qwerty",
	}

	mockRepo.EXPECT().CreateUser(gomock.Any()).Return(123, nil)

	userService := service.NewUserService(mockRepo)

	userID, err := userService.CreateUser(user)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if userID != 123 {
		t.Errorf("Expected userID 123, got %d", userID)
	}
}

func TestCheckUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks_service.NewMockIUserRepo(ctrl)

	user := pkg.User{
		Username: "username",
		Email:    "qwerty@mail.ru",
		Password: "qwerty",
		ID:       123,
	}

	mockRepo.EXPECT().GetUser(user.Username, user.Password).Return(user, nil)

	userService := service.NewUserService(mockRepo)

	userID, err := userService.CheckUser(user)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if userID != 123 {
		t.Errorf("Expected userID 123, got %d", userID)
	}
}
