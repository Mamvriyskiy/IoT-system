package repository

import (
	"fmt"

	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/logger"
	pkg "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"github.com/jmoiron/sqlx"
)

type AccessHomePostgres struct {
	db *sqlx.DB
}

func NewAccessHomePostgres(db *sqlx.DB) *AccessHomePostgres {
	return &AccessHomePostgres{db: db}
}

func (r *AccessHomePostgres) AddUser(userID, accessLevel int, email string) (int, error) {
	var homeID int
	const queryHomeID = `select h.homeid from home h 
	where h.homeid in (select a.homeid from accesshome a 
		where a.accessid in (select a.accessid from accessclient a 
			JOIN access ac ON a.accessid = ac.accessid where clientid = $1 AND accessLevel = 4));`
	
	err := r.db.Get(&homeID, queryHomeID, userID)
	if err != nil {
		logger.Log("Error", "Get", "Error get homeID:", err, &homeID, queryHomeID, userID)
		return 0, err
	}

	var id int
	query := fmt.Sprintf(`INSERT INTO %s (accessStatus, accessLevel) 
		values ($1, $2) RETURNING accessID`, "access")
	row := r.db.QueryRow(query, "active", accessLevel)
	err = row.Scan(&id)
	if err != nil {
		logger.Log("Error", "Scan", "Error insert into access:", err, &id)
		return 0, err
	}

	var newUserID int
	queryUserID := `select c.clientID from client c where email = $1;`
	err = r.db.Get(&newUserID, queryUserID, email)
	if err != nil {
		logger.Log("Error", "Get", "Error get newUserID:", err, &newUserID, queryUserID, email)
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

func (r *AccessHomePostgres) UpdateLevel(idUser int, access pkg.AddUserHome) error {
	var homeID int
	const queryHomeID = `select h.homeid from home h 
	where h.homeid in (select a.homeid from accesshome a 
		where a.accessid in (select a.accessid from accessclient a 
			JOIN access ac ON a.accessid = ac.accessid where clientid = $1 AND accessLevel = 4));`

	err := r.db.Get(&homeID, queryHomeID, idUser)
	if err != nil {
		logger.Log("Error", "Get", "Error get homeID:", err, &homeID, queryHomeID, idUser)
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
	_, err = r.db.Exec(query, access.AccessLevel, access.Email, homeID)

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

func (r *AccessHomePostgres) GetListUserHome(idHome int) ([]pkg.ClientHome, error) {
	var lists []pkg.ClientHome
	query := `select c.login, a.accesslevel, a.accessstatus from client c 
				join accessclient ac on c.clientid = ac.clientid
					join access a on a.accessid = ac.accessid 
						join accesshome ah on ah.accessid = a.accessid 
							where ah.homeid = $1;`
	err := r.db.Select(&lists, query, idHome)
	if err != nil {
		logger.Log("Error", "Select", "Error select ClientHome:", err, "")
		return nil, err
	}

	return lists, nil
}

func (r *AccessHomePostgres) DeleteUser(userID int, email string) error {
	var homeID int
	const queryHomeID = `select h.homeid from home h 
	where h.homeid in (select a.homeid from accesshome a 
		where a.accessid in (select a.accessid from accessclient a 
			JOIN access ac ON a.accessid = ac.accessid where clientid = $1 AND accessLevel = 4));`
	err := r.db.Get(&homeID, queryHomeID, userID)
	if err != nil {
		logger.Log("Error", "Get", "Error get homeID:", err, "")
		return err
	}

	query := `delete from access where accessid = 
	(select accessid from accesshome 
		where homeid = $1 and accessid = (select accessid from accessclient ac
			join client c on c.clientid = ac.clientid where c.email = $2));`
	_, err = r.db.Exec(query, homeID, email)

	return err
}
