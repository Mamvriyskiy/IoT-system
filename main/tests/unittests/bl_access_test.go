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

func (s *MyUnitTestsSuite) TestAddClientBL(t provider.T) {
	tests := []struct {
		nameTest   string
		accessUser factory.ObjectSystem
		userID     string
		accessID   string
	}{
		{
			nameTest:   "Test1",
			accessUser: factory.New("access", ""),
			userID:     "10",
			accessID:   "1",
		},
		{
			nameTest:   "Test2",
			accessUser: factory.New("access", ""),
			userID:     "11",
			accessID:   "2",
		},
		{
			nameTest:   "Test3",
			accessUser: factory.New("access", ""),
			userID:     "12",
			accessID:   "3",
		},
	}

	// Создаем новый контроллер для управления mock-объектами
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создаем mock для репозитория
	mockRepo := mocks_service.NewMockIAccessHomeRepo(ctrl)

	for _, test := range tests {
		t.Run(test.nameTest, func(t provider.T) {
			newAccessUser := test.accessUser.(*method.TestAccess)

			serviceAccess := pkg.AccessService{
				Access: newAccessUser.Access,
			}

			mockRepo.EXPECT().AddUser(test.userID, serviceAccess).Return(test.accessID, nil)

			accessService := service.NewAccessHomeService(mockRepo)

			handlerAccess := pkg.AccessHandler{
				Access: newAccessUser.Access,
			}

			accessID, err := accessService.AddUser(test.userID, handlerAccess)

			t.Assert().NoError(err)
			t.Assert().Equal(test.accessID, accessID)
		})
	}
}

func (s *MyUnitTestsSuite) TestUpdateLevelBL(t provider.T) {
	tests := []struct {
		nameTest   string
		accessUser factory.ObjectSystem
		userID     string
	}{
		{
			nameTest:   "Test1",
			accessUser: factory.New("access", ""),
			userID:     "10",
		},
		{
			nameTest:   "Test2",
			accessUser: factory.New("access", ""),
			userID:     "11",
		},
		{
			nameTest:   "Test3",
			accessUser: factory.New("access", ""),
			userID:     "12",
		},
	}

	// Создаем новый контроллер для управления mock-объектами
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создаем mock для репозитория
	mockRepo := mocks_service.NewMockIAccessHomeRepo(ctrl)

	for _, test := range tests {
		t.Run(test.nameTest, func(t provider.T) {
			newAccessUser := test.accessUser.(*method.TestAccess)

			serviceAccess := pkg.AccessService{
				Access: newAccessUser.Access,
			}

			mockRepo.EXPECT().UpdateLevel(test.userID, serviceAccess).Return(nil)

			accessService := service.NewAccessHomeService(mockRepo)

			handlerAccess := pkg.AccessHandler{
				Access: newAccessUser.Access,
			}

			err := accessService.UpdateLevel(test.userID, handlerAccess)

			t.Assert().NoError(err)
		})
	}
}

func (s *MyUnitTestsSuite) TestUpdateStatusBL(t provider.T) {
	tests := []struct {
		nameTest   string
		accessHome pkg.AccessService
		userID     string
	}{
		{
			nameTest:   "Test1",
			accessHome: pkg.AccessService{
				Access: pkg.Access{
					AccessStatus: "blocked",
				},
			},
			userID:     "10",
		},
		{
			nameTest:   "Test2",
			accessHome: pkg.AccessService{
				Access: pkg.Access{
					AccessStatus: "blocked",
				},
			},
			userID:     "11",
		},
		{
			nameTest:   "Test3",
			accessHome: pkg.AccessService{
				Access: pkg.Access{
					AccessStatus: "blocked",
				},
			},
			userID:     "12",
		},
	}

	// Создаем новый контроллер для управления mock-объектами
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создаем mock для репозитория
	mockRepo := mocks_service.NewMockIAccessHomeRepo(ctrl)

	for _, test := range tests {
		t.Run(test.nameTest, func(t provider.T) {
			mockRepo.EXPECT().UpdateStatus(test.userID, test.accessHome).Return(nil)

			accessService := service.NewAccessHomeService(mockRepo)

			handlerAccess := pkg.AccessHandler{
				Access: test.accessHome.Access,
			}

			err := accessService.UpdateStatus(test.userID, handlerAccess)

			t.Assert().NoError(err)
		})
	}
}

func (s *MyUnitTestsSuite) TestGetListUserHomeBL(t provider.T) {
	tests := []struct {
		nameTest string
		lenList  int
		homeID   string
	}{
		{
			nameTest: "Test1",
			lenList:  1,
			homeID:   "10",
		},
		{
			nameTest: "Test2",
			lenList:  5,
			homeID:   "11",
		},
		{
			nameTest: "Test3",
			lenList:  10,
			homeID:   "12",
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks_service.NewMockIAccessHomeRepo(ctrl)

	for _, test := range tests {
		t.Run(test.nameTest, func(t provider.T) {
			listUser := make([]pkg.AccessInfoData, test.lenList)
			for i := 0; i < test.lenList; i++ {
				newUser := factory.New("user", "")
				user := newUser.(*method.TestUser)

				newHome := factory.New("home", "")
				home := newHome.(*method.TestHome)

				newAccess := factory.New("access", "")
				access := newAccess.(*method.TestAccess)

				listUser[i].Home = home.Name
				listUser[i].Login = user.Username
				listUser[i].Email = user.Email
				listUser[i].AccessLevel = access.AccessLevel
				listUser[i].AccessStatus = "active"
			}

			mockRepo.EXPECT().GetListUserHome(test.homeID).Return(listUser, nil)

			accessService := service.NewAccessHomeService(mockRepo)

			resultList, err := accessService.GetListUserHome(test.homeID)

			t.Assert().NoError(err)
			t.Assert().Equal(listUser, resultList)
		})
	}
}
