package introtests

import (
	"fmt"
	//"github.com/jmoiron/sqlx"
	"github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/Mamvriyskiy/database_course/main/pkg/service"
	"github.com/Mamvriyskiy/database_course/main/pkg/repository"
	"github.com/Mamvriyskiy/database_course/main/tests/factory"
	method "github.com/Mamvriyskiy/database_course/main/tests/method"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"errors"
)

func (s *MyFirstSuite) TestCreateDeviceIntro(t provider.T) {
	tests := []struct {
		nameTest  string
		device    factory.ObjectSystem
		character factory.ObjectSystem
		home      factory.ObjectSystem
		access    factory.ObjectSystem
		user      factory.ObjectSystem
		devChar   pkg.DeviceCharacteristics
	}{
		{
			nameTest:  "Test1",
			device:    factory.New("device", ""),
			user:      factory.New("user", ""),
			home:      factory.New("home", ""),
			access:    factory.New("access", ""),
			character: factory.New("character", ""),
			devChar: pkg.DeviceCharacteristics{
				Values: 100,
			},
		},
		{
			nameTest:  "Test2",
			user:      factory.New("user", ""),
			home:      factory.New("home", ""),
			access:    factory.New("access", ""),
			device:    factory.New("device", ""),
			character: factory.New("character", ""),
			devChar: pkg.DeviceCharacteristics{
				Values: 200,
			},
		},
		{
			nameTest:  "Test3",
			device:    factory.New("device", ""),
			user:      factory.New("user", ""),
			access:    factory.New("access", ""),
			home:      factory.New("home", ""),
			character: factory.New("character", ""),
			devChar: pkg.DeviceCharacteristics{
				Values: 300,
			},
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

			//создание харакетристик устройства
			// newChar := test.character.(*method.TestCharacter)

			//создание устройства
			newDevice := test.device.(*method.TestDevice)
			newDevice.Home = newHome.Name
			deviceID, err := services.IDevice.CreateDevice(clientID, &newDevice.Devices)
			t.Require().NoError(err)

			query := `SELECT name FROM device WHERE deviceID = $1`
			row := connDB.QueryRow(query, deviceID)

			var nameDev string
			err = row.Scan(&nameDev)
			t.Require().NoError(err)

			t.Assert().Equal(newDevice.Name, nameDev)
		})
	}
}

func (s *MyFirstSuite) TestDeleteDeviceIntro(t provider.T) {
	tests := []struct {
		nameTest  string
		device    factory.ObjectSystem
		character factory.ObjectSystem
		home      factory.ObjectSystem
		access    factory.ObjectSystem
		user      factory.ObjectSystem
		devChar   pkg.DeviceCharacteristics
	}{
		{
			nameTest:  "Test1",
			device:    factory.New("device", ""),
			user:      factory.New("user", ""),
			home:      factory.New("home", ""),
			access:    factory.New("access", ""),
			character: factory.New("character", ""),
			devChar: pkg.DeviceCharacteristics{
				Values: 100,
			},
		},
		{
			nameTest:  "Test2",
			user:      factory.New("user", ""),
			home:      factory.New("home", ""),
			access:    factory.New("access", ""),
			device:    factory.New("device", ""),
			character: factory.New("character", ""),
			devChar: pkg.DeviceCharacteristics{
				Values: 200,
			},
		},
		{
			nameTest:  "Test3",
			device:    factory.New("device", ""),
			user:      factory.New("user", ""),
			access:    factory.New("access", ""),
			home:      factory.New("home", ""),
			character: factory.New("character", ""),
			devChar: pkg.DeviceCharacteristics{
				Values: 300,
			},
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
			newDevice.HomeID = homeID
			deviceID, err := newDevice.InsertObject(connDB)

			//создание харакетристик устройства
			newChar := test.character.(*method.TestCharacter)
			typeID, err := newChar.InsertObject(connDB)
			t.Require().NoError(err)
			
			query3 := fmt.Sprintf(`INSERT INTO devicecharacteristics (deviceID, valueschar, typecharacterid)
				values ($1, $2, $3)`)
			_ = connDB.QueryRow(query3, deviceID, test.devChar.Values, typeID)

			err = services.IDevice.DeleteDevice(clientID, newDevice.Name, newHome.Name) 
			t.Require().NoError(err)

			query := `SELECT name FROM device WHERE deviceID = $1`
			row := connDB.QueryRow(query, deviceID)

			var resultName string
			err = row.Scan(&resultName)

			t.Assert().Equal(err, errors.New("sql: no rows in result set"))
		})
	}
}
