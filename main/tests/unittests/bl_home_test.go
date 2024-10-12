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
		nameTest string
		home     string
		homeID   string
	}{
		{
			nameTest: "Test1",
			home: "home1",
			homeID: "1",
		},
		{
			nameTest: "Test2",
			home: "home2",
			homeID: "2",
		},
		{
			nameTest: "Test3",
			home: "home3",
			homeID: "3",
		},
	}

	// Создаем новый контроллер для управления mock-объектами
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создаем mock для репозитория
	mockRepo := mocks_service.NewMockIHomeRepo(ctrl)

	for _, test := range tests {
		t.Run(test.nameTest, func(t provider.T) {
			mockRepo.EXPECT().UpdateHome(test.homeID, test.home).Return(nil)

			homeService := service.NewHomeService(mockRepo)

			err := homeService.UpdateHome(test.homeID, test.home)

			t.Assert().NoError(err)
		})
	}
}

func (s *MyUnitTestsSuite) TestCreateHomeBL(t provider.T) {
	tests := []struct {
		nameTest string
		home     factory.ObjectSystem
		id       string
	}{
		{
			nameTest: "Test1",
			home:     factory.New("home", ""),
			id:       "10",
		},
		{
			nameTest: "Test2",
			home:     factory.New("home", ""),
			id:       "11",
		},
		{
			nameTest: "Test3",
			home:     factory.New("home", ""),
			id:       "12",
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

			homeService := pkg.HomeService{
				Home: newHome.Home,
			}

			mockRepo.EXPECT().CreateHome(homeService).Return(test.id, nil)
			mockRepo.EXPECT().GetHomeByID(test.id).Return(newHome.HomeData, nil)

			homeSRV := service.NewHomeService(mockRepo)

			homeHandler := pkg.HomeHandler{
				Home: newHome.Home,
			}

			insertHome, err := homeSRV.CreateHome(homeHandler)

			t.Assert().NoError(err)
			t.Assert().Equal(newHome.HomeData, insertHome)
		})
	}
}

func (s *MyUnitTestsSuite) TestDeleteHomeBL(t provider.T) {
	tests := []struct {
		nameTest string
		deviceID   string
	}{
		{
			nameTest: "Test1",
			deviceID:   "10",
		},
		{
			nameTest: "Test2",
			deviceID:   "11",
		},
		{
			nameTest: "Test3",
			deviceID:   "12",
		},
	}

	// Создаем новый контроллер для управления mock-объектами
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создаем mock для репозитория
	mockRepo := mocks_service.NewMockIHomeRepo(ctrl)

	for _, test := range tests {
		t.Run(test.nameTest, func(t provider.T) {
			mockRepo.EXPECT().DeleteHome(test.deviceID).Return(nil)

			homeService := service.NewHomeService(mockRepo)

			err := homeService.DeleteHome(test.deviceID)

			t.Assert().NoError(err)
		})
	}
}

func (s *MyUnitTestsSuite) TestGetListHomeBL(t provider.T) {
	tests := []struct {
		nameTest string
		lenList  int
		userID   string
	}{
		{
			nameTest: "Test1",
			lenList:  1,
			userID:   "10",
		},
		{
			nameTest: "Test2",
			lenList:  5,
			userID:   "11",
		},
		{
			nameTest: "Test3",
			lenList:  10,
			userID:   "12",
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks_service.NewMockIHomeRepo(ctrl)

	for _, test := range tests {
		t.Run(test.nameTest, func(t provider.T) {
			listHome := make([]pkg.HomeData, test.lenList)
			for i := 0; i < test.lenList; i++ {
				newHome := factory.New("home", "")
				home := newHome.(*method.TestHome)

				listHome[i] = home.HomeData
			}

			mockRepo.EXPECT().ListUserHome(test.userID).Return(listHome, nil)

			homeService := service.NewHomeService(mockRepo)

			resultList, err := homeService.ListUserHome(test.userID)

			t.Assert().NoError(err)
			t.Assert().Equal(listHome, resultList)
		})
	}
}
