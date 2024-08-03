package repositoryCH

import (
	"fmt"

	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/logger"
	pkg "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	//"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"database/sql"
	"math/rand/v2"
)

type HomePostgres struct {
	db *sql.DB
}

func NewHomePostgres(db *sql.DB) *HomePostgres {
	return &HomePostgres{db: db}
}

// rows, err := r.db.Query("select resetcode from resetPswrd where token = $1", token)
// if err != nil {
// 	logger.Log("Error", "r.db.Query", "select clientID", err)
// 	return "", fmt.Errorf("client", err)
// }
// defer rows.Close()

// for rows.Next() {
// 	if err := rows.Scan(&code); err != nil {
// 		logger.Log("Error", "rows.Scan(&id)", "client", err)
// 		return "", fmt.Errorf("client", err)
// 	}
// }

func (r *HomePostgres) ListUserHome(userID int) ([]pkg.Home, error) {
	// getHomeID := `select h.homeID, h.name from home h 
	// where h.homeID in (select a.homeID from accesshome a 
	// 	where a.accessID in (select a.accessID from accessclient a where clientID = $1));`

	var homeList []pkg.Home
	// err := r.db.Select(&homeList, getHomeID, userID)
	// if err != nil {
	// 	logger.Log("Error", "Select", "Error Select from home:", err, getHomeID, userID)
	// 	return nil, err
	// }

	rows, err := r.db.Query(`select h.homeID, h.name from home h 
	where h.homeID in (select a.homeID from accesshome a 
		where a.accessID in (select a.accessID from accessclient a where clientID = $1));`, userID)
	if err != nil {
		logger.Log("Error", "r.db.Query", "select clientID", err)
		return homeList, fmt.Errorf("client", err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id int
			name string
		)
		if err := rows.Scan(&id, &name); err != nil {
			logger.Log("Error", "rows.Scan(&id)", "client", err)
			return homeList, fmt.Errorf("client", err)
		}
		var home = pkg.Home{
			ID: id,
			Name: name,
		}
		homeList = append(homeList, home)
	}

	return homeList, nil
}

func (r *HomePostgres) CreateHome(ownerID int, home pkg.Home) (int, error) {
	var id int
	// query := fmt.Sprintf("INSERT INTO %s (ownerid, name) values ($1, $2) RETURNING homeID", "home")
	// row := r.db.QueryRow(query, ownerID, home.Name)
	// if err := row.Scan(&id); err != nil {
	// 	logger.Log("Error", "Scan", "Error insert into home:", err, ownerID, home.Name)
	// 	return 0, err
	// }

	// rows, err := r.db.Query("INSERT INTO home (ownerID, name) values ($1, $2)", ownerID, home.Name)
	// if err != nil {
	// 	logger.Log("Error", "r.db.Query", "select clientID", err)
	// 	return 0, fmt.Errorf("client", err)
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	if err := rows.Scan(&id); err != nil {
	// 		logger.Log("Error", "rows.Scan(&id)", "client", err)
	// 		return 0, fmt.Errorf("client", err)
	// 	}
	// }

	// id, err = rows.LastInsertId()
    // if err != nil {
    //     return 0, fmt.Errorf("get last insert ID: %v", err)
    // }
	
	id = rand.IntN(10000)
	_, err := r.db.Query("INSERT INTO home (homeID, ownerID, name) VALUES ($1, $2, $3)", id, ownerID, home.Name)
	if err != nil {
		logger.Log("Error", "r.db.Query", "insert into home", err)
		return 0, fmt.Errorf("client", err)
	}

	return id, nil
}

func (r *HomePostgres) DeleteHome(userID int) error {
	var homeID int
	// const queryHomeID = `select h.homeID from home h 
	// where h.homeID in (select a.homeID from accesshome a 
	// 	where a.accessID in (select a.accessID from accessclient a 
	// 		JOIN access ac ON a.accessID = ac.accessID where clientID = $1 AND accessLevel = 4));`

	// err := r.db.Get(&homeID, queryHomeID, userID)
	// if err != nil {
	// 	logger.Log("Error", "Get", "Error select from home:", err, userID)
	// 	return err
	// }

	rows, err := r.db.Query(`select h.homeID from home h 
	where h.homeID in (select a.homeID from accesshome a 
	where a.accessID in (select a.accessID from accessclient a 
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

	// query1 := `DELETE FROM access 
	// 	WHERE accessID IN (SELECT accessID 
	// 		FROM accesshome WHERE homeID = $1);`
	// _, err = r.db.Exec(query1, homeID)
	// if err != nil {
	// 	logger.Log("Error", "Exec", "Error delete from access:", err, homeID)
	// 	return err
	// }

	rows, err = r.db.Query(`DELETE FROM access 
	 	WHERE accessID IN (SELECT accessID
	 		FROM accesshome WHERE homeID = $1);`, homeID)
	if err != nil {
		logger.Log("Error", "r.db.Query", "select clientID", err)
		return fmt.Errorf("client", err)
	}


	// query2 := `DELETE FROM historydev 
	// 	WHERE historydevid IN (SELECT historydevid 
	// 		FROM historydevice WHERE deviceid 
	// 			IN (SELECT deviceid FROM devicehome WHERE homeID = $1));`
	// _, err = r.db.Exec(query2, homeID)
	// if err != nil {
	// 	logger.Log("Error", "Exec", "Error delete from historydev:", err, homeID)
	// 	return err
	// }

	rows, err = r.db.Query(`DELETE FROM historydev 
	 	WHERE historyDevID IN (SELECT historyDevID
	 		FROM historydevice WHERE deviceiID 
	 			IN (SELECT deviceid FROM devicehome WHERE homeID = $1));`, homeID)
	if err != nil {
		logger.Log("Error", "r.db.Query", "select clientID", err)
		return fmt.Errorf("client", err)
	}

	// query3 := `DELETE FROM device 
	// 	WHERE deviceid IN (SELECT deviceid 
	// 		FROM devicehome WHERE homeID = $1);`
	// _, err = r.db.Exec(query3, homeID)
	// if err != nil {
	// 	logger.Log("Error", "Exec", "Error delete from device:", err, homeID)
	// 	return err
	// }

	rows, err = r.db.Query(`DELETE FROM device 
	 	WHERE deviceID IN (SELECT deviceID
	 		FROM devicehome WHERE homeID = $1);`, homeID)
	if err != nil {
		logger.Log("Error", "r.db.Query", "select clientID", err)
		return fmt.Errorf("client", err)
	}


	// query4 := `DELETE FROM home
	// 	WHERE homeID = $1;`
	// _, err = r.db.Exec(query4, homeID)
	// if err != nil {
	// 	logger.Log("Error", "Exec", "Error delete from home:", err, homeID)
	// 	return err
	// }

	rows, err = r.db.Query(`DELETE FROM home
	 	WHERE homeID = $1;`, homeID)
	if err != nil {
		logger.Log("Error", "r.db.Query", "select clientID", err)
		return fmt.Errorf("client", err)
	}

	return err
}

func (r *HomePostgres) UpdateHome(home pkg.Home) error {
	var homeID int
	// queryHomeID := `select h.homeID from home h 
	// where h.homeID in (select a.homeID from accesshome a 
	// 	where a.accessID in (select a.accessID from accessclient a 
	// 		JOIN access ac ON a.accessID = ac.accessID where clientID = $1 AND accessLevel = 4));`

	// err := r.db.Get(&homeID, queryHomeID, home.OwnerID)
	// if err != nil {
	// 	logger.Log("Error", "Get", "Error select from home:", err, home.OwnerID)
	// 	return err
	// }

	rows, err := r.db.Query(`select h.homeID from home h 
	where h.homeID in (select a.homeID from accesshome a 
	 	where a.accessID in (select a.accessID from accessclient a 
	 		JOIN access ac ON a.accessID = ac.accessID where clientID = $1 AND accessLevel = 4));`, home.OwnerID)
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

	// query := `UPDATE home
	// 	SET name = $1
	// 	WHERE homeID = $2;`
	// result, err := r.db.Exec(query, home.Name, homeID)
	// if err != nil {
	// 	logger.Log("Error", "Exec", "Error update home:", err, home.Name, homeID)

	// 	return err
	// }

	rows, err = r.db.Query(`UPDATE home
	 	SET name = $1
	 	WHERE homeID = $2;`, home.Name, homeID)
	if err != nil {
		logger.Log("Error", "r.db.Query", "select clientID", err)
		return fmt.Errorf("client", err)
	}

	// _, err = result.RowsAffected()
	// if err != nil {
	// 	logger.Log("Error", "RowsAffected", "Error update home:", err, home.Name, homeID)
	// 	return err
	// }

	return nil
}

func (r *HomePostgres) GetHomeByID(homeID int) (pkg.Home, error) {
	var home pkg.Home
	// query := fmt.Sprintf("SELECT * from %s where homeID = $1", "home")
	// err := r.db.Get(&home, query, homeID)
	rows, err := r.db.Query("SELECT * from home where homeID = $1", homeID)
	if err != nil {
		logger.Log("Error", "r.db.Query", "select clientID", err)
		return home, fmt.Errorf("client", err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			OwnerID int
			name string
		)
		if err := rows.Scan(&homeID, &OwnerID, &name); err != nil {
			logger.Log("Error", "rows.Scan(&id)", "client", err)
			return home, fmt.Errorf("client", err)
		}

		home.Name = name
		home.OwnerID = OwnerID
		home.ID = homeID
	}

	return home, nil
}

