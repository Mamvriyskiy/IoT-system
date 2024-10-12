package repository

import (
	"fmt"

	"github.com/Mamvriyskiy/database_course/main/logger"
	pkg "github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/jmoiron/sqlx"
	"github.com/google/uuid"
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
	history pkg.HistoryService) (string, error) {
	var result int
	queryUpdateStatus := fmt.Sprintf(`select update_status($1, $2);`)
	err := r.db.Get(&result, queryUpdateStatus, deviceID, "active")
	if err != nil {
		logger.Log("Error", "Exec", "Error select from device:", err, deviceID)
		return "", err
	}

	if result == -2 {
		return "-2", err
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
		return "", err
	}

	historyDevID := uuid.New()
	var id string
	query := fmt.Sprintf(`INSERT INTO %s 
		(timeWork, AverageIndicator, EnergyConsumed, historyDevID) 
			values ($1, $2, $3, $4) RETURNING historyDevID`, "historyDev")
	row := r.db.QueryRow(query, history.TimeWork, history.AverageIndicator, history.EnergyConsumed, historyDevID)
	err = row.Scan(&id)
	if err != nil {
		logger.Log("Error", "Scan", "Error insert into historyDevID:", err,
			history.TimeWork, history.AverageIndicator, history.EnergyConsumed)
		return "", err
	}

	query = fmt.Sprintf("INSERT INTO %s (deviceID, historydevID) VALUES ($1, $2)", "historydevice")
	_, err = r.db.Exec(query, deviceID, id)
	if err != nil {
		logger.Log("Error", "Exec", "Error insert into historydevice:", err, deviceID, id)
		return "", err
	}

	return id, nil
}

func (r *DeviceHistoryPostgres) GetDeviceHistory(deviceID string) ([]pkg.DevicesHistoryData, error) {
	var lists []pkg.DevicesHistoryData
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
