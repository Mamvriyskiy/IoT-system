package repository

import (
	"fmt"

	"github.com/Mamvriyskiy/DBCourse/main/logger"
	pkg "github.com/Mamvriyskiy/DBCourse/main/pkg"
	"github.com/jmoiron/sqlx"
)

type DevicePostgres struct {
	db *sqlx.DB
}

func NewDevicePostgres(db *sqlx.DB) *DevicePostgres {
	return &DevicePostgres{db: db}
}

func (r *DevicePostgres) CreateDevice(userID int, device *pkg.Devices) (int, error) {
	var homeID int
	const queryHomeID = `select h.homeid from home h 
	where h.homeid in (select a.homeid from accesshome a 
		where a.accessid in (select a.accessid from accessclient a 
			JOIN access ac ON a.accessid = ac.accessid where clientid = $1 AND accessLevel = 4));`

	err := r.db.Get(&homeID, queryHomeID, userID)
	if err != nil {
		logger.Log("Error", "Get", "Error get homeID:", err, &homeID, userID)
		return 0, err
	}

	var id int
	query := fmt.Sprintf(`INSERT INTO %s (name, TypeDevice, Status, 
		Brand, PowerConsumption, MinParametr, MaxParametr) 
			values ($1, $2, $3, $4, $5, $6, $7) RETURNING deviceID`, "device")
	row := r.db.QueryRow(query, device.Name, device.TypeDevice,
		device.Status, device.Brand, device.PowerConsumption,
		device.MinParameter, device.MaxParameter)

	err = row.Scan(&id)
	if err != nil {
		logger.Log("Error", "Scan", "Error insert into device:", err, &id)
		return 0, err
	}

	query1 := fmt.Sprintf("INSERT INTO %s (homeID, deviceId) VALUES ($1, $2)", "deviceHome")
	_ = r.db.QueryRow(query1, homeID, id)

	return id, nil
}

func (r *DevicePostgres) DeleteDevice(userID int, name string) error {
	var homeID int
	const queryHomeID = `select h.homeid from home h 
	where h.homeid in (select a.homeid from accesshome a 
		where a.accessid in (select a.accessid from accessclient a 
			JOIN access ac ON a.accessid = ac.accessid where clientid = $1 AND accessLevel = 4));`

	err := r.db.Get(&homeID, queryHomeID, userID)
	if err != nil {
		logger.Log("Error", "Get", "Error get homeID:", err, &homeID, userID)
		return err
	}
	var deviceID int
	queryDeviceID := `select d.deviceid from device d 
	join devicehome d2 on d.deviceid = d2.deviceid 
		where d2.homeid = $1 and d.name = $2;`
	err = r.db.Get(&deviceID, queryDeviceID, homeID, name)
	if err != nil {
		logger.Log("Error", "Get", "Error get deviceID:", err, &deviceID, homeID, name)
		return err
	}

	query := `DELETE FROM historydev
			WHERE historydevid IN 
				(SELECT h2.historydevid FROM historydevice h2 
					WHERE h2.deviceid = $1);`

	_, err = r.db.Exec(query, deviceID)
	if err != nil {
		logger.Log("Error", "Exec", "Error delete historydev:", err, deviceID)
		return err
	}

	query = `DELETE FROM device 
				where deviceid = $1;`
	_, err = r.db.Exec(query, deviceID)
	if err != nil {
		logger.Log("Error", "Exec", "Error delete device:", err, deviceID)
		return err
	}

	return err
}

func (r *DevicePostgres) GetDeviceByID(deviceID int) (pkg.Devices, error) {
	var device pkg.Devices
	query := fmt.Sprintf("SELECT * from %s where deviceid = $1", "device")
	err := r.db.Get(&device, query, deviceID)

	return device, err
}
