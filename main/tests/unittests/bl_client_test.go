package unittests

import (
	//"github.com/Mamvriyskiy/database_course/main/pkg"
	mocks_service "github.com/Mamvriyskiy/database_course/main/pkg/repository/mocks"
	"github.com/Mamvriyskiy/database_course/main/pkg/service"
	"github.com/Mamvriyskiy/database_course/main/tests/factory"
	method "github.com/Mamvriyskiy/database_course/main/tests/method"
	"github.com/golang/mock/gomock"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"errors"
)

func (s *MyUnitTestsSuite) TestCreateClientBL(t provider.T) {
	tests := []struct {
		nameTest string
		user     factory.ObjectSystem
		userID   string
	}{
		{
			nameTest: "Test1",
			user:     factory.New("user", ""),
			userID:   "1",
		},
		{
			nameTest: "Test2",
			user:     factory.New("user", ""),
			userID:   "2",
		},
		{
			nameTest: "Test3",
			user:     factory.New("user", ""),
			userID:   "3",
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks_service.NewMockIUserRepo(ctrl)

	for _, test := range tests {
		t.Run(test.nameTest, func(t provider.T) {
			newUser := test.user.(*method.TestUser)

			mockRepo.EXPECT().CreateUser(gomock.Any()).Return(test.userID, nil)

			userService := service.NewUserService(mockRepo)

			userID, err := userService.CreateUser(newUser.User)

			t.Assert().NoError(err)
			t.Assert().Equal(test.userID, userID)
		})
	}
}

func (s *MyUnitTestsSuite) TestCheckUserBL(t provider.T) {
	tests := []struct {
		nameTest string
		user     factory.ObjectSystem
		userID   string
	}{
		{
			nameTest: "Test1",
			user:     factory.New("user", ""),
			userID:   "1",
		},
		{
			nameTest: "Test2",
			user:     factory.New("user", ""),
			userID:   "2",
		},
		{
			nameTest: "Test3",
			user:     factory.New("user", ""),
			userID:   "3",
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создаем mock для репозитория
	mockRepo := mocks_service.NewMockIUserRepo(ctrl)

	for _, test := range tests {
		t.Run(test.nameTest, func(t provider.T) {
			newUser := test.user.(*method.TestUser)
			newUser.ID = test.userID

			mockRepo.EXPECT().GetUser(newUser.User.Email, newUser.User.Password).Return(newUser.User, nil)

			userService := service.NewUserService(mockRepo)

			resultID, err := userService.CheckUser(newUser.User)

			t.Assert().NoError(err)
			t.Assert().Equal(resultID, test.userID)
		})
	}
}

func (s *MyUnitTestsSuite) TestSendCodeBL(t provider.T) {
	tests := []struct {
		nameTest string
		email    factory.ObjectSystem
		userID   int
	}{
		{
			nameTest: "Test1",
			email:    factory.New("email", ""),
		},
		{
			nameTest: "Test2",
			email:    factory.New("email", ""),
		},
		{
			nameTest: "Test3",
			email:    factory.New("email", ""),
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создаем mock для репозитория
	mockRepo := mocks_service.NewMockIUserRepo(ctrl)

	for _, test := range tests {
		t.Run(test.nameTest, func(t provider.T) {
			newEmail := test.email.(*method.TestEmail)

			mockRepo.EXPECT().AddCode(gomock.Any()).Return(nil)

			userService := service.NewUserService(mockRepo)

			err := userService.SendCode(newEmail.Email)

			t.Assert().NoError(err)
		})
	}
}

func (s *MyUnitTestsSuite) TestCheckCodeBL(t provider.T) {
	tests := []struct {
		nameTest   string
		code       string
		token      string
		resetCode  string
		checkError error
	}{
		{
			nameTest:   "Test1",
			code:       "123456",
			token:      "fsdjkgjksfjivniusdnivniusdnuigniusdngiunsdinfgiwr",
			resetCode:  "123456",
			checkError: nil,
		},
		{
			nameTest:   "Test2",
			code:       "2352523",
			token:      "kasdlfjsdncjhnhsbnfhbdshngfjndsjhgnjhwrgnjsdnfjhn",
			resetCode:  "2352523",
			checkError: nil,
		},
		{
			nameTest:   "Test3",
			code:       "67445735",
			token:      "gsdhgsjkanjcksnjngejhfnghndsjnfnsdjnfjnsdjknfsdnj",
			resetCode:  "1324134",
			checkError: errors.New("Negative code"),
		},
	}

	errors.New("Negative code")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks_service.NewMockIUserRepo(ctrl)

	for _, test := range tests {
		t.Run(test.nameTest, func(t provider.T) {
			mockRepo.EXPECT().GetCode(test.token).Return(test.resetCode, nil)

			userService := service.NewUserService(mockRepo)

			err := userService.CheckCode(test.code, test.token)

			t.Assert().Equal(test.checkError, err)
		})
	}
}
