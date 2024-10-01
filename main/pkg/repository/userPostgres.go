package repository

import (
	"fmt"

	"github.com/Mamvriyskiy/database_course/main/logger"
	pkg "github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/jmoiron/sqlx"
	"errors"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) CreateUser(user pkg.User) (int, error) {
	var id int
	query := fmt.Sprintf(`INSERT INTO %s (password, login, email) 
		values ($1, $2, $3) RETURNING clientid`, "client")
	row := r.db.QueryRow(query, user.Password, user.Username, user.Email)
	if err := row.Scan(&id); err != nil {
		logger.Log("Error", "Scan", "Error insert into client:", err, id)
		return 0, err
	}

	return id, nil
}

func (r *UserPostgres) GetUserByEmail(email string) (int, error) {
	var count int
	query := fmt.Sprintf("SELECT count(clientid) from %s where email = $1", "client")
	err := r.db.Get(&count, query, email)
	fmt.Println(count)
	
	return count, err
}

func (r *UserPostgres) GetUser(email, password string) (pkg.User, error) {
	var user pkg.User
	query := fmt.Sprintf("SELECT clientid from %s where email = $1 and password = $2", "client")
	err := r.db.Get(&user, query, email, password)

	return user, err
}

func (r *UserPostgres) ChangePassword(password, token string) (error) {
	query := `
	UPDATE client
	SET password = $1
	WHERE clientid = (
		SELECT clientid FROM resetpswrd rp
		WHERE rp.token = $2
	);`
	result, err := r.db.Exec(query, password, token)
	if err != nil {
		//fmt.Println("Ошибка выполнения запроса:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		//fmt.Println("Ошибка получения количества обновлённых строк:", err)
		return err
	}

	if rowsAffected == 0 {
		return errors.New("No updated rows")
	}

	return err
} 

func (r *UserPostgres) GetCode(token string) (string, error) {
	var code string
	query := fmt.Sprintf("select resetcode from resetPswrd where token = $1")
	err := r.db.Get(&code, query, token)

	return code, err
}

func (r *UserPostgres) AddCode(email pkg.Email) error {
	var userID int
	query := fmt.Sprintf("select clientID from client where email = $1")
	err := r.db.Get(&userID, query, email.Email)
	if err != nil {
		return err
	}

	query = fmt.Sprintf(`INSERT INTO resetpswrd (clientID, resetcode, token) 
	values ($1, $2, $3)`)
	_ = r.db.QueryRow(query, userID, email.Code, email.Token)

	return err
}
