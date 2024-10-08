package repository

import (
	"fmt"

	"github.com/Mamvriyskiy/database_course/main/logger"
	pkg "github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/jmoiron/sqlx"
	"github.com/google/uuid"
)

type AccessHomePostgres struct {
	db *sqlx.DB
}

func NewAccessHomePostgres(db *sqlx.DB) *AccessHomePostgres {
	return &AccessHomePostgres{db: db}
}

func (r *AccessHomePostgres) AddUser(homeID string, access pkg.Access) (string, error) {
	var userID string
	queryUserID := `select c.clientID from client c where email = $1;`
	err := r.db.Get(&userID, queryUserID, access.Email)
	if err != nil {
		logger.Log("Error", "Get", "Error get newUserID:", err, &userID, queryUserID, access.Email)
		return "", err
	}

	accessID := uuid.New()

	var id string
	query := fmt.Sprintf(`INSERT INTO %s (accessStatus, accessLevel, homeid, clientid, accessID) 
		values ($1, $2, $3, $4, $5) RETURNING accessID`, "access")
	row := r.db.QueryRow(query, "active", access.AccessLevel, homeID, userID, accessID)
	err = row.Scan(&id)
	if err != nil {
		logger.Log("Error", "Scan", "Error insert into access:", err, access.AccessLevel, homeID, userID, id)
		return "", err
	}

	return id, nil
}

func (r *AccessHomePostgres) AddOwner(userID, homeID string) (string, error) {
	accessID := uuid.New()

	var id string
	query := fmt.Sprintf(`INSERT INTO %s (accessStatus, accessLevel, clientid, homeid, accessID) 
		values ($1, $2, $3, $4, $5) RETURNING accessID`, "access")
	row := r.db.QueryRow(query, "active", 4, userID, homeID, accessID)
	err := row.Scan(&id)
	if err != nil {
		logger.Log("Error", "Scan", "Error insert into access:", err, "")
		return "", err
	}

	return id, nil
}

func (r *AccessHomePostgres) UpdateLevel(accessID string, updateAccess pkg.Access) error {
	query := `
	UPDATE access
	SET accesslevel = $1
	WHERE accessID = $2;`
	_, err := r.db.Exec(query, updateAccess.AccessLevel, accessID)

	return err
}

func (r *AccessHomePostgres) UpdateStatus(accessID string, access pkg.AccessHome) error {
	query := `
	UPDATE access
	SET accessstatus = $1
	WHERE accessID = $2;`
	_, err := r.db.Exec(query, access.AccessStatus, accessID)

	return err
}

func (r *AccessHomePostgres) GetListUserHome(homeID string) ([]pkg.ClientHome, error) {
	var lists []pkg.ClientHome
	query := `SELECT h.name, c.login, c.email, a.accesslevel, a.accessstatus
		FROM client c 
			JOIN access a ON a.clientid = c.clientid
				JOIN home h ON h.homeid = a.homeid
					WHERE h.homeid = $1;`
					
	err := r.db.Select(&lists, query, homeID)
	if err != nil {
		logger.Log("Error", "Select", "Error select ClientHome:", err, "")
		return nil, err
	}

	return lists, nil
}

func (r *AccessHomePostgres) DeleteUser(accessID string) error {
	queryDeleteAccessHomeID := `delete from access where accessid = $1`
	_, err := r.db.Exec(queryDeleteAccessHomeID, accessID)

	return err
}

func (r *AccessHomePostgres) GetInfoAccessByID(accessID string) (pkg.Access, error) {
	var access pkg.Access
	query := `SELECT c.login, c.email, a.accesslevel, a.accessid
              FROM client c 
              JOIN access a ON a.clientid = c.clientid
              WHERE a.accessID = $1;`

	err := r.db.Get(&access, query, accessID) // Используем одну структуру
	if err != nil {
		logger.Log("Error", "Get", "Ошибка при получении Access информации:", err, query, accessID)
		return access, err
	}

	return access, nil
}

