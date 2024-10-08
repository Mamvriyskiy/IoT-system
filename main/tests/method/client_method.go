package method 

import (
	"crypto/rand"
	"math/big"
	"github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/jmoiron/sqlx"
	"github.com/google/uuid"
)

const (
	charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	lengthPassword = 12
	lengthUserName = 6
	lengtDomainLeft= 5
	lengtDomainRight = 3
	lengthEmail = 10 
)

type TestUser struct {
	pkg.User
}

func NewUser(email string) *TestUser {
	var b TestUser

	return b.BuilderUser(email)
}

func (b *TestUser) BuilderUser(email string) *TestUser {
	b.generatePassword()
	b.generateUserName()
	if email == "" {
		b.generateEmail()
	} else {
		b.Email = email
	}

	return b
}

func (b *TestUser) generatePassword() {
	password := make([]byte, lengthPassword)
	for i := range password {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		password[i] = charset[n.Int64()]
	}

	b.Password = string(password)
}

func (b *TestUser) generateUserName() {
	userName := make([]byte, lengthUserName)
	for i := range userName {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		userName[i] = charset[n.Int64()]
	}

	b.Username = string(userName)
}

func (b *TestUser) generateEmail() {
	b.Email = createEmail(lengthEmail, lengtDomainLeft, lengtDomainRight)
}

func createEmail(a, b, c int) string {
	email := make([]byte, a + b + c + 2)
	i := 0
	for j := 0; j < a; j++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		email[i] = charset[n.Int64()]
		i++
	}

	email[i] = '@'
	i++

	for j := 0; j < b; j++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		email[i] = charset[n.Int64()]
		i++
	}

	email[i] = '.'
	i++

	for j := 0; j < c; j++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		email[i] = charset[n.Int64()]
		i++
	}

	return string(email)
}

func (tu TestUser) InsertObject(connDB *sqlx.DB) (string, error) {
	clientid := uuid.New()
	query := `INSERT INTO client (password, login, email, clientid) values ($1, $2, $3, $4) RETURNING clientid`
	row := connDB.QueryRow(query, tu.Password, tu.Username, tu.Email, clientid)

	var clientID string
	err := row.Scan(&clientID)
	
	return clientID, err
}
