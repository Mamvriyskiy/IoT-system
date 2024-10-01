package unittests

import (
	"github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/Mamvriyskiy/database_course/main/pkg/repository"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/Mamvriyskiy/database_course/main/tests/factory"
	method "github.com/Mamvriyskiy/database_course/main/tests/method"
	//"reflect"
	"errors"
	"strconv"
)

func (s *MyFirstSuite) TestCreateClient(t provider.T) {
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
			user: 	  factory.New("user", ""),
		},
		{
			nameTest: "Test3",
			user: 	  factory.New("user", ""),
		},
	}

	repos := repository.NewRepository(connDB)

	for _, test := range tests {
		t.Run(test.nameTest, func(t provider.T) {
			newUser := test.user.(*method.TestUser)
			
			resultID, err := repos.IUserRepo.CreateUser(newUser.User)
			t.Require().NoError(err)

			newUser.User.ID = resultID
			query := `SELECT password, login, email FROM client WHERE clientid = $1`
			row := connDB.QueryRow(query, resultID)

			retrievedUser := pkg.User{
				ID: resultID,
			}

			err = row.Scan(&retrievedUser.Password, &retrievedUser.Username, &retrievedUser.Email)
			t.Require().NoError(err)
			t.Assert().Equal(newUser.User, retrievedUser)
		})
	}
}

func (s *MyFirstSuite) TestGetClient(t provider.T) {
	tests := []struct {
		nameTest    string
		user        factory.ObjectSystem
		SearchEmail string
		ResultError error
	}{
		{
			nameTest: "Test1",
			user: factory.New("user", "email4"),
			SearchEmail: "email4",
			ResultError: nil,
		},
		{
			nameTest: "Test2",
			user: factory.New("user", "email5"),
			SearchEmail: "email5",
			ResultError: nil,
		},
		{
			nameTest: "Test3",
			user: factory.New("user", "email6"),
			SearchEmail: "email6",
			ResultError: nil,
		},
		{
			nameTest: "Test4",
			user: factory.New("user", "email7"),
			SearchEmail: "email8",
			ResultError: errors.New("sql: no rows in result set"),
		},
	}

	repos := repository.NewRepository(connDB)

	for _, test := range tests {
		t.Run(test.nameTest, func(t provider.T) {
			newUser := test.user.(*method.TestUser)

			clientID, err := newUser.InsertObject(connDB)

			t.Require().NoError(err)

			resultUser, err := repos.IUserRepo.GetUser(test.SearchEmail, newUser.Password)
			if err != nil {
				t.Require().Equal(err, test.ResultError)
			} else {
				t.Assert().Equal(clientID, resultUser.ID)
			}
		})
	}
}

func (s *MyFirstSuite) TestChangePassword(t provider.T) {
	tests := []struct {
		nameTest       string
		user           factory.ObjectSystem
		NewPassword    string
		Token          string
		ResetCode      string
		SearchToken    string
		ResultError    error
	}{
		{
			nameTest: "Test1",
			user: factory.New("user", ""),
			NewPassword:    "pswrd2",
			Token:          "kakfksdkfmksdv",
			SearchToken:    "kakfksdkfmksdv",
			ResetCode:      "123456",
		},
		{
			nameTest: "Test2",
			user: factory.New("user", ""),
			NewPassword:    "pswrd3",
			Token:          "meqwefdf",
			SearchToken:    "meqwefdf",
			ResetCode:      "asdfgd",
		},
		{
			nameTest: "Test4",
			user: factory.New("user", ""),
			NewPassword:    "pswrd5",
			Token:          "naskfnjsndjg",
			SearchToken:    "alkmckmjvnjnfh",
			ResetCode:      "fldkfdkfj",
			ResultError: errors.New("No updated rows"),
		},
	}

	repos := repository.NewRepository(connDB)

	for _, test := range tests {
		t.Run(test.nameTest, func(t provider.T) {
			newUser := test.user.(*method.TestUser)

			clientID, err := newUser.InsertObject(connDB)
			t.Require().NoError(err)

			query1 := `INSERT INTO resetpswrd (clientID, resetCode, token) values ($1, $2, $3);`
			_ = connDB.QueryRow(query1, clientID, test.ResetCode, test.Token)

			err = repos.IUserRepo.ChangePassword(test.NewPassword, test.SearchToken)
			if err != nil {
				t.Require().Equal(err, test.ResultError)
				return
			}

			query := `SELECT password FROM client WHERE clientid = $1`
			row := connDB.QueryRow(query, clientID)

			var pswrd string
			err = row.Scan(&pswrd)
			t.Require().NoError(err)

			t.Assert().Equal(test.NewPassword, pswrd)
		})
	}
}

func (s *MyFirstSuite) TestGetCode(t provider.T) {
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
		t.Run(test.nameTest, func(t provider.T) {
			query := `INSERT INTO resetpswrd (clientid, token, resetcode) values ($1, $2, $3);`
			_ = connDB.QueryRow(query, test.ClientID, test.Token, test.Code)

			resultCode, err := repos.IUserRepo.GetCode(test.SearchToken)
			if err != nil {
				t.Assert().Equal(err, test.ResultError)
			} else {
				t.Assert().Equal(resultCode, test.Code)
			}	
		})
	}
}

func (s *MyFirstSuite) TestAddCode(t provider.T) {
	tests := []struct {
		nameTest    string
		user        factory.ObjectSystem
		reset       factory.ObjectSystem
		ResultError error
	}{
		{
			nameTest: "Test1",
			user: factory.New("user", "email121"),
			reset: factory.New("email", "email121"),
			ResultError: nil,
		},
		{
			nameTest: "Test2",
			user: factory.New("user", "email221"),
			reset: factory.New("email", "email221"),
			ResultError: nil,
		},
		{
			nameTest: "Test3",
			user: factory.New("user", "email321"),
			reset: factory.New("email", "email421"),
			ResultError: errors.New("sql: no rows in result set"),
		},
	}

	repos := repository.NewRepository(connDB)

	for _, test := range tests {
		t.Run(test.nameTest, func(t provider.T) {
			newUser := test.user.(*method.TestUser)
			newEmail := test.reset.(*method.TestEmail)

			_, err := newUser.InsertObject(connDB)

			t.Require().NoError(err)

			err = repos.IUserRepo.AddCode(newEmail.Email)
			if err != nil {
				t.Require().Equal(err, test.ResultError)
				return
			}

			query := `SELECT resetcode FROM resetpswrd WHERE token = $1`
			row := connDB.QueryRow(query, newEmail.Token)

			var code string
			err = row.Scan(&code)
			t.Require().NoError(err)

			res := strconv.Itoa(newEmail.Code)
			t.Assert().Equal(code, res)
		})
	}
}
