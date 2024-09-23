package testsdatabase

import (
	"github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/Mamvriyskiy/database_course/main/pkg/repository"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/Mamvriyskiy/database_course/main/testsdatabase/factory"
	method "github.com/Mamvriyskiy/database_course/main/testsdatabase/method"
	//"fmt"
	//"reflect"
	"errors"
	// "strconv"
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

// func (s *MyFirstSuite) TestChangePassword(t provider.T) {
// 	tests := []struct {
// 		nameTest       string
// 		user           pkg.User
// 		NewPassword    string
// 		Token          string
// 		ResetCode      string
// 		SearchToken    string
// 		SearchPassword string
// 	}{
// 		{
// 			nameTest: "Test1",
// 			user: pkg.User{
// 				Password: "pswrd4",
// 				Username: "user4",
// 				Email:    "email10",
// 			},
// 			NewPassword:    "pswrd2",
// 			SearchPassword: "pswrd2",
// 			Token:          "kakfksdkfmksdv",
// 			SearchToken:    "kakfksdkfmksdv",
// 			ResetCode:      "123456",
// 		},
// 		{
// 			nameTest: "Test2",
// 			user: pkg.User{
// 				Password: "pswrd5",
// 				Username: "user5",
// 				Email:    "email11",
// 			},
// 			NewPassword:    "pswrd3",
// 			SearchPassword: "pswrd3",
// 			Token:          "meqwefdf",
// 			SearchToken:    "meqwefdf",
// 			ResetCode:      "asdfgd",
// 		},
// 		{
// 			nameTest: "Test4",
// 			user: pkg.User{
// 				Password: "pswrd6",
// 				Username: "user6",
// 				Email:    "email17",
// 			},
// 			NewPassword:    "pswrd5",
// 			SearchPassword: "pswrd6",
// 			Token:          "naskfnjsndjg",
// 			SearchToken:    "alkmckmjvnjnfh",
// 			ResetCode:      "fldkfdkfj",
// 		},
// 	}

// 	repos := repository.NewRepository(connDB)

// 	for _, test := range tests {
// 		t.Run(test.nameTest, func(t provider.T) {
// 			query := `INSERT INTO client (password, login, email) values ($1, $2, $3) RETURNING clientid;`
// 			row := connDB.QueryRow(query, test.user.Password, test.user.Username, test.user.Email)

// 			var clientID int
// 			err := row.Scan(&clientID)
// 			t.Require().NoError(err)

// 			query1 := `INSERT INTO resetpswrd (clientID, resetCode, token) values ($1, $2, $3);`
// 			_ = connDB.QueryRow(query1, clientID, test.ResetCode, test.Token)

// 			_ = repos.IUserRepo.ChangePassword(test.NewPassword, test.SearchToken)

// 			query = `SELECT password FROM client WHERE clientid = $1`
// 			row = connDB.QueryRow(query, clientID)

// 			var pswrd string
// 			err = row.Scan(&pswrd)
// 			t.Require().NoError(err)

// 			t.Assert().Equal(test.SearchPassword, pswrd)
// 		})
// 	}
// }

// func (s *MyFirstSuite) TestGetCode(t provider.T) {
// 	tests := []struct {
// 		nameTest    string
// 		ClientID    int
// 		Token       string
// 		SearchToken string
// 		Code        string
// 		ResultError error
// 	}{
// 		{
// 			nameTest:    "Test1",
// 			ClientID:    111,
// 			Token:       "mcakmsfkdfkdf",
// 			SearchToken: "mcakmsfkdfkdf",
// 			Code:        "avbdkk",
// 			ResultError: nil,
// 		},
// 		{
// 			nameTest:    "Test3",
// 			ClientID:    131,
// 			Token:       "dasnfjajkcddj",
// 			SearchToken: "dasnfjajkcddj",
// 			Code:        "czxc",
// 			ResultError: nil,
// 		},
// 		{
// 			nameTest:    "Test4",
// 			ClientID:    141,
// 			Token:       "ghfdffcbhjbhjbbf",
// 			SearchToken: "asklfkdmjf",
// 			Code:        "3214dsaf",
// 			ResultError: errors.New("sql: no rows in result set"),
// 		},
// 	}

// 	repos := repository.NewRepository(connDB)

// 	for _, test := range tests {
// 		t.Run(test.nameTest, func(t provider.T) {
// 			query := `INSERT INTO resetpswrd (clientid, token, resetcode) values ($1, $2, $3);`
// 			_ = connDB.QueryRow(query, test.ClientID, test.Token, test.Code)

// 			resultCode, err := repos.IUserRepo.GetCode(test.SearchToken)
// 			if err != nil {
// 				t.Assert().Equal(err, test.ResultError)
// 			} else {
// 				t.Assert().Equal(resultCode, test.Code)
// 			}	
// 		})
// 	}
// }

// func (s *MyFirstSuite) TestAddCode(t provider.T) {
// 	tests := []struct {
// 		nameTest    string
// 		user        pkg.User
// 		reset       pkg.Email
// 		ResultError error
// 	}{
// 		{
// 			nameTest: "Test1",
// 			user: pkg.User{
// 				Password: "pswrd1",
// 				Username: "user1",
// 				Email:    "email121",
// 				ID:       100,
// 			},
// 			reset: pkg.Email{
// 				Email: "email121",
// 				Code:  123,
// 				Token: "fdskgdsg",
// 			},
// 			ResultError: nil,
// 		},
// 		{
// 			nameTest: "Test2",
// 			user: pkg.User{
// 				Password: "pswrd2",
// 				Username: "user2",
// 				Email:    "email221",
// 				ID:       200,
// 			},
// 			reset: pkg.Email{
// 				Email: "email221",
// 				Code:  5435,
// 				Token: "fldasklgfksdg",
// 			},
// 			ResultError: nil,
// 		},
// 		{
// 			nameTest: "Test3",
// 			user: pkg.User{
// 				Password: "pswrd3",
// 				Username: "user3",
// 				Email:    "email321",
// 				ID:       300,
// 			},
// 			reset: pkg.Email{
// 				Email: "email521",
// 				Code:  5435,
// 				Token: "fldasklgfksdg",
// 			},
// 			ResultError: errors.New("sql: no rows in result set"),
// 		},
// 	}

// 	repos := repository.NewRepository(connDB)

// 	for _, test := range tests {
// 		t.Run(test.nameTest, func(t provider.T) {
// 			query := `INSERT INTO client (password, login, email) values ($1, $2, $3) RETURNING clientid`
// 			row := connDB.QueryRow(query, test.user.Password, test.user.Username, test.user.Email)

// 			var clientID int
// 			err := row.Scan(&clientID)
// 			t.Require().NoError(err)

// 			err = repos.IUserRepo.AddCode(test.reset)
// 			t.Require().Equal(err, test.ResultError)

// 			query = `SELECT resetcode FROM resetpswrd WHERE token = $1`
// 			row = connDB.QueryRow(query, test.reset.Token)

// 			var code string
// 			err = row.Scan(&code)
// 			t.Require().NoError(err)

// 			res := strconv.Itoa(test.reset.Code)
// 			t.Assert().Equal(code, res)
// 		})
// 	}
// }
