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

func (r *DeviceHistoryPostgres) CreateDeviceHistory(deviceID string,
	history pkg.AddHistory) (int, error) {
	var result int
	queryUpdateStatus := fmt.Sprintf(`select update_status($1, $2);`)
	err := r.db.Get(&result, queryUpdateStatus, deviceID, "active")
	if err != nil {
		logger.Log("Error", "Exec", "Error select from device:", err, deviceID)
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
		logger.Log("Error", "Exec", "Error select from device:", err, deviceID)
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

func (r *DeviceHistoryPostgres) GetDeviceHistory(deviceID string) ([]pkg.DevicesHistory, error) {
	var lists []pkg.DevicesHistory
	query := `select hi.timework, hi.averageindicator, hi.energyconsumed 
		from historydev as hi join historydevice as hd on hi.historydevid = hd.historydevid 
			where hd.deviceid = $1`
	err := r.db.Select(&lists, query, deviceID)
	if err != nil {
		logger.Log("Error", "Select", "Error Select from historydev:", err, deviceID)
		return nil, err
	}

	return lists, nil
}
