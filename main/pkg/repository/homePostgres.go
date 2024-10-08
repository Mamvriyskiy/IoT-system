package repository

import (
	"fmt"

	"github.com/Mamvriyskiy/database_course/main/logger"
	pkg "github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/jmoiron/sqlx"
	"github.com/google/uuid"
)

type HomePostgres struct {
	db *sqlx.DB
}

func NewHomePostgres(db *sqlx.DB) *HomePostgres {
	return &HomePostgres{db: db}
}

func (r *HomePostgres) ListUserHome(userID string) ([]pkg.Home, error) {
	getHomeID := `select * from home h 
	where h.homeid in (select a.homeid from access a 
		where a.clientid = $1);`

	var homeList []pkg.Home
	err := r.db.Select(&homeList, getHomeID, userID)
	if err != nil {
		logger.Log("Error", "Select", "Error Select from home:", err, getHomeID, userID)
		return nil, err
	}

	return homeList, nil
}

func (r *HomePostgres) CreateHome(home pkg.Home) (string, error) {
	id := uuid.New()
	var homeID string
	query := fmt.Sprintf("INSERT INTO %s (longitude, latitude, name, homeID) values ($1, $2, $3, $4) RETURNING homeid", "home")
	row := r.db.QueryRow(query, home.Longitude, home.Latitude, home.Name, id)
	if err := row.Scan(&homeID); err != nil {
		logger.Log("Error", "Scan", "Error insert into home:", err, home.Longitude, home.Latitude, home.Name)
		return "", err
	}

	return homeID, nil
}

func (r *HomePostgres) DeleteHome(homeID string) error {
	query1 := `DELETE FROM access 
		WHERE homeid = $1;`
	_, err := r.db.Exec(query1, homeID)
	if err != nil {
		logger.Log("Error", "Exec", "Error delete from access:", err, homeID)
		return err
	}

	query2 := `DELETE FROM historydev 
		WHERE historydevid IN (SELECT historydevid 
			FROM historydevice hd join device d on d.deviceid = hd.deviceid WHERE d.homeid = $1);`
	_, err = r.db.Exec(query2, homeID)
	if err != nil {
		logger.Log("Error", "Exec", "Error delete from historydev:", err, homeID)
		return err
	}

	query3 := `DELETE FROM device 
		WHERE homeid = $1;`

	_, err = r.db.Exec(query3, homeID)
	if err != nil {
		logger.Log("Error", "Exec", "Error delete from device:", err, homeID)
		return err
	}


	query4 := `DELETE FROM home 
		WHERE homeid = $1;`

	_, err = r.db.Exec(query4, homeID)
	if err != nil {
		logger.Log("Error", "Exec", "Error delete from home:", err, homeID)
		return err
	}

	return err
}

func (r *HomePostgres) UpdateHome(homeID, name string) error {
	query := `UPDATE home
		SET name = $1
		WHERE homeid = $2;`
	result, err := r.db.Exec(query, name, homeID)
	if err != nil {
		logger.Log("Error", "Exec", "Error update home:", err, name, homeID)

		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		logger.Log("Error", "RowsAffected", "Error update home:", err, name, homeID)
		return err
	}

	return err
}

func (r *HomePostgres) GetHomeByID(homeID string) (pkg.Home, error) {
	var home pkg.Home
	query := fmt.Sprintf("SELECT * from %s where homeid = $1", "home")
	err := r.db.Get(&home, query, homeID)

	return home, err
}
