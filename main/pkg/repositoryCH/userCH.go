package repositoryCH

import (
	"fmt"
	"context"

	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/logger"
	pkg "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	//"github.com/jmoiron/sqlx"
	//"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"database/sql"

)

type UserPostgres struct {
	db *sql.DB
}

func NewUserPostgres(db *sql.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) CreateUser(user pkg.User) (int, error) {
	var id int
	ctx := context.Background()
    // Insert statement without RETURNING
    query := `INSERT INTO client (password, login, email) 
              VALUES ($1, $2, $3)`

    _, err := r.db.ExecContext(ctx, query, user.Password, user.Username, user.Email)
    if err != nil {
        logger.Log("Error", "r.db.ExecContext", "Error inserting into client:", err, user.Password, user.Username, user.Email)
        return 0, fmt.Errorf("error inserting into client: %w", err)
    }

	rows, err := r.db.Query("SELECT clientID FROM client where email = ?", user.Email)
    if err != nil {
		logger.Log("Error", "r.db.Query", "select clientID", err)
        return 0, fmt.Errorf("select clientID %q: %v", user.Email, err)
    }
    defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			logger.Log("Error", "rows.Scan(&id)", "client", err)
			return 0, fmt.Errorf("client %q: %v", id, err)
		}
	}

    return id, err 
}

func (r *UserPostgres) GetUser(login, password string) (pkg.User, error) {
	var user pkg.User
	// query := fmt.Sprintf("SELECT clientid from %s where login = $1 and password = $2", "client")
	// err := r.db.Get(&user, query, login, password)

	rows, err := r.db.Query("SELECT clientID from client where login = $1 and password = $2", login, password)
    if err != nil {
		logger.Log("Error", "r.db.Query", "select clientID", err)
        return user, fmt.Errorf("select clientID %q: %v", user.Email, err)
    }
    defer rows.Close()

	for rows.Next() {
		var (
            id int
        )
		if err := rows.Scan(&id); err != nil {
			logger.Log("Error", "rows.Scan(&id)", "client", err)
			return user, fmt.Errorf("client %q: %v", id, err)
		}

		user.ID = id
	}

	return user, nil
}

func (r *UserPostgres) ChangePassword(password, token string) (error) {
	_, err := r.db.Query(`UPDATE client SET password = $1 WHERE clientid = 
		(SELECT clientid FROM resetpswrd rp
			WHERE rp.token = $2`, password, token)
    if err != nil {
		logger.Log("Error", "r.db.Query", "select clientID", err)
        return fmt.Errorf("select clientID", err)
    }

	return nil
} 

func (r *UserPostgres) GetCode(token string) (string, error) {
	var code string
	// query := fmt.Sprintf("select resetcode from resetPswrd where token = $1")
	// err := r.db.Get(&code, query, token)

	rows, err := r.db.Query("select resetcode from resetPswrd where token = $1", token)
    if err != nil {
		logger.Log("Error", "r.db.Query", "select clientID", err)
        return "", fmt.Errorf("client", err)
    }
    defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&code); err != nil {
			logger.Log("Error", "rows.Scan(&id)", "client", err)
			return "", fmt.Errorf("client", err)
		}
	}

	return code, nil
}

func (r *UserPostgres) AddCode(email pkg.Email) error {
	// var userID int
	// query := fmt.Sprintf("select clientID from client where email = $1")
	// err := r.db.Get(&userID, query, email.Email)

	// query = fmt.Sprintf(`INSERT INTO resetpswrd (clientID, resetcode, token) 
	// values ($1, $2, $3)`)
	// row := r.db.QueryRow(query, userID, email.Code, email.Token)
	// fmt.Println(row)



	return nil
}
