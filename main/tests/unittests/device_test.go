package unittests

import (
	"fmt"
	//"github.com/jmoiron/sqlx"
	"github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/Mamvriyskiy/database_course/main/pkg/repository"
	"github.com/Mamvriyskiy/database_course/main/tests/factory"
	method "github.com/Mamvriyskiy/database_course/main/tests/method"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"errors"
)

func (s *MyUnitTestsSuite) TestUpdateStatusFunc(t provider.T) {
	insertData := []struct {
		device factory.ObjectSystem
		ID     int
	}{
		{
			device: factory.New("device", ""),
			ID:     0,
		},
		{
			device: factory.New("device", ""),
			ID:     1,
		},
	}

	tests := []struct {
		nameTest   string
		nameDev    string
		status     string
		resultCode int
		devID      int
	}{
		{
			nameTest:   "Test1",
			nameDev:    "dev1",
			status:     "inactive",
			resultCode: -3,
			devID:     0,
		},
		{
			nameTest:   "Test2",
			nameDev:    "dev2",
			status:     "inactive",
			resultCode: -3,
			devID:     1,
		},
		{
			nameTest:   "Test3",
			nameDev:    "dev2",
			status:     "active",
			resultCode: 0,
			devID:      1,
		},
		{
			nameTest:   "Test4",
			nameDev:    "dev1",
			status:     "active",
			resultCode: 0,
			devID:      0,
		},
		{
			nameTest:   "Test5",
			nameDev:    "dev2",
			status:     "active",
			resultCode: -2,
			devID:      1,
		},
		{
			nameTest:   "Test6",
			nameDev:    "dev2",
			status:     "inactive",
			resultCode: 0,
			devID:      1,
		},
		{
			nameTest:   "Test7",
			nameDev:    "dev1",
			status:     "inactive",
			resultCode: 0,
			devID:      0,
		},
		{
			nameTest:   "Test8",
			nameDev:    "dev4",
			status:     "inactive",
			resultCode: -1,
			devID:      3,
		},
	}

	devicesID := make([]string, 2)
	for i, data := range insertData {
		id, err := data.device.InsertObject(connDB)
		t.Require().NoError(err)
		devicesID[i] = id
	}

	for _, test := range tests {
		t.Run(test.nameTest, func(t provider.T) {

			var result int
			queryUpdateStatus := fmt.Sprintf(`select update_status($1, $2);`)
			id := "123e4567-e89b-12d3-a456-426614174000"
			if test.devID < 2 {
				id = devicesID[test.devID]
			}

			err := connDB.Get(&result, queryUpdateStatus, id, test.status)
			t.Require().NoError(err)

			t.Assert().Equal(test.resultCode, result)
		})
	}
}

func (s *MyUnitTestsSuite) TestCreateDevice(t provider.T) {
	tests := []struct {
		nameTest  string
		device    factory.ObjectSystem
		character factory.ObjectSystem
		home      factory.ObjectSystem
		devChar   pkg.DeviceCharacteristicsService
	}{
		{
			nameTest:  "Test1",
			device:    factory.New("device", ""),
			home:      factory.New("home", ""),
			character: factory.New("character", ""),
			devChar: pkg.DeviceCharacteristicsService{
				Values: 100,
			},
		},
		{
			nameTest:  "Test2",
			home:      factory.New("home", ""),
			device:    factory.New("device", ""),
			character: factory.New("character", ""),
			devChar: pkg.DeviceCharacteristicsService{
				Values: 200,
			},
		},
		{
			nameTest:  "Test3",
			device:    factory.New("device", ""),
			home:      factory.New("home", ""),
			character: factory.New("character", ""),
			devChar: pkg.DeviceCharacteristicsService{
				Values: 300,
			},
		},
	}

	repos := repository.NewRepository(connDB)

	for _, test := range tests {
		t.Run(test.nameTest, func(t provider.T) {
			//создание дома
			newHome := test.home.(*method.TestHome)

			homeID, err := newHome.InsertObject(connDB)
			t.Require().NoError(err)

			//создание харакетристик устройства
			newChar := test.character.(*method.TestCharacter)

			//создание устройства
			newDevice := test.device.(*method.TestDevice)
			// newDevice.Home = newHome.Name

			deviceService := pkg.DevicesService{
				Devices: newDevice.Devices,
				DevicesInfo: newDevice.DevicesInfo,
			}

			deviceID, err := repos.IDeviceRepo.CreateDevice(homeID, deviceService, test.devChar, newChar.TypeCharacterService)
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

func (s *MyUnitTestsSuite) TestDeleteDevice(t provider.T) {
	tests := []struct {
		nameTest  string
		device    factory.ObjectSystem
		character factory.ObjectSystem
		home      factory.ObjectSystem
		devChar   pkg.DeviceCharacteristicsService
	}{
		{
			nameTest:  "Test1",
			device:    factory.New("device", ""),
			home:      factory.New("home", ""),
			character: factory.New("character", ""),
			devChar: pkg.DeviceCharacteristicsService{
				Values: 100,
			},
		},
		{
			nameTest:  "Test2",
			home:      factory.New("home", ""),
			device:    factory.New("device", ""),
			character: factory.New("character", ""),
			devChar: pkg.DeviceCharacteristicsService{
				Values: 200,
			},
		},
		{
			nameTest:  "Test3",
			device:    factory.New("device", ""),
			home:      factory.New("home", ""),
			character: factory.New("character", ""),
			devChar: pkg.DeviceCharacteristicsService{
				Values: 300,
			},
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
			newDevice.HomeID = homeID
			deviceID, err := newDevice.InsertObject(connDB)

			//создание харакетристик устройства
			newChar := test.character.(*method.TestCharacter)
			typeID, err := newChar.InsertObject(connDB)
			t.Require().NoError(err)
			
			query3 := fmt.Sprintf(`INSERT INTO DeviceCharacteristicsService (deviceID, valueschar, typecharacterid)
				values ($1, $2, $3)`)
			_ = connDB.QueryRow(query3, deviceID, test.devChar.Values, typeID)

			err = repos.IDeviceRepo.DeleteDevice(deviceID) 
			t.Require().NoError(err)

			query := `SELECT name FROM device WHERE deviceID = $1`
			row := connDB.QueryRow(query, deviceID)

			var resultName string
			err = row.Scan(&resultName)

			t.Assert().Equal(err, errors.New("sql: no rows in result set"))
		})
	}
}
