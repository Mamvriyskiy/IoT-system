package unittests

import (
	"github.com/Mamvriyskiy/database_course/main/pkg"
	mocks_service "github.com/Mamvriyskiy/database_course/main/pkg/repository/mocks"
	"github.com/Mamvriyskiy/database_course/main/pkg/service"
	"github.com/Mamvriyskiy/database_course/main/tests/factory"
	method "github.com/Mamvriyskiy/database_course/main/tests/method"
	"github.com/golang/mock/gomock"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

func (s *MyUnitTestsSuite) TestUpdateHomeBL(t provider.T) {
	tests := []struct {
		nameTest   string
		updateHome pkg.UpdateNameHome
	}{
		{
			nameTest: "Test1",
			updateHome: pkg.UpdateNameHome{
				NewName: "home1",
			},
		},
		{
			nameTest: "Test2",
			updateHome: pkg.UpdateNameHome{
				NewName: "home2",
			},
		},
		{
			nameTest: "Test3",
			updateHome: pkg.UpdateNameHome{
				NewName: "home3",
			},
		},
	}

	// Создаем новый контроллер для управления mock-объектами
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создаем mock для репозитория
	mockRepo := mocks_service.NewMockIHomeRepo(ctrl)

	for _, test := range tests {
		t.Run(test.nameTest, func(t provider.T) {
			mockRepo.EXPECT().UpdateHome(test.updateHome).Return(nil)

			homeService := service.NewHomeService(mockRepo)

			err := homeService.UpdateHome(test.updateHome)

			t.Assert().NoError(err)
		})
	}
}

func (s *MyUnitTestsSuite) TestCreateHomeBL(t provider.T) {
	tests := []struct {
		nameTest string
		home     factory.ObjectSystem
		id       int
	}{
		{
			nameTest: "Test1",
			home:     factory.New("home", ""),
			id:       10,
		},
		{
			nameTest: "Test2",
			home:     factory.New("home", ""),
			id:       11,
		},
		{
			nameTest: "Test3",
			home:     factory.New("home", ""),
			id:       12,
		},
	}

	// Создаем новый контроллер для управления mock-объектами
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создаем mock для репозитория
	mockRepo := mocks_service.NewMockIHomeRepo(ctrl)

	for _, test := range tests {
		t.Run(test.nameTest, func(t provider.T) {
			newHome := test.home.(*method.TestHome)

			mockRepo.EXPECT().CreateHome(newHome.Home).Return(test.id, nil)

			homeService := service.NewHomeService(mockRepo)

			homeID, err := homeService.CreateHome(newHome.Home)

			t.Assert().NoError(err)
			t.Assert().Equal(test.id, homeID)
		})
	}
}

func (s *MyUnitTestsSuite) TestDeleteHomeBL(t provider.T) {
	tests := []struct {
		nameTest string
		home     factory.ObjectSystem
		userID   int
	}{
		{
			nameTest: "Test1",
			home:     factory.New("home", ""),
			userID:   10,
		},
		{
			nameTest: "Test2",
			home:     factory.New("home", ""),
			userID:   11,
		},
		{
			nameTest: "Test3",
			home:     factory.New("home", ""),
			userID:   12,
		},
	}

	// Создаем новый контроллер для управления mock-объектами
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создаем mock для репозитория
	mockRepo := mocks_service.NewMockIHomeRepo(ctrl)

	for _, test := range tests {
		t.Run(test.nameTest, func(t provider.T) {
			newHome := test.home.(*method.TestHome)

			mockRepo.EXPECT().DeleteHome(test.userID, newHome.Name).Return(nil)

			homeService := service.NewHomeService(mockRepo)

			err := homeService.DeleteHome(test.userID, newHome.Name)

			t.Assert().NoError(err)
		})
	}
}

func (s *MyUnitTestsSuite) TestGetListHomeBL(t provider.T) {
	tests := []struct {
		nameTest string
		lenList  int
		userID   int
	}{
		{
			nameTest: "Test1",
			lenList:  1,
			userID:   10,
		},
		{
			nameTest: "Test2",
			lenList:  5,
			userID:   11,
		},
		{
			nameTest: "Test3",
			lenList:  10,
			userID:   12,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks_service.NewMockIHomeRepo(ctrl)

	for _, test := range tests {
		t.Run(test.nameTest, func(t provider.T) {
			listHome := make([]pkg.Home, test.lenList)
			for i := 0; i < test.lenList; i++ {
				newHome := factory.New("home", "")
				home := newHome.(*method.TestHome)

				listHome[i] = home.Home
			}

			mockRepo.EXPECT().ListUserHome(test.userID).Return(listHome, nil)

			homeService := service.NewHomeService(mockRepo)

			resultList, err := homeService.ListUserHome(test.userID)

			t.Assert().NoError(err)
			t.Assert().Equal(listHome, resultList)
		})
	}
}
