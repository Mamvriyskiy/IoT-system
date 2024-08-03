package repositoryCH

import (
	"fmt"

	"github.com/Mamvriyskiy/DBCourse/main/logger"
	pkg "github.com/Mamvriyskiy/DBCourse/main/pkg"
	// "github.com/jmoiron/sqlx"
	//"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"database/sql"
	"math/rand/v2"
)

type DevicePostgres struct {
	db *sql.DB
}

func NewDevicePostgres(db *sql.DB) *DevicePostgres {
	return &DevicePostgres{db: db}
}

func (r *DevicePostgres) CreateDevice(userID int, device *pkg.Devices) (int, error) {
	var homeID int
	// const queryHomeID = `select h.homeID from home h 
	// where h.homeID in (select a.homeID from accessHome a 
	// 	where a.accessID in (select a.accessID from accessClient a 
	// 		JOIN access ac ON a.accessID = ac.accessID where clientID = $1 AND accessLevel = 4));`

	// err := r.db.Get(&homeID, queryHomeID, userID)
	// if err != nil {
	// 	logger.Log("Error", "Get", "Error get homeID:", err, &homeID, userID)
	// 	return 0, err
	// }

	rows, err := r.db.Query(`select h.homeID from home h 
	where h.homeID in (select a.homeID from accessHome a 
	 	where a.accessID in (select a.accessID from accessClient a 
	 		JOIN access ac ON a.accessID = ac.accessID where clientID = $1 AND accessLevel = 4));`, userID)
	if err != nil {
		logger.Log("Error", "r.db.Query", "select clientID", err)
		return 0, fmt.Errorf("client", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&homeID); err != nil {
			logger.Log("Error", "rows.Scan(&id)", "client", err)
			return 0, fmt.Errorf("client", err)
		}
	}

	var id int
	// query := fmt.Sprintf(`INSERT INTO %s (name, TypeDevice, Status, 
	// 	Brand, PowerConsumption, MinParametr, MaxParametr) 
	// 		values ($1, $2, $3, $4, $5, $6, $7) RETURNING deviceID`, "device")
	// row := r.db.QueryRow(query, device.Name, device.TypeDevice,
	// 	device.Status, device.Brand, device.PowerConsumption,
	// 	device.MinParameter, device.MaxParameter)

	// err = row.Scan(&id)
	// if err != nil {
	// 	logger.Log("Error", "Scan", "Error insert into device:", err, &id)
	// 	return 0, err
	// }

	// rows, err = r.db.Query(`INSERT INTO device (name, typeDevice, status, 
	// 	 	brand, powerConsumption, minParametr, maxParametr) 
	// 			values ($1, $2, $3, $4, $5, $6, $7)`, device.Name, device.TypeDevice,
	// 			 	device.Status, device.Brand, device.PowerConsumption,
	// 			 	device.MinParameter, device.MaxParameter)
	// if err != nil {
	// 	logger.Log("Error", "r.db.Query", "select clientID", err)
	// 	return 0, fmt.Errorf("client", err)
	// }

	id = rand.IntN(10000)
	_, err = r.db.Query(`INSERT INTO device (deviceID, name, typeDevice, status, 
		brand, powerConsumption, minParametr, maxParametr) 
		   values ($1, $2, $3, $4, $5, $6, $7, $8)`, id, device.Name, device.TypeDevice,
				device.Status, device.Brand, device.PowerConsumption,
				device.MinParameter, device.MaxParameter)
	if err != nil {
		logger.Log("Error", "r.db.Query", "insert into home", err)
		return 0, fmt.Errorf("client", err)
	}

	// for rows.Next() {
	// 	if err := rows.Scan(&id); err != nil {
	// 		logger.Log("Error", "rows.Scan(&id)", "client", err)
	// 		return 0, fmt.Errorf("client", err)
	// 	}
	// }

	rows, err = r.db.Query(`INSERT INTO deviceHome (homeID, deviceID) VALUES ($1, $2)`, homeID, id)
	if err != nil {
		logger.Log("Error", "r.db.Query", "select clientID", err)
		return 0, fmt.Errorf("client", err)
	}

	// query1 := fmt.Sprintf("INSERT INTO %s (homeID, deviceID) VALUES ($1, $2)", "deviceHome")
	// _ = r.db.QueryRow(query1, homeID, id)

	return id, nil
}

func (r *DevicePostgres) DeleteDevice(userID int, name string) error {
	var homeID int
	// const queryHomeID = `select h.homeID from home h 
	// where h.homeID in (select a.homeID from accessHome a 
	// 	where a.accessID in (select a.accessID from accessClient a 
	// 		JOIN access ac ON a.accessID = ac.accessID where clientID = $1 AND accessLevel = 4));`

	// err := r.db.Get(&homeID, queryHomeID, userID)
	// if err != nil {
	// 	logger.Log("Error", "Get", "Error get homeID:", err, &homeID, userID)
	// 	return err
	// }

	rows, err := r.db.Query(`select h.homeID from home h 
	where h.homeID in (select a.homeID from accessHome a 
	 	where a.accessID in (select a.accessID from accessClient a 
	 		JOIN access ac ON a.accessID = ac.accessID where clientID = $1 AND accessLevel = 4));`, userID)
	if err != nil {
		logger.Log("Error", "r.db.Query", "select clientID", err)
		return fmt.Errorf("client", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&homeID); err != nil {
			logger.Log("Error", "rows.Scan(&id)", "client", err)
			return fmt.Errorf("client", err)
		}
	}

	var deviceID int
	// queryDeviceID := `select d.deviceID from device d 
	// join devicehome d2 on d.deviceID = d2.deviceID 
	// 	where d2.homeID = $1 and d.name = $2;`
	// err = r.db.Get(&deviceID, queryDeviceID, homeID, name)
	// if err != nil {
	// 	logger.Log("Error", "Get", "Error get deviceID:", err, &deviceID, homeID, name)
	// 	return err
	// }

	rows, err = r.db.Query(`select d.deviceID from device d 
	 	join devicehome d2 on d.deviceID = d2.deviceID 
	 	where d2.homeID = $1 and d.name = $2;`, homeID, name)
	if err != nil {
		logger.Log("Error", "r.db.Query", "select clientID", err)
		return fmt.Errorf("client", err)
	}

	// query := `DELETE FROM historydev
	// 		WHERE historydevid IN 
	// 			(SELECT h2.historydevid FROM historydevice h2 
	// 				WHERE h2.deviceID = $1);`

	// _, err = r.db.Exec(query, deviceID)
	// if err != nil {
	// 	logger.Log("Error", "Exec", "Error delete historydev:", err, deviceID)
	// 	return err
	// }

	rows, err = r.db.Query(`DELETE FROM historydev
	WHERE historydevid IN 
		(SELECT h2.historydevid FROM historydevice h2 
			WHERE h2.deviceID = $1);`, deviceID)
	if err != nil {
		logger.Log("Error", "r.db.Query", "select clientID", err)
		return fmt.Errorf("client", err)
	}

	// query = `DELETE FROM device 
	// 			where deviceID = $1;`
	// _, err = r.db.Exec(query, deviceID)
	// if err != nil {
	// 	logger.Log("Error", "Exec", "Error delete device:", err, deviceID)
	// 	return err
	// }

	rows, err = r.db.Query(`DELETE FROM device 
			where deviceID = $1;`, deviceID)
	if err != nil {
		logger.Log("Error", "r.db.Query", "select clientID", err)
		return fmt.Errorf("client", err)
	}

	return err
}

func (r *DevicePostgres) GetDeviceByID(deviceID int) (pkg.Devices, error) {
	var device pkg.Devices
	// query := fmt.Sprintf("SELECT * from %s where deviceID = $1", "device")
	// err := r.db.Get(&device, query, deviceID)

	rows, err := r.db.Query(`SELECT * from device where deviceID = $1`, deviceID)
	if err != nil {
		logger.Log("Error", "r.db.Query", "select clientID", err)
		return device, fmt.Errorf("client", err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			Name             string 
			TypeDevice       string 
			Status           string 
			Brand            string 
			DeviceID         int    
			PowerConsumption int    
			MinParameter     int    
			MaxParameter     int    
		)

		if err := rows.Scan(&DeviceID, &Name, &TypeDevice, &Status, 
			&Brand, &MaxParameter, &MinParameter, &PowerConsumption); err != nil {
			logger.Log("Error", "rows.Scan(&id)", "client", err)
			return device, fmt.Errorf("client", err)
		}

		device.Name = Name
		device.TypeDevice = TypeDevice
		device.Status = Status
		device.Brand = Brand
		device.DeviceID = DeviceID
		device.PowerConsumption = PowerConsumption
		device.MinParameter = MinParameter
		device.MaxParameter = MaxParameter
	}

	return device, err
}
