package unittests

import (
	"fmt"
	"github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/Mamvriyskiy/database_course/main/pkg/repository"
	"github.com/Mamvriyskiy/database_course/main/tests/factory"
	method "github.com/Mamvriyskiy/database_course/main/tests/method"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

func (s *MyUnitTestsSuite) TestAddClient(t provider.T) {
	tests := []struct {
		nameTest    string
		user        factory.ObjectSystem
		owner       factory.ObjectSystem
		home        factory.ObjectSystem
		accessUser  factory.ObjectSystem
		accessOwner factory.ObjectSystem
	}{
		{
			nameTest:    "Test1",
			user:        factory.New("user", "", "DB"),
			owner:       factory.New("user", "", "DB"),
			home:        factory.New("home", "", "DB"),
			accessUser:  factory.New("access", "", "DB"),
			accessOwner: factory.New("access", "", "DB"),
		},
		{
			nameTest:    "Test1",
			user:        factory.New("user", "", "DB"),
			owner:       factory.New("user", "", "DB"),
			home:        factory.New("home", "", "DB"),
			accessUser:  factory.New("access", "", "DB"),
			accessOwner: factory.New("access", "", "DB"),
		},
		{
			nameTest:    "Test1",
			user:        factory.New("user", "", "DB"),
			owner:       factory.New("user", "", "DB"),
			home:        factory.New("home", "", "DB"),
			accessUser:  factory.New("access", "", "DB"),
			accessOwner: factory.New("access", "", "DB"),
		},
	}

	repos := repository.NewRepository(connDB)

	for _, test := range tests {
		t.Run(test.nameTest, func(t provider.T) {
			newUser := test.user.(*method.TestUser)
			newHome := test.home.(*method.TestHome)
			newOwner := test.user.(*method.TestUser)
			newAccessOwner := test.accessOwner.(*method.TestAccess)
			newAccessUser := test.accessUser.(*method.TestAccess)

			homeID, err := newHome.InsertObject(connDB)
			t.Require().NoError(err)

			userID, err := newUser.InsertObject(connDB)
			t.Require().NoError(err)

			ownerID, err := newOwner.InsertObject(connDB)
			t.Require().NoError(err)

			newAccessOwner.ClientID = ownerID
			newAccessOwner.HomeID = homeID
			_, err = newAccessOwner.InsertObject(connDB)

			newAccessUser.Email = newUser.Email
			// newAccessUser.Home = newHome.Name

			// accessService := pkg.AccessService{
			// 	Access:
			// }

			accessID, err := repos.IAccessHomeRepo.AddUser(homeID, newAccessUser.AccessService)
			t.Require().NoError(err)

			var clientID string
			query := `SELECT clientID FROM access WHERE accessID = $1`
			row := connDB.QueryRow(query, accessID)

			err = row.Scan(&clientID)
			t.Require().NoError(err)
			t.Assert().Equal(userID, clientID)
		})
	}
}

func (s *MyUnitTestsSuite) TestUpdateLevel(t provider.T) {
	tests := []struct {
		nameTest    string
		user        factory.ObjectSystem
		home        factory.ObjectSystem
		accessUser  factory.ObjectSystem
		updateLevel int
	}{
		{
			nameTest:    "Test1",
			user:        factory.New("user", "", "DB"),
			home:        factory.New("home", "", "DB"),
			accessUser:  factory.New("access", "", "DB"),
			updateLevel: 10,
		},
		{
			nameTest:    "Test1",
			user:        factory.New("user", "", "DB"),
			home:        factory.New("home", "", "DB"),
			accessUser:  factory.New("access", "", "DB"),
			updateLevel: 10,
		},
		{
			nameTest:    "Test1",
			user:        factory.New("user", "", "DB"),
			home:        factory.New("home", "", "DB"),
			accessUser:  factory.New("access", "", "DB"),
			updateLevel: 10,
		},
	}

	repos := repository.NewRepository(connDB)

	for _, test := range tests {
		t.Run(test.nameTest, func(t provider.T) {
			newUser := test.user.(*method.TestUser)
			newHome := test.home.(*method.TestHome)
			newAccessUser := test.accessUser.(*method.TestAccess)

			homeID, err := newHome.InsertObject(connDB)
			t.Require().NoError(err)

			userID, err := newUser.InsertObject(connDB)
			t.Require().NoError(err)

			newAccessUser.Email = newUser.Email
			//newAccessUser.Home = newHome.Name
			newAccessUser.AccessLevel = test.updateLevel
			newAccessUser.ClientID = userID
			newAccessUser.HomeID = homeID

			_, err = newAccessUser.InsertObject(connDB)

			err = repos.IAccessHomeRepo.UpdateLevel(userID, newAccessUser.AccessService)
			t.Require().NoError(err)

			var accessLevel int
			query := `select accesslevel from access
				WHERE clientid = $1
					and homeid = $2;`
			row := connDB.QueryRow(query, userID, homeID)

			err = row.Scan(&accessLevel)

			t.Require().NoError(err)
			t.Assert().Equal(test.updateLevel, accessLevel)
		})
	}
}

func (s *MyUnitTestsSuite) TestUpdateStatus(t provider.T) {
	tests := []struct {
		nameTest   string
		user       factory.ObjectSystem
		accessUser factory.ObjectSystem
		accessHome pkg.AccessService
		home       factory.ObjectSystem
	}{
		{
			nameTest:   "Test1",
			user:       factory.New("user", "", "DB"),
			home:       factory.New("home", "", "DB"),
			accessUser: factory.New("access", "", "DB"),
			accessHome: pkg.AccessService{
				Access: pkg.Access{
					AccessStatus: "blocked",
				},
			},
		},
		{
			nameTest:   "Test1",
			user:       factory.New("user", "", "DB"),
			home:       factory.New("home", "", "DB"),
			accessUser: factory.New("access", "", "DB"),
			accessHome: pkg.AccessService{
				Access: pkg.Access{
					AccessStatus: "blocked",
				},
			},
		},
		{
			nameTest:   "Test1",
			user:       factory.New("user", "", "DB"),
			home:       factory.New("home", "", "DB"),
			accessUser: factory.New("access", "", "DB"),
			accessHome: pkg.AccessService{
				Access: pkg.Access{
					AccessStatus: "blocked",
				},
			},
		},
	}

	repos := repository.NewRepository(connDB)

	for _, test := range tests {
		t.Run(test.nameTest, func(t provider.T) {
			newUser := test.user.(*method.TestUser)
			newAccessUser := test.accessUser.(*method.TestAccess)

			userID, err := newUser.InsertObject(connDB)
			t.Require().NoError(err)

			newHome := test.home.(*method.TestHome)
			homeID, err := newHome.InsertObject(connDB)
			fmt.Println("=====", homeID, err, "=====")
			t.Require().NoError(err)

			newAccessUser.Email = newUser.Email
			newAccessUser.ClientID = userID
			newAccessUser.HomeID = homeID

			accessID, err := newAccessUser.InsertObject(connDB)
			t.Require().NoError(err)

			t.Log("userID: ", userID, "accessID: ", accessID)

			err = repos.IAccessHomeRepo.UpdateStatus(accessID, test.accessHome)
			t.Require().NoError(err)

			var accessstatus string
			query := `select accessstatus from access
					WHERE accessID = $1;`
			row := connDB.QueryRow(query, accessID)

			err = row.Scan(&accessstatus)

			t.Require().NoError(err)
			t.Assert().Equal(test.accessHome.AccessStatus, accessstatus)
		})
	}
}

func (s *MyUnitTestsSuite) TestGetListUserHome(t provider.T) {
	tests := []struct {
		nameTest string
		lenList  int
		home     factory.ObjectSystem
	}{
		{
			nameTest: "Test1",
			lenList:  1,
			home:     factory.New("home", "", "DB"),
		},
		{
			nameTest: "Test2",
			lenList:  10,
			home:     factory.New("home", "", "DB"),
		},
	}

	repos := repository.NewRepository(connDB)

	for _, test := range tests {
		t.Run(test.nameTest, func(t provider.T) {
			newHome := test.home.(*method.TestHome)

			homeID, err := newHome.InsertObject(connDB)

			t.Require().NoError(err)

			listHome := make([]pkg.ClientHome, test.lenList)
			for i := 0; i < test.lenList; i++ {
				newUser := factory.New("user", "")
				user := newUser.(*method.TestUser)

				newAccess := factory.New("access", "")
				access := newAccess.(*method.TestAccess)

				userID, err := user.InsertObject(connDB)
				t.Require().NoError(err)

				access.ClientID = userID
				access.HomeID = homeID
				_, err = access.InsertObject(connDB)
				t.Require().NoError(err)

				listHome[i].Home = newHome.Name
				listHome[i].Username = user.Username
				listHome[i].Email = user.Email
				listHome[i].AccessLevel = access.AccessLevel
				listHome[i].AccessStatus = "active"
			}

			resultListHome, err := repos.IAccessHomeRepo.GetListUserHome(homeID)

			t.Require().NoError(err)

			t.Assert().Equal(listHome, resultListHome)
		})
	}
}
