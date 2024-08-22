package repository

import (
	"fmt"

	"github.com/Mamvriyskiy/DBCourse/main/logger"
	pkg "github.com/Mamvriyskiy/DBCourse/main/pkg"
	"github.com/jmoiron/sqlx"
	"sync"
	"time"
)

type DeviceHistoryPostgres struct {
	db *sqlx.DB
}

func NewDeviceHistoryPostgres(db *sqlx.DB) *DeviceHistoryPostgres {
	return &DeviceHistoryPostgres{db: db}
}

func (r *DeviceHistoryPostgres) CreateDeviceHistory(userID int,
	history pkg.AddHistory,
) (int, error) {
	var homeID int
	const queryHomeID = `select h.homeid from home h 
	where h.homeid in (select a.homeid from accesshome a 
		where a.accessid in (select a.accessid from accessclient a 
			JOIN access ac ON a.accessid = ac.accessid where clientid = $1));`
	err := r.db.Get(&homeID, queryHomeID, userID)
	if err != nil {
		logger.Log("Error", "Get", "Error select from home:", err, userID)
		return 0, err
	}

	var deviceID int
	const querDeviceID = `select d.deviceid from device d join 
		devicehome dh on d.deviceid = dh.deviceid 
			where dh.homeid = $1 and d.name = $2;`
	err = r.db.Get(&deviceID, querDeviceID, homeID, history.Name)
	if err != nil {
		logger.Log("Error", "Get", "Error select from device:", err, homeID, history.Name)
		return 0, err
	}

	updateQuery := fmt.Sprintf(`UPDATE device SET status = 'Active' where deviceid = $1;`)
	_, err = r.db.Exec(updateQuery, deviceID)
	if err != nil {
		logger.Log("Error", "Exec", "Error update device status:", err, homeID, history.Name)
		return 0, err
	}

	queryUpdateStatus := fmt.Sprintf(`select update_status($1, $2);`)
	_, err = r.db.Exec(queryUpdateStatus, deviceID, "Active")
	if err != nil {
		logger.Log("Error", "Exec", "Error select from device:", err, homeID, history.Name)
		return 0, err
	}

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func(i int) {
		defer wg.Done()
		time.Sleep(time.Duration(i) * time.Second)
	}( history.TimeWork)
	
	wg.Wait()

	queryUpdateStatus = fmt.Sprintf(`select update_status($1, $2);`)
	_, err = r.db.Exec(queryUpdateStatus, deviceID, "Inactive")
	if err != nil {
		logger.Log("Error", "Exec", "Error select from device:", err, homeID, history.Name)
		return 0, err
	}

	var id int
	query := fmt.Sprintf(`INSERT INTO %s 
		(timeWork, AverageIndicator, EnergyConsumed) 
			values ($1, $2, $3) RETURNING historyDevID`, "historyDev")
	row := r.db.QueryRow(query, history.TimeWork, history.AverageIndicator, history.EnergyConsumed)
	err = row.Scan(&id)
	if err != nil {
		logger.Log("Error", "Scan", "Error instert into historyDevID:", err,
			history.TimeWork, history.AverageIndicator, history.EnergyConsumed)
		return 0, err
	}

	query = fmt.Sprintf("INSERT INTO %s (deviceID, historydevID) VALUES ($1, $2)", "historydevice")
	result, err := r.db.Exec(query, deviceID, id)
	if err != nil {
		logger.Log("Error", "Exec", "Error insert into historydevice:", err, deviceID, id)
		return 0, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		logger.Log("Error", "RowsAffected", "Error insert into historydevice:", err)
		return 0, err
	}


	return id, nil
}

func (r *DeviceHistoryPostgres) GetDeviceHistory(userID int,
	name string,
) ([]pkg.DevicesHistory, error) {
	var homeID int
	queryHomeID := `select h.homeid from home h 
	where h.homeid in (select a.homeid from accesshome a 
		where a.accessid in (select a.accessid from accessclient a 
			JOIN access ac ON a.accessid = ac.accessid
			where clientid = $1));`
	err := r.db.Get(&homeID, queryHomeID, userID)
	if err != nil {
		logger.Log("Error", "Get", "Error select from home:", err, userID)
		return nil, err
	}

	var deviceID int
	querDeviceID := `select d.deviceid from device d 
		join devicehome dh on d.deviceid = dh.deviceid 
			where dh.homeid = $1 and d.name = $2;`
	err = r.db.Get(&deviceID, querDeviceID, homeID, name)
	if err != nil {
		return nil, err
	}

	var lists []pkg.DevicesHistory
	query := `select hi.timework, hi.averageindicator, hi.energyconsumed 
		from historydev as hi join historydevice as hd on hi.historydevid = hd.historydevid 
			where hd.deviceid = $1`
	err = r.db.Select(&lists, query, deviceID)
	if err != nil {
		logger.Log("Error", "Select", "Error Select from historydev:", err, deviceID)
		return nil, err
	}

	return lists, nil
}
