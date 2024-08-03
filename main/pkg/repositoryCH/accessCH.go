package repositoryCH

import (
	"fmt"

	"github.com/Mamvriyskiy/DBCourse/main/logger"
	pkg "github.com/Mamvriyskiy/DBCourse/main/pkg"
	//"github.com/jmoiron/sqlx"
	"database/sql"
	//"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type AccessHomePostgres struct {
	db *sql.DB
}

func NewAccessHomePostgres(db *sql.DB) *AccessHomePostgres {
	return &AccessHomePostgres{db: db}
}

func (r *AccessHomePostgres) AddUser(userID, accessLevel int, email string) (int, error) {
	var homeID int
	// const queryHomeID = `select h.homeID from home h 
	// where h.homeID in (select a.homeID from accessHome a 
	// where a.accessID in (select a.accessID from accessClient a 
	// JOIN access ac ON a.accessID = ac.accessID where clientID = $1 AND accessLevel = 4));`
	
	// err := r.db.Get(&homeID, queryHomeID, userID)
	// if err != nil {
	// 	logger.Log("Error", "Get", "Error get homeID:", err, &homeID, queryHomeID, userID)
	// 	return 0, err
	// }

	rows, err := r.db.Query(`select h.homeID from home h 
		where h.homeID in (select a.homeID from accessHome a 
		where a.accessID in (select a.accessID from accessClient a 
		JOIN access ac ON a.accessID = ac.accessID where clientID = $1 AND accessLevel = 4));`, userID)
	if err != nil {
		logger.Log("Error", "r.db.Query", "home", err)
		return 0, fmt.Errorf("client", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&homeID); err != nil {
			logger.Log("Error", "rows.Scan(&id)", "client", err)
			return 0, fmt.Errorf("client", err)
		}
	}


	var id int
	// query := fmt.Sprintf(`INSERT INTO %s (accessStatus, accessLevel) 
	// 	values ($1, $2) RETURNING accessID`, "access")
	// row := r.db.QueryRow(query, "active", accessLevel)
	// err = row.Scan(&id)
	// if err != nil {
	// 	logger.Log("Error", "Scan", "Error insert into access:", err, &id)
	// 	return 0, err
	// }

	rows, err = r.db.Query(`INSERT INTO access (accessStatus, accessLevel) 
	values ($1, $2)`, "active", accessLevel)
	if err != nil {
		logger.Log("Error", "r.db.Query", "access", err)
		return 0, fmt.Errorf("client", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			logger.Log("Error", "rows.Scan(&id)", "access", err)
			return 0, fmt.Errorf("client", err)
		}
	}

	var newUserID int
	// queryUserID := `select c.clientID from client c where email = $1;`
	// err = r.db.Get(&newUserID, queryUserID, email)
	// if err != nil {
	// 	logger.Log("Error", "Get", "Error get newUserID:", err, &newUserID, queryUserID, email)
	// 	return 0, err
	// }

	rows, err = r.db.Query(`select c.clientID from client c where email = $1;`, email)
	if err != nil {
		logger.Log("Error", "r.db.Query", "access", err)
		return 0, fmt.Errorf("client", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&newUserID); err != nil {
			logger.Log("Error", "rows.Scan(&id)", "access", err)
			return 0, fmt.Errorf("client", err)
		}
	}

	// query2 := fmt.Sprintf("INSERT INTO %s (clientID, accessID) VALUES ($1, $2)", "accessClient")
	// result, err := r.db.Exec(query2, newUserID, id)
	// if err != nil {
	// 	logger.Log("Error", "Exec", "Error insert into accessClient:", err, newUserID, id)
	// 	return 0, err
	// }

	rows, err = r.db.Query(`INSERT INTO accessClient (clientID, accessID) VALUES ($1, $2);`, newUserID, id)
	if err != nil {
		logger.Log("Error", "r.db.Query", "access", err)
		return 0, fmt.Errorf("client", err)
	}


	// _, err = result.RowsAffected()
	// if err != nil {
	// 	logger.Log("Error", "RowsAffected", "Error insert into accessClient:", err, "")
	// 	return 0, err
	// }

	rows, err = r.db.Query(`INSERT INTO accessHome (homeID, accessID) VALUES ($1, $2)`, homeID, id)
	if err != nil {
		logger.Log("Error", "r.db.Query", "access", err)
		return 0, fmt.Errorf("client", err)
	}

	// query3 := fmt.Sprintf("INSERT INTO %s (homeID, accessID) VALUES ($1, $2)", "accessHome")
	// r.db.QueryRow(query3, homeID, id)

	return 0, nil
}

func (r *AccessHomePostgres) AddOwner(userID, homeID int) (int, error) {
	var id int
	// query := fmt.Sprintf(`INSERT INTO %s (accessStatus, accessLevel) 
	// 	values ($1, $2) RETURNING accessID`, "access")
	// row := r.db.QueryRow(query, "active", 4)
	// err := row.Scan(&id)
	// if err != nil {
	// 	logger.Log("Error", "Scan", "Error insert into access:", err, "")
	// 	return 0, err
	// }

	rows, err := r.db.Query(`INSERT INTO access (accessStatus, accessLevel) 
	values ($1, $2)`, "active", 4)
	if err != nil {
		logger.Log("Error", "r.db.Query", "home", err)
		return 0, fmt.Errorf("client", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			logger.Log("Error", "rows.Scan(&id)", "access", err)
			return 0, fmt.Errorf("client", err)
		}
	}


	// query2 := fmt.Sprintf("INSERT INTO %s (clientID, accessID) VALUES ($1, $2)", "accessClient")

	// result, err := r.db.Exec(query2, userID, id)
	// if err != nil {
	// 	logger.Log("Error", "Exec", "Error insert into accessClient:", err, userID, id)
	// 	return 0, err
	// }

	rows, err = r.db.Query(`INSERT INTO accessClient (clientID, accessID) VALUES ($1, $2)`, userID, id)
	if err != nil {
		logger.Log("Error", "r.db.Query", "home", err)
		return 0, fmt.Errorf("client", err)
	}

	// _, err = result.RowsAffected()
	// if err != nil {
	// 	logger.Log("Error", "RowsAffected", "Error insert into accessClient:", err, "")
	// 	return 0, err
	// }

	// query3 := fmt.Sprintf("INSERT INTO %s (homeID, accessID) VALUES ($1, $2)", "accessHome")
	// r.db.QueryRow(query3, homeID, id)

	rows, err = r.db.Query(`INSERT INTO accessHome (homeID, accessID) VALUES ($1, $2)`, homeID, id)
	if err != nil {
		logger.Log("Error", "r.db.Query", "home", err)
		return 0, fmt.Errorf("client", err)
	}

	return id, nil
}

func (r *AccessHomePostgres) UpdateLevel(idUser int, access pkg.AddUserHome) error {
	var homeID int
	// const queryHomeID = `select h.homeID from home h 
	// where h.homeID in (select a.homeID from accessHome a 
	// 	where a.accessID in (select a.accessID from accessClient a 
	// 		JOIN access ac ON a.accessID = ac.accessID where clientID = $1 AND accessLevel = 4));`

	// err := r.db.Get(&homeID, queryHomeID, idUser)
	// if err != nil {
	// 	logger.Log("Error", "Get", "Error get homeID:", err, &homeID, queryHomeID, idUser)
	// 	return err
	// }

	rows, err := r.db.Query(`select h.homeID from home h 
		where h.homeID in (select a.homeID from accessHome a 
		where a.accessID in (select a.accessID from accessClient a 
		JOIN access ac ON a.accessID = ac.accessID where clientID = $1 AND accessLevel = 4));`, idUser)

	if err != nil {
		logger.Log("Error", "r.db.Query", "home", err)
		return fmt.Errorf("client", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&homeID); err != nil {
			logger.Log("Error", "rows.Scan(&id)", "client", err)
			return fmt.Errorf("client", err)
		}
	}

	// query := `
	// UPDATE access
	// SET accesslevel = $1
	// WHERE accessID = (
	// 	SELECT ac.accessID FROM accessClient ac 
	// 		JOIN client c ON c.clientID = ac.clientID
	// 		JOIN accessHome ah ON ah.accessID = ac.accessID
	// 	WHERE c.email = $2 AND ah.homeID = $3
	// );`
	// _, err = r.db.Exec(query, access.AccessLevel, access.Email, homeID)

	rows, err = r.db.Query(`UPDATE access
	SET accesslevel = $1
	WHERE accessID = (
		SELECT ac.accessID FROM accessClient ac 
			JOIN client c ON c.clientID = ac.clientID
			JOIN accessHome ah ON ah.accessID = ac.accessID
		WHERE c.email = $2 AND ah.homeID = $3`, access.AccessLevel, access.Email, homeID)
		
	if err != nil {
		logger.Log("Error", "r.db.Query", "home", err)
		return fmt.Errorf("client", err)
	}

	return err
}

func (r *AccessHomePostgres) UpdateStatus(idUser int, access pkg.AccessHome) error {
	query := `
	UPDATE access
		SET accessStatus = $1
			WHERE accessID = (
				SELECT accessID FROM accessClient WHERE clientID = $2
	);`
	_, err := r.db.Exec(query, access.AccessStatus, idUser)

	return err
}

func (r *AccessHomePostgres) GetListUserHome(idHome int) ([]pkg.ClientHome, error) {
	var lists []pkg.ClientHome
	// query := `select c.login, a.accesslevel, a.accessStatus from client c 
	// 			join accessClient ac on c.clientID = ac.clientID
	// 				join access a on a.accessID = ac.accessID 
	// 					join accessHome ah on ah.accessID = a.accessID 
	// 						where ah.homeID = $1;`
	// err := r.db.Select(&lists, query, idHome)
	// if err != nil {
	// 	logger.Log("Error", "Select", "Error select ClientHome:", err, "")
	// 	return nil, err
	// }

	rows, err := r.db.Query(`select c.login, a.accesslevel, a.accessStatus from client c 
	join accessClient ac on c.clientID = ac.clientID
		join access a on a.accessID = ac.accessID 
			join accessHome ah on ah.accessID = a.accessID 
				where ah.homeID = $1;`, idHome)

	if err != nil {
		logger.Log("Error", "r.db.Query", "home", err)
		return lists, fmt.Errorf("client", err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			name, st string
			level int
		)
		if err := rows.Scan(&name, &st, &level); err != nil {
			logger.Log("Error", "rows.Scan(&id)", "client", err)
			return lists, fmt.Errorf("client", err)
		}

		var el = pkg.ClientHome{
			Username     : name,
			AccessStatus : st,
			AccessLevel  : level,
		}

		lists = append(lists, el)
	}

	return lists, nil
}

func (r *AccessHomePostgres) DeleteUser(userID int, email string) error {
	var homeID int
	// const queryHomeID = `select h.homeID from home h 
	// where h.homeID in (select a.homeID from accessHome a 
	// 	where a.accessID in (select a.accessID from accessClient a 
	// 		JOIN access ac ON a.accessID = ac.accessID where clientID = $1 AND accessLevel = 4));`
	// err := r.db.Get(&homeID, queryHomeID, userID)
	// if err != nil {
	// 	logger.Log("Error", "Get", "Error get homeID:", err, "")
	// 	return err
	// }

	rows, err := r.db.Query(`select h.homeID from home h 
	where h.homeID in (select a.homeID from accessHome a 
		where a.accessID in (select a.accessID from accessClient a 
			JOIN access ac ON a.accessID = ac.accessID where clientID = $1 AND accessLevel = 4));`, userID)

	if err != nil {
		logger.Log("Error", "r.db.Query", "home", err)
		return fmt.Errorf("client", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&homeID); err != nil {
			logger.Log("Error", "rows.Scan(&id)", "client", err)
			return fmt.Errorf("client", err)
		}
	}

	// query := `delete from access where accessID = 
	// (select accessID from accessHome 
	// 	where homeID = $1 and accessID = (select accessID from accessClient ac
	// 		join client c on c.clientID = ac.clientID where c.email = $2));`
	// _, err = r.db.Exec(query, homeID, email)

	rows, err = r.db.Query(`delete from access where accessID = 
	(select accessID from accessHome 
		where homeID = $1 and accessID = (select accessID from accessClient ac
			join client c on c.clientID = ac.clientID where c.email = $2));`, homeID, email)

	if err != nil {
		logger.Log("Error", "r.db.Query", "home", err)
		return fmt.Errorf("client", err)
	}

	return err
}

