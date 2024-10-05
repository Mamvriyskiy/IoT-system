package introtests

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/Mamvriyskiy/database_course/main/pkg/repository"
	"github.com/Mamvriyskiy/database_course/main/pkg/service"
	"github.com/Mamvriyskiy/database_course/main/tests/factory"
	method "github.com/Mamvriyskiy/database_course/main/tests/method"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

const (
	salt = "hfdjmaxckdk20"
)

func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return hex.EncodeToString(hash.Sum([]byte(salt)))
}

func (s *MyIntroTestsSuite) TestCreateClientIntro(t provider.T) {
	tests := []struct {
		nameTest string
		user     factory.ObjectSystem
	}{
		{
			nameTest: "Test1",
			user:     factory.New("user", ""),
		},
		{
			nameTest: "Test2",
			user:     factory.New("user", ""),
		},
		{
			nameTest: "Test3",
			user:     factory.New("user", ""),
		},
	}

	repos := repository.NewRepository(connDB)
	services := service.NewServicesPsql(repos)

	for _, test := range tests {
		t.Run(test.nameTest, func(t provider.T) {
			newUser := test.user.(*method.TestUser)

			resultID, err := services.IUser.CreateUser(newUser.User)
			t.Require().NoError(err)

			newUser.User.ID = resultID
			query := `SELECT password, login, email FROM client WHERE clientid = $1`
			row := connDB.QueryRow(query, resultID)

			retrievedUser := pkg.User{
				ID: resultID,
			}
			newUser.Password = generatePasswordHash(newUser.Password)

			err = row.Scan(&retrievedUser.Password, &retrievedUser.Username, &retrievedUser.Email)
			t.Require().NoError(err)
			t.Assert().Equal(newUser.User, retrievedUser)
		})
	}
}

func (s *MyIntroTestsSuite) TestGetClientIntro(t provider.T) {
	tests := []struct {
		nameTest    string
		user        factory.ObjectSystem
		SearchEmail string
		ResultError error
	}{
		{
			nameTest:    "Test1",
			user:        factory.New("user", "email4"),
			SearchEmail: "email4",
			ResultError: nil,
		},
		{
			nameTest:    "Test2",
			user:        factory.New("user", "email5"),
			SearchEmail: "email5",
			ResultError: nil,
		},
		{
			nameTest:    "Test3",
			user:        factory.New("user", "email6"),
			SearchEmail: "email6",
			ResultError: nil,
		},
		{
			nameTest:    "Test4",
			user:        factory.New("user", "email7"),
			SearchEmail: "email8",
			ResultError: errors.New("sql: no rows in result set"),
		},
	}

	repos := repository.NewRepository(connDB)
	services := service.NewServicesPsql(repos)

	for _, test := range tests {
		t.Run(test.nameTest, func(t provider.T) {
			newUser := test.user.(*method.TestUser)

			clientID, err := newUser.InsertObject(connDB)

			t.Require().NoError(err)
			newUser.Email = test.SearchEmail

			resultUserID, err := services.IUser.CheckUser(newUser.User)
			if err != nil {
				t.Require().Equal(err, test.ResultError)
			} else {
				t.Assert().Equal(clientID, resultUserID)
			}
		})
	}
}

func (s *MyIntroTestsSuite) TestChangePasswordIntro(t provider.T) {
	tests := []struct {
		nameTest    string
		user        factory.ObjectSystem
		NewPassword string
		Token       string
		ResetCode   string
		SearchToken string
		ResultError error
	}{
		{
			nameTest:    "Test1",
			user:        factory.New("user", ""),
			NewPassword: "pswrd2",
			Token:       "kakfksdkfmksdv",
			SearchToken: "kakfksdkfmksdv",
			ResetCode:   "123456",
		},
		{
			nameTest:    "Test2",
			user:        factory.New("user", ""),
			NewPassword: "pswrd3",
			Token:       "meqwefdf",
			SearchToken: "meqwefdf",
			ResetCode:   "asdfgd",
		},
		{
			nameTest:    "Test4",
			user:        factory.New("user", ""),
			NewPassword: "pswrd5",
			Token:       "naskfnjsndjg",
			SearchToken: "alkmckmjvnjnfh",
			ResetCode:   "fldkfdkfj",
			ResultError: errors.New("No updated rows"),
		},
	}

	repos := repository.NewRepository(connDB)
	services := service.NewServicesPsql(repos)

	for _, test := range tests {
		t.Run(test.nameTest, func(t provider.T) {
			newUser := test.user.(*method.TestUser)

			clientID, err := newUser.InsertObject(connDB)
			t.Require().NoError(err)

			query1 := `INSERT INTO resetpswrd (clientID, resetCode, token) values ($1, $2, $3);`
			_ = connDB.QueryRow(query1, clientID, test.ResetCode, test.Token)

			err = services.IUser.ChangePassword(test.NewPassword, test.SearchToken)
			if err != nil {
				t.Require().Equal(err, test.ResultError)
				return
			}

			test.NewPassword = generatePasswordHash(test.NewPassword)

			query := `SELECT password FROM client WHERE clientid = $1`
			row := connDB.QueryRow(query, clientID)

			var pswrd string
			err = row.Scan(&pswrd)
			t.Require().NoError(err)

			t.Assert().Equal(test.NewPassword, pswrd)
		})
	}
}

func (s *MyIntroTestsSuite) TestCheckCodeIntro(t provider.T) {
	tests := []struct {
		nameTest    string
		ClientID    int
		Token       string
		SearchToken string
		Code        string
		SearchCode  string
		ResultError error
	}{
		{
			nameTest:    "Test1",
			ClientID:    111,
			Token:       "mcakmsfkdfkdf",
			SearchToken: "mcakmsfkdfkdf",
			Code:        "avbdkk",
			SearchCode:  "avbdkk",
			ResultError: nil,
		},
		{
			nameTest:    "Test3",
			ClientID:    131,
			Token:       "dasnfjajkcddj",
			SearchToken: "dasnfjajkcddj",
			Code:        "czxc",
			SearchCode:  "kfdsfkksd",
			ResultError: errors.New("Negative code"),
		},
		{
			nameTest:    "Test4",
			ClientID:    141,
			Token:       "ghfdffcbhjbhjbbf",
			SearchToken: "asklfkdmjf",
			Code:        "3214dsaf",
			SearchCode:  "3214dsaf",
			ResultError: errors.New("Negative code"),
		},
	}

	repos := repository.NewRepository(connDB)
	services := service.NewServicesPsql(repos)

	for _, test := range tests {
		t.Run(test.nameTest, func(t provider.T) {
			query := `INSERT INTO resetpswrd (clientid, token, resetcode) values ($1, $2, $3);`
			_ = connDB.QueryRow(query, test.ClientID, test.Token, test.Code)

			err := services.IUser.CheckCode(test.SearchCode, test.SearchToken)

			if err != nil {
				t.Assert().Equal(err, test.ResultError)
			}
		})
	}
}
