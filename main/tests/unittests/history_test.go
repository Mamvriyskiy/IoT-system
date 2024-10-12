package unittests

import (
	//"github.com/jmoiron/sqlx"
	"github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/Mamvriyskiy/database_course/main/pkg/repository"
	"github.com/Mamvriyskiy/database_course/main/tests/factory"
	method "github.com/Mamvriyskiy/database_course/main/tests/method"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

func (s *MyUnitTestsSuite) TestCreateHistory(t provider.T) {
	tests := []struct {
		nameTest  string
		device    factory.ObjectSystem
		home      factory.ObjectSystem
		access    factory.ObjectSystem
		user      factory.ObjectSystem
		history   factory.ObjectSystem
	}{
		{
			nameTest:  "Test1",
			device:    factory.New("device", ""),
			user:      factory.New("user", ""),
			home:      factory.New("home", ""),
			access:    factory.New("access", ""),
			history:   factory.New("history", ""),
		},
		{
			nameTest:  "Test2",
			user:      factory.New("user", ""),
			home:      factory.New("home", ""),
			access:    factory.New("access", ""),
			device:    factory.New("device", ""),
			history:   factory.New("history", ""),
		},
		{
			nameTest:  "Test3",
			device:    factory.New("device", ""),
			user:      factory.New("user", ""),
			access:    factory.New("access", ""),
			home:      factory.New("home", ""),
			history:   factory.New("history", ""),
		},
	}

	repos := repository.NewRepository(connDB)

	for _, test := range tests {
		t.Run(test.nameTest, func(t provider.T) {
			//Создание пользователя
			newUser := test.user.(*method.TestUser)

			clientID, err := newUser.InsertObject(connDB)
			t.Require().NoError(err)

			//создание дома
			newHome := test.home.(*method.TestHome)

			homeID, err := newHome.InsertObject(connDB)
			t.Require().NoError(err)

			//создание доступа к дому
			newAccess := test.access.(*method.TestAccess)
			newAccess.ClientID = clientID
			newAccess.HomeID = homeID
			_, err = newAccess.InsertObject(connDB)
			t.Require().NoError(err)

			//создание устройства
			newDevice := test.device.(*method.TestDevice)
			// newDevice.Home = newHome.Name
			newDevice.HomeID = homeID
			_, err = newDevice.InsertObject(connDB)
			t.Require().NoError(err)

			newHistory := test.history.(*method.TestHistory)

			// newHistory.Home = newHome.Name
			// newHistory.Name = newDevice.Name
			historyID, err := repos.IHistoryDeviceRepo.CreateDeviceHistory(clientID, newHistory.HistoryService)
			t.Require().NoError(err)

			testHistory := pkg.DevicesHistoryData{
				ID: historyID,
				History: pkg.History{
					TimeWork: newHistory.TimeWork,
					AverageIndicator: newHistory.AverageIndicator,
					EnergyConsumed: newHistory.EnergyConsumed,
				},
			}

			query := `SELECT * FROM historyDev WHERE historydevID = $1`
			row := connDB.QueryRow(query, historyID)

			var resultHistory pkg.DevicesHistoryData
			err = row.Scan(&resultHistory.ID, &resultHistory.TimeWork, &resultHistory.AverageIndicator, &resultHistory.EnergyConsumed)
			t.Require().NoError(err)

			t.Assert().Equal(testHistory, resultHistory)
		})
	}
}

func (s *MyUnitTestsSuite) TestGetDeviceHistory(t provider.T) {
	tests := []struct {
		nameTest string
		lenList  int
		device    factory.ObjectSystem
		home      factory.ObjectSystem
	}{
		{
			nameTest: "Test1",
			lenList:  1,
			device:    factory.New("device", ""),
			home:      factory.New("home", ""),
		},
		{
			nameTest: "Test2",
			lenList:  10,
			device:    factory.New("device", ""),
			home:      factory.New("home", ""),
		},
	}

	repos := repository.NewRepository(connDB)

	for _, test := range tests {
		t.Run(test.nameTest, func(t provider.T) {
			//создание дома
			newHome := test.home.(*method.TestHome)

			homeID, err := newHome.InsertObject(connDB)
			t.Require().NoError(err)

			//создание устройства
			newDevice := test.device.(*method.TestDevice)
			// newDevice.Home = newHome.Name
			newDevice.HomeID = homeID
			deviceID, err := newDevice.InsertObject(connDB)
			t.Require().NoError(err)

			listHistory := make([]pkg.DevicesHistoryData, test.lenList)
			for i := 0; i < test.lenList; i++ {
				history := factory.New("history", "")
				newHistory := history.(*method.TestHistory)

				// newHistory.Home = newHome.Name
				// newHistory.Name = newDevice.Name
				newHistory.DeviceID = deviceID
				historyID, err := newHistory.InsertObject(connDB)
				t.Require().NoError(err)

				currentHistory := pkg.DevicesHistoryData{
					ID: historyID,
					History: pkg.History{
						TimeWork: newHistory.TimeWork,
						AverageIndicator: newHistory.AverageIndicator,
						EnergyConsumed: newHistory.EnergyConsumed,
					},
				}

				listHistory[i] = currentHistory
			}

			resultListHistory, err := repos.IHistoryDeviceRepo.GetDeviceHistory(deviceID)

			t.Require().NoError(err)

			t.Assert().Equal(listHistory, resultListHistory)
		})
	}
}
