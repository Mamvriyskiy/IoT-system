package repository

import (
	"fmt"

	"github.com/Mamvriyskiy/database_course/main/logger"
	pkg "github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/jmoiron/sqlx"
)

type HomePostgres struct {
	db *sqlx.DB
}

func NewHomePostgres(db *sqlx.DB) *HomePostgres {
	return &HomePostgres{db: db}
}

func (r *HomePostgres) ListUserHome(userID int) ([]pkg.Home, error) {
	getHomeID := `select h.homeid, h.name, h.coords from home h 
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

func (r *HomePostgres) CreateHome(home pkg.Home) (int, error) {
	var homeID int
	query := fmt.Sprintf("INSERT INTO %s (coords, name) values ($1, $2) RETURNING homeID", "home")
	row := r.db.QueryRow(query, home.GeographCoords, home.Name)
	if err := row.Scan(&homeID); err != nil {
		logger.Log("Error", "Scan", "Error insert into home:", err, home.GeographCoords, home.Name)
		return 0, err
	}

	return homeID, nil
}

func (r *HomePostgres) DeleteHome(userID int, nameHome string) error {
	var homeID int
	
	const queryHomeID = `select * from getHomeID($1, $2, $3);`

	err := r.db.Get(&homeID, queryHomeID, userID, 4, nameHome)
	if err != nil {
		logger.Log("Error", "Get", "Error select from home:", err, userID)
		return err
	}

	query1 := `DELETE FROM access 
		WHERE homeid = $1;`
	_, err = r.db.Exec(query1, homeID)
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

func (r *HomePostgres) UpdateHome(home pkg.UpdateNameHome) error {
	var homeID int
	const queryHomeID = `select * from getHomeID($1, $2, $3);`

	err := r.db.Get(&homeID, queryHomeID, home.UserID, 4, home.LastName)
	if err != nil {
		logger.Log("Error", "Get", "Error select from home:", err, home.UserID, 4, home.LastName)
		return err
	}

	query := `UPDATE home
		SET name = $1
		WHERE homeid = $2;`
	result, err := r.db.Exec(query, home.NewName, homeID)
	if err != nil {
		logger.Log("Error", "Exec", "Error update home:", err, home.NewName, homeID)

		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		logger.Log("Error", "RowsAffected", "Error update home:", err, home.NewName, homeID)
		return err
	}

	return err
}

func (r *HomePostgres) GetHomeByID(homeID int) (pkg.Home, error) {
	var home pkg.Home
	query := fmt.Sprintf("SELECT * from %s where homeid = $1", "home")
	err := r.db.Get(&home, query, homeID)

	return home, err
}
