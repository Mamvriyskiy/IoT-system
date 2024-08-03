package repository

import (
	"fmt"

	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/logger"
	pkg "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"github.com/jmoiron/sqlx"
)

type HomePostgres struct {
	db *sqlx.DB
}

func NewHomePostgres(db *sqlx.DB) *HomePostgres {
	return &HomePostgres{db: db}
}

func (r *HomePostgres) ListUserHome(userID int) ([]pkg.Home, error) {
	getHomeID := `select h.homeid, h.name from home h 
	where h.homeid in (select a.homeid from accesshome a 
		where a.accessid in (select a.accessid from accessclient a where clientid = $1));`

	var homeList []pkg.Home
	err := r.db.Select(&homeList, getHomeID, userID)
	if err != nil {
		logger.Log("Error", "Select", "Error Select from home:", err, getHomeID, userID)
		return nil, err
	}

	return homeList, nil
}

func (r *HomePostgres) CreateHome(ownerID int, home pkg.Home) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (ownerid, name) values ($1, $2) RETURNING homeID", "home")
	row := r.db.QueryRow(query, ownerID, home.Name)
	if err := row.Scan(&id); err != nil {
		logger.Log("Error", "Scan", "Error insert into home:", err, ownerID, home.Name)
		return 0, err
	}

	return id, nil
}

func (r *HomePostgres) DeleteHome(userID int) error {
	var homeID int
	const queryHomeID = `select h.homeid from home h 
	where h.homeid in (select a.homeid from accesshome a 
		where a.accessid in (select a.accessid from accessclient a 
			JOIN access ac ON a.accessid = ac.accessid where clientid = $1 AND accessLevel = 4));`

	err := r.db.Get(&homeID, queryHomeID, userID)
	if err != nil {
		logger.Log("Error", "Get", "Error select from home:", err, userID)
		return err
	}

	query1 := `DELETE FROM access 
		WHERE accessid IN (SELECT accessid 
			FROM accesshome WHERE homeid = $1);`
	_, err = r.db.Exec(query1, homeID)
	if err != nil {
		logger.Log("Error", "Exec", "Error delete from access:", err, homeID)
		return err
	}

	query2 := `DELETE FROM historydev 
		WHERE historydevid IN (SELECT historydevid 
			FROM historydevice WHERE deviceid 
				IN (SELECT deviceid FROM devicehome WHERE homeid = $1));`
	_, err = r.db.Exec(query2, homeID)
	if err != nil {
		logger.Log("Error", "Exec", "Error delete from historydev:", err, homeID)
		return err
	}

	query3 := `DELETE FROM device 
		WHERE deviceid IN (SELECT deviceid 
			FROM devicehome WHERE homeid = $1);`
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

func (r *HomePostgres) UpdateHome(home pkg.Home) error {
	var homeID int
	queryHomeID := `select h.homeid from home h 
	where h.homeid in (select a.homeid from accesshome a 
		where a.accessid in (select a.accessid from accessclient a 
			JOIN access ac ON a.accessid = ac.accessid where clientid = $1 AND accessLevel = 4));`

	err := r.db.Get(&homeID, queryHomeID, home.OwnerID)
	if err != nil {
		logger.Log("Error", "Get", "Error select from home:", err, home.OwnerID)
		return err
	}

	query := `UPDATE home
		SET name = $1
		WHERE homeid = $2;`
	result, err := r.db.Exec(query, home.Name, homeID)
	if err != nil {
		logger.Log("Error", "Exec", "Error update home:", err, home.Name, homeID)

		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		logger.Log("Error", "RowsAffected", "Error update home:", err, home.Name, homeID)
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
