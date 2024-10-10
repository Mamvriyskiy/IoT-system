package unittests

import (
	// "github.com/jmoiron/sqlx"
	"github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/Mamvriyskiy/database_course/main/pkg/repository"
	"github.com/Mamvriyskiy/database_course/main/tests/factory"
	method "github.com/Mamvriyskiy/database_course/main/tests/method"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"errors"
)

func (s *MyUnitTestsSuite) TestCreateHome(t provider.T) {
	tests := []struct {
		nameTest string
		home     factory.ObjectSystem
	}{
		{
			nameTest: "Test1",
			home:     factory.New("home", "", "DB"),
		},
		{
			nameTest: "Test2",
			home:     factory.New("home", "", "DB"),
		},
		{
			nameTest: "Test3",
			home:     factory.New("home", "", "DB"),
		},
	}

	repos := repository.NewRepository(connDB)

	for _, test := range tests {
		t.Run(test.nameTest, func(t provider.T) {
			newHome := test.home.(*method.TestHome)

			homeID, err := repos.IHomeRepo.CreateHome(newHome.Home)
			t.Require().NoError(err)

			newHome.ID = homeID

			query := `SELECT name, latitude, longitude FROM home WHERE homeid = $1`
			row := connDB.QueryRow(query, homeID)

			retrievedHome := pkg.Home{
				ID: homeID,
			}

			err = row.Scan(&retrievedHome.Name, &retrievedHome.Latitude, &retrievedHome.Longitude)
			t.Require().NoError(err)

			t.Assert().Equal(newHome.Home, retrievedHome)
		})
	}
}

func (s *MyUnitTestsSuite) TestGetListHome(t provider.T) {
	tests := []struct {
		nameTest string
		lenList  int
		user     factory.ObjectSystem
	}{
		{
			nameTest: "Test1",
			lenList:  1,
			user:     factory.New("user", "", "DB"),
		},
		{
			nameTest: "Test2",
			lenList:  10,
			user:     factory.New("user", "", "DB"),
		},
	}

	repos := repository.NewRepository(connDB)

	for _, test := range tests {
		t.Run(test.nameTest, func(t provider.T) {
			newUser := test.user.(*method.TestUser)

			clientID, err := newUser.InsertObject(connDB)

			t.Require().NoError(err)

			listHome := make([]pkg.Home, test.lenList)
			for i := 0; i < test.lenList; i++ {
				newHome := factory.New("home", "")
				home := newHome.(*method.TestHome)

				homeID, err := home.InsertObject(connDB)
				t.Require().NoError(err)

				home.Home.ID = homeID
				listHome[i] = home.Home

				newAccess := factory.New("access", "")
				access := newAccess.(*method.TestAccess)

				access.ClientID = clientID
				access.HomeID = homeID

				_, err = access.InsertObject(connDB)
				t.Require().NoError(err)
			}

			resultListHome, err := repos.IHomeRepo.ListUserHome(clientID)

			t.Require().NoError(err)

			t.Assert().Equal(listHome, resultListHome)
		})
	}
}

func (s *MyUnitTestsSuite) TestUpdateHome(t provider.T) {
	tests := []struct {
		nameTest   string
		home       factory.ObjectSystem
		updateHome string
	}{
		{
			nameTest: "Test1",
			home:     factory.New("home", "", "DB"),
			updateHome: "home1",
		},
		{
			nameTest: "Test2",
			home:     factory.New("home", "", "DB"),
			updateHome: "home2",
		},
		{
			nameTest: "Test3",
			home:     factory.New("home", "", "DB"),
			updateHome: "home3",
		},
	}

	repos := repository.NewRepository(connDB)

	for _, test := range tests {
		t.Run(test.nameTest, func(t provider.T) {
			newHome := test.home.(*method.TestHome)

			homeID, err := newHome.InsertObject(connDB)
			t.Require().NoError(err)

			err = repos.IHomeRepo.UpdateHome(homeID, test.updateHome)
			t.Require().NoError(err)

			query := `SELECT name FROM home WHERE homeid = $1`
			row := connDB.QueryRow(query, homeID)

			var resultName string
			err = row.Scan(&resultName)
			t.Require().NoError(err)

			t.Assert().Equal(test.updateHome, resultName)
		})
	}
}

func (s *MyUnitTestsSuite) TestDeleteHome(t provider.T) {
	tests := []struct {
		nameTest   string
		home       factory.ObjectSystem
	}{
		{
			nameTest: "Test1",
			home:     factory.New("home", "", "DB"),
		},
		{
			nameTest: "Test2",
			home:     factory.New("home", "", "DB"),
		},
		{
			nameTest: "Test3",
			home:     factory.New("home", "", "DB"),
		},
	}

	repos := repository.NewRepository(connDB)

	for _, test := range tests {
		t.Run(test.nameTest, func(t provider.T) {
			newHome := test.home.(*method.TestHome)

			homeID, err := newHome.InsertObject(connDB)
			t.Require().NoError(err)

			err = repos.IHomeRepo.DeleteHome(homeID)
			t.Require().NoError(err)

			query := `SELECT name FROM home WHERE homeid = $1`
			row := connDB.QueryRow(query, homeID)

			var resultName string
			err = row.Scan(&resultName)

			t.Assert().Equal(err, errors.New("sql: no rows in result set"))
		})
	}
}
