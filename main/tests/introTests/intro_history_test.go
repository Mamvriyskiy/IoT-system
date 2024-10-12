package introtests

import (
	//"github.com/jmoiron/sqlx"
	"github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/Mamvriyskiy/database_course/main/pkg/service"
	"github.com/Mamvriyskiy/database_course/main/pkg/repository"
	"github.com/Mamvriyskiy/database_course/main/tests/factory"
	method "github.com/Mamvriyskiy/database_course/main/tests/method"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

func (s *MyIntroTestsSuite) TestCreateHistoryIntro(t provider.T) {
	tests := []struct {
		nameTest  string
		device    factory.ObjectSystem
		home      factory.ObjectSystem
		history   factory.ObjectSystem
	}{
		{
			nameTest:  "Test1",
			device:    factory.New("device", ""),
			home:      factory.New("home", ""),
			history:   factory.New("history", ""),
		},
		{
			nameTest:  "Test2",
			home:      factory.New("home", ""),
			device:    factory.New("device", ""),
			history:   factory.New("history", ""),
		},
		{
			nameTest:  "Test3",
			device:    factory.New("device", ""),
			home:      factory.New("home", ""),
			history:   factory.New("history", ""),
		},
	}

	repos := repository.NewRepository(connDB)
	services := service.NewServicesPsql(repos)

	for _, test := range tests {
		t.Run(test.nameTest, func(t provider.T) {
			//создание дома
			newHome := test.home.(*method.TestHome)

			homeID, err := newHome.InsertObject(connDB)
			t.Require().NoError(err)

			//создание устройства
			newDevice := test.device.(*method.TestDevice)
			newDevice.Home = newHome.Name
			newDevice.HomeID = homeID
			deviceID, err := newDevice.InsertObject(connDB)
			t.Require().NoError(err)

			idHistory, err := services.IHistoryDevice.CreateDeviceHistory(deviceID)
			t.Require().NoError(err)

			var dbHistoryID string
			query := `SELECT historydevID FROM historyDev WHERE historydevID = $1`
			row := connDB.QueryRow(query, idHistory)
	
			err = row.Scan(&dbHistoryID)
			t.Require().NoError(err)

			t.Assert().Equal(idHistory, dbHistoryID)
		})
	}
}

func (s *MyIntroTestsSuite) TestGetDeviceHistoryIntro(t provider.T) {
	tests := []struct {
		nameTest string
		lenList  int
		user     factory.ObjectSystem
		device    factory.ObjectSystem
		home      factory.ObjectSystem
		access    factory.ObjectSystem
	}{
		{
			nameTest: "Test1",
			lenList:  1,
			device:    factory.New("device", ""),
			user:      factory.New("user", ""),
			access:    factory.New("access", ""),
			home:      factory.New("home", ""),
		},
		{
			nameTest: "Test2",
			lenList:  10,
			device:    factory.New("device", ""),
			user:      factory.New("user", ""),
			access:    factory.New("access", ""),
			home:      factory.New("home", ""),
		},
	}

	repos := repository.NewRepository(connDB)
	services := service.NewServicesPsql(repos)

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
			newDevice.Home = newHome.Name
			newDevice.HomeID = homeID
			deviceID, err := newDevice.InsertObject(connDB)
			t.Require().NoError(err)

			listHistory := make([]pkg.DevicesHistory, test.lenList)
			for i := 0; i < test.lenList; i++ {
				history := factory.New("history", "")
				newHistory := history.(*method.TestHistory)

				newHistory.Home = newHome.Name
				newHistory.Name = newDevice.Name
				newHistory.DeviceID = deviceID
				historyID, err := newHistory.InsertObject(connDB)
				t.Require().NoError(err)

				currentHistory := pkg.DevicesHistory{
					ID: historyID,
					TimeWork: newHistory.TimeWork,
					AverageIndicator: newHistory.AverageIndicator,
					EnergyConsumed: newHistory.EnergyConsumed,
				}

				listHistory[i] = currentHistory
			}

			resultListHistory, err := services.IHistoryDevice.GetDeviceHistory(deviceID)

			t.Require().NoError(err)

			t.Assert().Equal(listHistory, resultListHistory)
		})
	}
}
