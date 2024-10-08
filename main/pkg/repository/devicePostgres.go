package repository

import (
	"fmt"

	"github.com/Mamvriyskiy/database_course/main/logger"
	pkg "github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/jmoiron/sqlx"
	"github.com/google/uuid"
)

type DevicePostgres struct {
	db *sqlx.DB
}

func NewDevicePostgres(db *sqlx.DB) *DevicePostgres {
	return &DevicePostgres{db: db}
}

func (r *DevicePostgres) GetListDevices(homeID string) ([]pkg.DevicesInfo, error) {
	const queryHomeID = `select d.name, d.status, d.Brand, d.deviceid from device d 
	where d.homeid = $1;`

	var deviceList []pkg.DevicesInfo
	err := r.db.Select(&deviceList, queryHomeID, homeID)
	if err != nil {
		logger.Log("Error", "Select", "Error get list devices:", err, homeID)
		return []pkg.DevicesInfo{}, err
	}

	return deviceList, nil 
}

func (r *DevicePostgres) CreateDevice(homeID string, device pkg.Devices, 
		character pkg.DeviceCharacteristics, typeCharacter pkg.TypeCharacter) (string, error) {
	var id string
	deviceID := uuid.New()
	query := fmt.Sprintf(`INSERT INTO %s (name, TypeDevice, Status, 
		Brand, homeid, deviceID)
			values ($1, $2, $3, $4, $5, $6) RETURNING deviceID`, "device")
	row := r.db.QueryRow(query, device.Name, device.TypeDevice,
		"inactive", device.Brand, homeID, deviceID)

	err := row.Scan(&id)
	if err != nil {
		logger.Log("Error", "Scan", "Error insert into device:", err, &id)
		return "", err
	}
	
	typecharacterID := uuid.New()
	var characterID string
	query2 := fmt.Sprintf(`INSERT INTO typecharacter (typecharacter, unitmeasure, typecharacterID)
		values ($1, $2, $3) RETURNING typecharacterID`)
	row = r.db.QueryRow(query2, typeCharacter.Type, typeCharacter.UnitMeasure, typecharacterID)
	
	err = row.Scan(&characterID)
	if err != nil {
		logger.Log("Error", "Scan", "Error insert into typecharacter:", err, &id)
		return "", err
	}

	query3 := fmt.Sprintf(`INSERT INTO devicecharacteristics (deviceID, valueschar, typecharacterid)
		values ($1, $2, $3)`)
	row = r.db.QueryRow(query3, id, character.Values, characterID)

	return id, nil
}

func (r *DevicePostgres) DeleteDevice(deviceID string) error {
	query := `DELETE FROM historydev
			WHERE historydevid IN 
				(SELECT h2.historydevid FROM historydevice h2 
					WHERE h2.deviceid = $1);`

	_, err := r.db.Exec(query, deviceID)
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

func (r *DevicePostgres) GetDeviceByID(deviceID string) (pkg.Devices, error) {
	var device pkg.Devices
	query := fmt.Sprintf("SELECT * from %s where deviceid = $1", "device")
	err := r.db.Get(&device, query, deviceID)

	return device, err
}
