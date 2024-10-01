package unittests

import (
	"github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/Mamvriyskiy/database_course/main/pkg/repository"
	"github.com/Mamvriyskiy/database_course/main/tests/factory"
	method "github.com/Mamvriyskiy/database_course/main/tests/method"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

func (s *MyFirstSuite) TestAddClient(t provider.T) {
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
			user:        factory.New("user", ""),
			owner:       factory.New("user", ""),
			home:        factory.New("home", ""),
			accessUser:  factory.New("access", ""),
			accessOwner: factory.New("access", ""),
		},
		{
			nameTest:    "Test1",
			user:        factory.New("user", ""),
			owner:       factory.New("user", ""),
			home:        factory.New("home", ""),
			accessUser:  factory.New("access", ""),
			accessOwner: factory.New("access", ""),
		},
		{
			nameTest:    "Test1",
			user:        factory.New("user", ""),
			owner:       factory.New("user", ""),
			home:        factory.New("home", ""),
			accessUser:  factory.New("access", ""),
			accessOwner: factory.New("access", ""),
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

			newAccessOwner.ClientID = ownerID
			newAccessOwner.HomeID = homeID
			_, err = newAccessOwner.InsertObject(connDB)

			newAccessUser.Email = newUser.Email
			newAccessUser.Home = newHome.Name

			accessID, err := repos.IAccessHomeRepo.AddUser(ownerID, newAccessUser.Access)
			t.Require().NoError(err)

			var clientID int
			query := `SELECT clientID FROM access WHERE accessID = $1`
			row := connDB.QueryRow(query, accessID)

			err = row.Scan(&clientID)
			t.Require().NoError(err)
			t.Assert().Equal(userID, clientID)
		})
	}
}

func (s *MyFirstSuite) TestUpdateLevel(t provider.T) {
	tests := []struct {
		nameTest    string
		user        factory.ObjectSystem
		home        factory.ObjectSystem
		accessUser  factory.ObjectSystem
		updateLevel int
	}{
		{
			nameTest:    "Test1",
			user:        factory.New("user", ""),
			home:        factory.New("home", ""),
			accessUser:  factory.New("access", ""),
			updateLevel: 10,
		},
		{
			nameTest:    "Test1",
			user:        factory.New("user", ""),
			home:        factory.New("home", ""),
			accessUser:  factory.New("access", ""),
			updateLevel: 10,
		},
		{
			nameTest:    "Test1",
			user:        factory.New("user", ""),
			home:        factory.New("home", ""),
			accessUser:  factory.New("access", ""),
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
			newAccessUser.Home = newHome.Name
			newAccessUser.AccessLevel = test.updateLevel
			newAccessUser.ClientID = userID
			newAccessUser.HomeID = homeID

			_, err = newAccessUser.InsertObject(connDB)

			err = repos.IAccessHomeRepo.UpdateLevel(userID, newAccessUser.Access)
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

func (s *MyFirstSuite) TestUpdateStatus(t provider.T) {
	tests := []struct {
		nameTest     string
		user         factory.ObjectSystem
		accessUser   factory.ObjectSystem
		accessHome   pkg.AccessHome
	}{
		{
			nameTest:     "Test1",
			user:         factory.New("user", ""),
			accessUser:   factory.New("access", ""),
			accessHome: pkg.AccessHome{
				AccessStatus: "blocked",
			},
		},
		{
			nameTest:     "Test1",
			user:         factory.New("user", ""),
			accessUser:   factory.New("access", ""),
			accessHome: pkg.AccessHome{
				AccessStatus: "blocked",
			},
		},
		{
			nameTest:     "Test1",
			user:         factory.New("user", ""),
			accessUser:   factory.New("access", ""),
			accessHome: pkg.AccessHome{
				AccessStatus: "blocked",
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
			
			newAccessUser.Email = newUser.Email
			newAccessUser.ClientID = userID

			_, err = newAccessUser.InsertObject(connDB)

			err = repos.IAccessHomeRepo.UpdateStatus(userID, test.accessHome)
			t.Require().NoError(err)

			var accessstatus string
			query := `select accessstatus from access
					WHERE clientid = $1;`
			row := connDB.QueryRow(query, userID)

			err = row.Scan(&accessstatus)

			t.Require().NoError(err)
			t.Assert().Equal(test.accessHome.AccessStatus, accessstatus)
		})
	}
}

func (s *MyFirstSuite) TestGetListUserHome(t provider.T) {
	tests := []struct {
		nameTest string
		lenList  int
		user     factory.ObjectSystem
	}{
		{
			nameTest: "Test1",
			lenList:  1,
			user:     factory.New("user", ""),
		},
		{
			nameTest: "Test2",
			lenList:  10,
			user:     factory.New("user", ""),
		},
	}

	repos := repository.NewRepository(connDB)

	for _, test := range tests {
		t.Run(test.nameTest, func(t provider.T) {
			newUser := test.user.(*method.TestUser)

			clientID, err := newUser.InsertObject(connDB)

			t.Require().NoError(err)

			listHome := make([]pkg.ClientHome, test.lenList)
			for i := 0; i < test.lenList; i++ {
				newHome := factory.New("home", "")
				home := newHome.(*method.TestHome)

				newAccess := factory.New("access", "")
				access := newAccess.(*method.TestAccess)

				homeID, err := home.InsertObject(connDB)
				t.Require().NoError(err)

				access.ClientID = clientID
				access.HomeID = homeID
				_, err = access.InsertObject(connDB)
				t.Require().NoError(err)

				home.Home.ID = homeID
				listHome[i].Home = home.Name
				listHome[i].Username = newUser.Username
				listHome[i].Email = newUser.Email
				listHome[i].AccessLevel = access.AccessLevel
				listHome[i].AccessStatus = "active"
			}

			resultListHome, err := repos.IAccessHomeRepo.GetListUserHome(clientID)

			t.Require().NoError(err)

			t.Assert().Equal(listHome, resultListHome)
		})
	}
}
