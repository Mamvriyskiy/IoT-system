package repository

import (
	"fmt"

	"github.com/Mamvriyskiy/database_course/main/logger"
	pkg "github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/jmoiron/sqlx"
	// "sync"
	// "time"
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
	const queryHomeID = `select * from getHomeID($1, $2, $3);`
	err := r.db.Get(&homeID, queryHomeID, userID, 4, history.Home)
	if err != nil {
		logger.Log("Error", "Get", "Error select from home:", err, userID,  history.Home, history.Name)
		return 0, err
	}

	var deviceID int
	const querDeviceID = `select d.deviceid from device d
			where d.homeid = $1 and d.name = $2;`
	err = r.db.Get(&deviceID, querDeviceID, homeID, history.Name)
	if err != nil {
		logger.Log("Error", "Get", "Error select from device:", err, homeID, history.Name)
		return 0, err
	}

	var result int
	queryUpdateStatus := fmt.Sprintf(`select update_status($1, $2);`)
	err = r.db.Get(&result, queryUpdateStatus, deviceID, "active")
	if err != nil {
		logger.Log("Error", "Exec", "Error select from device:", err, homeID, history.Name)
		return 0, err
	}

	if result == -2 {
		return -2, err
	}

	// wg := &sync.WaitGroup{}

	// wg.Add(1)
	// go func(i int) {
	// 	defer wg.Done()
	// 	time.Sleep(10 * time.Second)
	// }(history.TimeWork)
	
	// wg.Wait()

	queryUpdateStatus = fmt.Sprintf(`select update_status($1, $2);`)
	_, err = r.db.Exec(queryUpdateStatus, deviceID, "inactive")
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
		logger.Log("Error", "Scan", "Error insert into historyDevID:", err,
			history.TimeWork, history.AverageIndicator, history.EnergyConsumed)
		return 0, err
	}

	query = fmt.Sprintf("INSERT INTO %s (deviceID, historydevID) VALUES ($1, $2)", "historydevice")
	_, err = r.db.Exec(query, deviceID, id)
	if err != nil {
		logger.Log("Error", "Exec", "Error insert into historydevice:", err, deviceID, id)
		return 0, err
	}

	return id, nil
}

func (r *DeviceHistoryPostgres) GetDeviceHistory(userID int,
	name, home string) ([]pkg.DevicesHistory, error) {
	var homeID int
	const queryHomeID = `select * from getHomeID($1, $2, $3);`
	err := r.db.Get(&homeID, queryHomeID, userID, 4, home)
	if err != nil {
		logger.Log("Error", "Get", "Error select from home:", err, userID, home)
		return nil, err
	}

	var deviceID int
	querDeviceID := `select d.deviceid from device d join 
		home h on d.homeid = h.homeid 
			where h.homeid = $1 and d.name = $2;`
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
