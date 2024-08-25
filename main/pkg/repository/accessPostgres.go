package repository

import (
	"fmt"

	"github.com/Mamvriyskiy/database_course/main/logger"
	pkg "github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/jmoiron/sqlx"
)

type AccessHomePostgres struct {
	db *sqlx.DB
}

func NewAccessHomePostgres(db *sqlx.DB) *AccessHomePostgres {
	return &AccessHomePostgres{db: db}
}

func (r *AccessHomePostgres) AddUser(userID int, access pkg.Access) (int, error) {
	var homeID int
	const queryHomeID = `select * from getHomeID($1, $2, $3);`
	
	err := r.db.Get(&homeID, queryHomeID, userID, 4, access.Home)
	if err != nil {
		logger.Log("Error", "Get", "Error get homeID:", err, &homeID, queryHomeID, access.Home, userID)
		return 0, err
	}

	var id int
	query := fmt.Sprintf(`INSERT INTO %s (accessStatus, accessLevel) 
		values ($1, $2) RETURNING accessID`, "access")
	row := r.db.QueryRow(query, "active", access.AccessLevel)
	err = row.Scan(&id)
	if err != nil {
		logger.Log("Error", "Scan", "Error insert into access:", err, &id)
		return 0, err
	}

	var newUserID int
	queryUserID := `select c.clientID from client c where email = $1;`
	err = r.db.Get(&newUserID, queryUserID, access.Email)
	if err != nil {
		logger.Log("Error", "Get", "Error get newUserID:", err, &newUserID, queryUserID, access.Email)
		return 0, err
	}

	query2 := fmt.Sprintf("INSERT INTO %s (clientID, accessID) VALUES ($1, $2)", "accessClient")
	result, err := r.db.Exec(query2, newUserID, id)
	if err != nil {
		logger.Log("Error", "Exec", "Error insert into accessClient:", err, newUserID, id)
		return 0, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		logger.Log("Error", "RowsAffected", "Error insert into accessClient:", err, "")
		return 0, err
	}

	query3 := fmt.Sprintf("INSERT INTO %s (homeID, accessID) VALUES ($1, $2)", "accessHome")
	r.db.QueryRow(query3, homeID, id)

	return id, nil
}

func (r *AccessHomePostgres) AddOwner(userID, homeID int) (int, error) {
	var id int
	query := fmt.Sprintf(`INSERT INTO %s (accessStatus, accessLevel) 
		values ($1, $2) RETURNING accessID`, "access")
	row := r.db.QueryRow(query, "active", 4)
	err := row.Scan(&id)
	if err != nil {
		logger.Log("Error", "Scan", "Error insert into access:", err, "")
		return 0, err
	}

	query2 := fmt.Sprintf("INSERT INTO %s (clientID, accessID) VALUES ($1, $2)", "accessClient")

	result, err := r.db.Exec(query2, userID, id)
	if err != nil {
		logger.Log("Error", "Exec", "Error insert into accessClient:", err, userID, id)
		return 0, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		logger.Log("Error", "RowsAffected", "Error insert into accessClient:", err, "")
		return 0, err
	}

	query3 := fmt.Sprintf("INSERT INTO %s (homeID, accessID) VALUES ($1, $2)", "accessHome")
	r.db.QueryRow(query3, homeID, id)

	return id, nil
}

func (r *AccessHomePostgres) UpdateLevel(userID int, updateAccess pkg.Access) error {
	var homeID int
	const queryHomeID = `select * from getHomeID($1, $2, $3);`

	err := r.db.Get(&homeID, queryHomeID, userID, 4, updateAccess.Home)
	if err != nil {
		logger.Log("Error", "Get", "Error get homeID:", err, &homeID, queryHomeID, updateAccess.Home, userID)
		return err
	}

	query := `
	UPDATE access
	SET accesslevel = $1
	WHERE accessid = (
		SELECT ac.accessid FROM accessclient ac 
			JOIN client c ON c.clientid = ac.clientid
			JOIN accesshome ah ON ah.accessid = ac.accessid
		WHERE c.email = $2 AND ah.homeid = $3
	);`
	_, err = r.db.Exec(query, updateAccess.AccessLevel, updateAccess.Email, homeID)

	return err
}

func (r *AccessHomePostgres) UpdateStatus(idUser int, access pkg.AccessHome) error {
	query := `
	UPDATE access
		SET accessstatus = $1
			WHERE accessid = (
				SELECT accessid FROM accessclient WHERE clientid = $2
	);`
	_, err := r.db.Exec(query, access.AccessStatus, idUser)

	return err
}

func (r *AccessHomePostgres) GetListUserHome(userID int) ([]pkg.ClientHome, error) {
	var lists []pkg.ClientHome
	query := `SELECT h.name, c.login, c.email, a.accesslevel, a.accessstatus
FROM client c 
	JOIN accessclient ac ON c.clientid = ac.clientid
		JOIN access a ON a.accessid = ac.accessid
			JOIN accesshome ah ON ah.accessid = a.accessid
					JOIN home h ON h.homeid = ah.homeid
						WHERE ah.homeid IN (SELECT a.homeid FROM accesshome a WHERE a.accessid IN (SELECT a.accessid FROM accessclient a WHERE clientid = $1));`
	err := r.db.Select(&lists, query, userID)
	if err != nil {
		logger.Log("Error", "Select", "Error select ClientHome:", err, "")
		return nil, err
	}

	return lists, nil
}

func (r *AccessHomePostgres) DeleteUser(userID int, access pkg.Access) error {
	var homeID int
	const queryHomeID = `select * from getHomeID($1, $2, $3);`
	err := r.db.Get(&homeID, queryHomeID, userID, 4, access.Home)
	if err != nil {
		logger.Log("Error", "Get", "Error get homeID:", err, "")
		return err
	}

	var accessID int
	queryAccessID := `select a.accessid from access a where a.accessid =
	(select accessid from accesshome 
		where homeid = $1 and accessid in (select accessid from accessclient ac
			join client c on c.clientid = ac.clientid where c.email = $2));`
	err = r.db.Get(&accessID, queryAccessID, homeID, access.Email)
	if err != nil {
		logger.Log("Error", "Get", "Error get AccessID:", err, "")
		return err
	}

	queryDeleteAccessID := `delete from accessclient where accessid = $1`
	_, err = r.db.Exec(queryDeleteAccessID, accessID)

	queryDeleteAccessClientID := `delete from accesshome where accessid = $1`
	_, err = r.db.Exec(queryDeleteAccessClientID, accessID)

	queryDeleteAccessHomeID := `delete from access where accessid = $1`
	_, err = r.db.Exec(queryDeleteAccessHomeID, accessID)

	return err
}
