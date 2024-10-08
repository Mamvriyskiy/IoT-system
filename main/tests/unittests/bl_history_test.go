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

func (s *MyUnitTestsSuite) TestCreateDeviceHistoryBL(t provider.T) {
	tests := []struct {
		nameTest  string
		history   factory.ObjectSystem
		deviceID  string
		historyID string
	}{
		{
			nameTest:  "Test1",
			deviceID:  "1",
			historyID: "4",
		},
		{
			nameTest:  "Test2",
			deviceID: "2",
			historyID: "5",
		},
		{
			nameTest:  "Test3",
			deviceID:  "3",
			historyID: "6",
		},
	}

	// Создаем новый контроллер для управления mock-объектами
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создаем mock для репозитория
	mockRepo := mocks_service.NewMockIHistoryDeviceRepo(ctrl)

	for _, test := range tests {
		t.Run(test.nameTest, func(t provider.T) {
			mockRepo.EXPECT().CreateDeviceHistory(test.deviceID, gomock.Any()).Return(test.historyID, nil)

			homeService := service.NewHistoryDeviceService(mockRepo)

			historyID, err := homeService.CreateDeviceHistory(test.deviceID)

			t.Assert().NoError(err)
			t.Assert().Equal(test.historyID, historyID)
		})
	}
}

func (s *MyUnitTestsSuite) TestGetDeviceHistoryBL(t provider.T) {
	tests := []struct {
		nameTest   string
		lenList    int
		deviceID     string
	}{
		{
			nameTest:   "Test1",
			lenList:    1,
			deviceID:     "10",
		},
		{
			nameTest:   "Test2",
			lenList:    5,
			deviceID:     "11",
		},
		{
			nameTest:   "Test3",
			lenList:    10,
			deviceID:     "12",
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks_service.NewMockIHistoryDeviceRepo(ctrl)

	for _, test := range tests {
		t.Run(test.nameTest, func(t provider.T) {
			listHistory := make([]pkg.DevicesHistory, test.lenList)
			for i := 0; i < test.lenList; i++ {
				newHistory := factory.New("history", "")
				history := newHistory.(*method.TestHistory)
	
				curHistory := pkg.DevicesHistory{
					TimeWork: history.TimeWork,
					AverageIndicator: history.AverageIndicator,
					EnergyConsumed: history.EnergyConsumed,
				}

				listHistory[i] = curHistory
			}

			mockRepo.EXPECT().GetDeviceHistory(test.deviceID).Return(listHistory, nil)

			homeService := service.NewHistoryDeviceService(mockRepo)

			resultList, err := homeService.GetDeviceHistory(test.deviceID)

			t.Assert().NoError(err)
			t.Assert().Equal(listHistory, resultList)
		})
	}
}
