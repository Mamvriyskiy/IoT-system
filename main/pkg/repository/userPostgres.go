package repository

import (
	"fmt"

	"github.com/Mamvriyskiy/DBCourse/main/logger"
	pkg "github.com/Mamvriyskiy/DBCourse/main/pkg"
	"github.com/jmoiron/sqlx"
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

func (r *UserPostgres) GetUser(login, password string) (pkg.User, error) {
	var user pkg.User
	query := fmt.Sprintf("SELECT clientid from %s where login = $1 and password = $2", "client")
	err := r.db.Get(&user, query, login, password)

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
	_, err := r.db.Exec(query, password, token)

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

	query = fmt.Sprintf(`INSERT INTO resetpswrd (clientID, resetcode, token) 
	values ($1, $2, $3)`)
	row := r.db.QueryRow(query, userID, email.Code, email.Token)
	fmt.Println(row)

	return err
}
