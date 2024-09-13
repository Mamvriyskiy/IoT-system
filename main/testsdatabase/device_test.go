package testsdatabase

import (
	"fmt"
	"testing"
	//"github.com/jmoiron/sqlx"
	"github.com/Mamvriyskiy/database_course/main/pkg"
	//"github.com/Mamvriyskiy/database_course/main/pkg/repository"
	"github.com/stretchr/testify/assert"
)

func TestUpdateStatusFunc(t *testing.T) {
	insertData := []struct {
		device pkg.Devices
		ID     int
	}{
		{
			device: pkg.Devices{
				Name:       "dev1",
				TypeDevice: "type1",
				Status:     "inactive",
				Brand:      "brand1",
			},
			ID: 1,
		},
		{
			device: pkg.Devices{
				Name:       "dev2",
				TypeDevice: "type2",
				Status:     "inactive",
				Brand:      "brand2",
			},
			ID: 2,
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
			devID:      1,
		},
		{
			nameTest:   "Test2",
			nameDev:    "dev2",
			status:     "inactive",
			resultCode: -3,
			devID:      2,
		},
		{
			nameTest:   "Test3",
			nameDev:    "dev2",
			status:     "active",
			resultCode: 0,
			devID:      2,
		},
		{
			nameTest:   "Test4",
			nameDev:    "dev1",
			status:     "active",
			resultCode: 0,
			devID:      1,
		},
		{
			nameTest:   "Test5",
			nameDev:    "dev2",
			status:     "active",
			resultCode: -2,
			devID:      2,
		},
		{
			nameTest:   "Test6",
			nameDev:    "dev2",
			status:     "inactive",
			resultCode: 0,
			devID:      2,
		},
		{
			nameTest:   "Test7",
			nameDev:    "dev1",
			status:     "inactive",
			resultCode: 0,
			devID:      1,
		},
		{
			nameTest:   "Test8",
			nameDev:    "dev4",
			status:     "inactive",
			resultCode: -1,
			devID:      4,
		},
	}

	//Заполнение таблицы Device данными

	for _, data := range insertData {
		query := fmt.Sprintf(`INSERT INTO %s (name, TypeDevice, Status, Brand)
			values ($1, $2, $3, $4) RETURNING deviceID`, "device")
		row := connDB.QueryRow(query, data.device.Name, data.device.TypeDevice,
			data.device.Status, data.device.Brand)

		var id int
		err := row.Scan(&id)
		if err != nil {
			assert.NoError(t, err, "")
		}
	}

	for _, test := range tests {
		t.Run(test.nameTest, func(t *testing.T) {

			var result int
			queryUpdateStatus := fmt.Sprintf(`select update_status($1, $2);`)
			err := connDB.Get(&result, queryUpdateStatus, test.devID, test.status)

			if err != nil {
				assert.NoError(t, err, "")
			}

			assert.Equal(t, test.resultCode, result, "The result status code should be the same.")
		})
	}
}
