package testsdatabase

import (
	"errors"
	"github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/Mamvriyskiy/database_course/main/pkg/repository"
	"github.com/stretchr/testify/assert"
	"testing"
	"strconv"
)

func TestCreateClient(t *testing.T) {
	tests := []struct {
		nameTest string
		user     pkg.User
	}{
		{
			nameTest: "Test1",
			user: pkg.User{
				Password: "pswrd1",
				Username: "user1",
				Email:    "email1",
				ID:       1,
			},
		},
		{
			nameTest: "Test2",
			user: pkg.User{
				Password: "pswrd2",
				Username: "user2",
				Email:    "email2",
				ID:       2,
			},
		},
		{
			nameTest: "Test3",
			user: pkg.User{
				Password: "pswrd3",
				Username: "user3",
				Email:    "email3",
				ID:       3,
			},
		},
	}

	repos := repository.NewRepository(connDB)

	for _, test := range tests {
		t.Run(test.nameTest, func(t *testing.T) {
			resultID, err := repos.IUserRepo.CreateUser(test.user)
			if err != nil {
				assert.NoError(t, err, "")
			}

			query := `SELECT password, login, email FROM client WHERE clientid = $1`
			row := connDB.QueryRow(query, resultID)

			retrievedUser := pkg.User{
				ID: resultID,
			}

			err = row.Scan(&retrievedUser.Password, &retrievedUser.Username, &retrievedUser.Email)
			if err != nil {
				assert.NoError(t, err, "")
			}

			assert.Equal(t, test.user, retrievedUser, "The passwords should be the same.")
		})
	}
}

func TestGetClient(t *testing.T) {
	tests := []struct {
		nameTest    string
		user        pkg.User
		SearchEmail string
		ResultError error
	}{
		{
			nameTest: "Test1",
			user: pkg.User{
				Password: "pswrd4",
				Username: "user4",
				Email:    "email4",
			},
			SearchEmail: "email4",
			ResultError: nil,
		},
		{
			nameTest: "Test2",
			user: pkg.User{
				Password: "pswrd5",
				Username: "user5",
				Email:    "email5",
			},
			SearchEmail: "email5",
			ResultError: nil,
		},
		{
			nameTest: "Test3",
			user: pkg.User{
				Password: "pswrd6",
				Username: "user6",
				Email:    "email6",
			},
			SearchEmail: "email6",
			ResultError: nil,
		},
		{
			nameTest: "Test4",
			user: pkg.User{
				Password: "pswrd6",
				Username: "user6",
				Email:    "email7",
			},
			SearchEmail: "email8",
			ResultError: errors.New("sql: no rows in result set"),
		},
	}

	repos := repository.NewRepository(connDB)

	for _, test := range tests {
		t.Run(test.nameTest, func(t *testing.T) {
			query := `INSERT INTO client (password, login, email) values ($1, $2, $3) RETURNING clientid`
			row := connDB.QueryRow(query, test.user.Password, test.user.Username, test.user.Email)

			var clientID int
			if err := row.Scan(&clientID); err != nil {
				assert.NoError(t, err, "")
			}

			resultUser, err := repos.IUserRepo.GetUser(test.SearchEmail, test.user.Password)
			assert.Equal(t, err, test.ResultError)
			if test.ResultError == nil {
				assert.Equal(t, clientID, resultUser.ID, "The clientIDs should be the same.")
			}
		})
	}
}

// ChangePassword(password, token string)
func TestChangePassword(t *testing.T) {
	tests := []struct {
		nameTest       string
		user           pkg.User
		NewPassword    string
		Token          string
		ResetCode      string
		SearchToken    string
		SearchPassword string
	}{
		{
			nameTest: "Test1",
			user: pkg.User{
				Password: "pswrd4",
				Username: "user4",
				Email:    "email10",
			},
			NewPassword:    "pswrd2",
			SearchPassword: "pswrd2",
			Token:          "kakfksdkfmksdv",
			SearchToken:    "kakfksdkfmksdv",
			ResetCode:      "123456",
		},
		{
			nameTest: "Test2",
			user: pkg.User{
				Password: "pswrd5",
				Username: "user5",
				Email:    "email11",
			},
			NewPassword:    "pswrd3",
			SearchPassword: "pswrd3",
			Token:          "meqwefdf",
			SearchToken:    "meqwefdf",
			ResetCode:      "asdfgd",
		},
		{
			nameTest: "Test4",
			user: pkg.User{
				Password: "pswrd6",
				Username: "user6",
				Email:    "email17",
			},
			NewPassword:    "pswrd5",
			SearchPassword: "pswrd6",
			Token:          "naskfnjsndjg",
			SearchToken:    "alkmckmjvnjnfh",
			ResetCode:      "fldkfdkfj",
		},
	}

	repos := repository.NewRepository(connDB)

	for _, test := range tests {
		t.Run(test.nameTest, func(t *testing.T) {
			query := `INSERT INTO client (password, login, email) values ($1, $2, $3) RETURNING clientid;`
			row := connDB.QueryRow(query, test.user.Password, test.user.Username, test.user.Email)

			var clientID int
			if err := row.Scan(&clientID); err != nil {
				assert.NoError(t, err, "")
			}

			query1 := `INSERT INTO resetpswrd (clientID, resetCode, token) values ($1, $2, $3);`
			_ = connDB.QueryRow(query1, clientID, test.ResetCode, test.Token)

			_ = repos.IUserRepo.ChangePassword(test.NewPassword, test.SearchToken)

			query = `SELECT password FROM client WHERE clientid = $1`
			row = connDB.QueryRow(query, clientID)

			var pswrd string
			err := row.Scan(&pswrd)
			if err != nil {
				assert.NoError(t, err, "")
			}

			assert.Equal(t, test.SearchPassword, pswrd)
		})
	}
}

// GetCode(token string) (string, error)
func TestGetCode(t *testing.T) {
	tests := []struct {
		nameTest    string
		ClientID    int
		Token       string
		SearchToken string
		Code        string
		ResultError error
	}{
		{
			nameTest:    "Test1",
			ClientID:    111,
			Token:       "mcakmsfkdfkdf",
			SearchToken: "mcakmsfkdfkdf",
			Code:        "avbdkk",
			ResultError: nil,
		},
		{
			nameTest:    "Test3",
			ClientID:    131,
			Token:       "dasnfjajkcddj",
			SearchToken: "dasnfjajkcddj",
			Code:        "czxc",
			ResultError: nil,
		},
		{
			nameTest:    "Test4",
			ClientID:    141,
			Token:       "ghfdffcbhjbhjbbf",
			SearchToken: "asklfkdmjf",
			Code:        "3214dsaf",
			ResultError: errors.New("sql: no rows in result set"),
		},
	}

	repos := repository.NewRepository(connDB)

	for _, test := range tests {
		t.Run(test.nameTest, func(t *testing.T) {
			query := `INSERT INTO resetpswrd (clientid, token, resetcode) values ($1, $2, $3);`
			_ = connDB.QueryRow(query, test.ClientID, test.Token, test.Code)

			resultCode, err := repos.IUserRepo.GetCode(test.SearchToken)
			assert.Equal(t, err, test.ResultError)
			if test.ResultError == nil {
				assert.Equal(t, resultCode, test.Code, "The codes should be the same.")
			}
		})
	}
}

func TestAddCode(t *testing.T) {
	tests := []struct {
		nameTest    string
		user        pkg.User
		reset       pkg.Email
		ResultError error
	}{
		{
			nameTest: "Test1",
			user: pkg.User{
				Password: "pswrd1",
				Username: "user1",
				Email:    "email121",
				ID:       100,
			},
			reset: pkg.Email{
				Email: "email121",
				Code:  123,
				Token: "fdskgdsg",
			},
			ResultError: nil,
		},
		{
			nameTest: "Test2",
			user: pkg.User{
				Password: "pswrd2",
				Username: "user2",
				Email:    "email221",
				ID:       200,
			},
			reset: pkg.Email{
				Email: "email221",
				Code:  5435,
				Token: "fldasklgfksdg",
			},
			ResultError: nil,
		},
		{
			nameTest: "Test3",
			user: pkg.User{
				Password: "pswrd3",
				Username: "user3",
				Email:    "email321",
				ID:       300,
			},
			reset: pkg.Email{
				Email: "email521",
				Code:  5435,
				Token: "fldasklgfksdg",
			},
			ResultError: errors.New("sql: no rows in result set"),
		},
	}

	repos := repository.NewRepository(connDB)

	for _, test := range tests {
		t.Run(test.nameTest, func(t *testing.T) {
			query := `INSERT INTO client (password, login, email) values ($1, $2, $3) RETURNING clientid`
			row := connDB.QueryRow(query, test.user.Password, test.user.Username, test.user.Email)

			var clientID int
			if err := row.Scan(&clientID); err != nil {
				assert.NoError(t, err, "")
			}

			err := repos.IUserRepo.AddCode(test.reset)
			assert.Equal(t, err, test.ResultError)

			if test.ResultError == nil {
				query = `SELECT resetcode FROM resetpswrd WHERE token = $1`
				row = connDB.QueryRow(query, test.reset.Token)
	
				var code string
				err := row.Scan(&code)
				if err != nil {
					assert.NoError(t, err, "")
				}

				res := strconv.Itoa(test.reset.Code)
				assert.Equal(t, code, res)
			}
		})
	}
}
