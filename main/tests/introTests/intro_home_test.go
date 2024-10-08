package introtests

import (
	// "github.com/jmoiron/sqlx"
	"errors"
	"github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/Mamvriyskiy/database_course/main/pkg/repository"
	"github.com/Mamvriyskiy/database_course/main/pkg/service"
	"github.com/Mamvriyskiy/database_course/main/tests/factory"
	method "github.com/Mamvriyskiy/database_course/main/tests/method"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

func (s *MyIntroTestsSuite) TestCreateHomeIntro(t provider.T) {
	tests := []struct {
		nameTest string
		home     factory.ObjectSystem
	}{
		{
			nameTest: "Test1",
			home:     factory.New("home", ""),
		},
		{
			nameTest: "Test2",
			home:     factory.New("home", ""),
		},
		{
			nameTest: "Test3",
			home:     factory.New("home", ""),
		},
	}

	repos := repository.NewRepository(connDB)
	services := service.NewServicesPsql(repos)

	for _, test := range tests {
		t.Run(test.nameTest, func(t provider.T) {
			newHome := test.home.(*method.TestHome)

			homeID, err := services.IHome.CreateHome(newHome.Home)
			t.Require().NoError(err)

			newHome.ID = homeID

			query := `SELECT name, Latitude, Longitude FROM home WHERE homeid = $1`
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

func (s *MyIntroTestsSuite) TestGetListHomeIntro(t provider.T) {
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
	services := service.NewServicesPsql(repos)

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

			resultListHome, err := services.IHome.ListUserHome(clientID)

			t.Require().NoError(err)

			t.Assert().Equal(listHome, resultListHome)
		})
	}
}

func (s *MyIntroTestsSuite) TestUpdateHomeIntro(t provider.T) {
	tests := []struct {
		nameTest string
		home     factory.ObjectSystem
		name     string
	}{
		{
			nameTest: "Test1",
			home:     factory.New("home", ""),
			name: "home1",
		},
		{
			nameTest: "Test2",
			home:     factory.New("home", ""),
			name: "home1",
		},
		{
			nameTest: "Test3",
			home:     factory.New("home", ""),
			name: "home1",
		},
	}

	repos := repository.NewRepository(connDB)
	services := service.NewServicesPsql(repos)

	for _, test := range tests {
		t.Run(test.nameTest, func(t provider.T) {
			homeID, err := test.home.InsertObject(connDB)
			t.Require().NoError(err)

			err = services.IHome.UpdateHome(homeID, test.name)
			t.Require().NoError(err)

			query := `SELECT name FROM home WHERE homeid = $1`
			row := connDB.QueryRow(query, homeID)

			var resultName string
			err = row.Scan(&resultName)
			t.Require().NoError(err)

			t.Assert().Equal(test.name, resultName)
		})
	}
}

func (s *MyIntroTestsSuite) TestDeleteHomeIntro(t provider.T) {
	tests := []struct {
		nameTest string
		user     factory.ObjectSystem
		home     factory.ObjectSystem
		access   factory.ObjectSystem
	}{
		{
			nameTest: "Test1",
			home:     factory.New("home", ""),
		},
		{
			nameTest: "Test2",
			home:     factory.New("home", ""),
		},
		{
			nameTest: "Test3",
			home:     factory.New("home", ""),
		},
	}

	repos := repository.NewRepository(connDB)
	services := service.NewServicesPsql(repos)

	for _, test := range tests {
		t.Run(test.nameTest, func(t provider.T) {
			newHome := test.home.(*method.TestHome)

			homeID, err := newHome.InsertObject(connDB)
			t.Require().NoError(err)

			err = services.IHome.DeleteHome(homeID)
			t.Require().NoError(err)

			query := `SELECT name FROM home WHERE homeid = $1`
			row := connDB.QueryRow(query, homeID)

			var resultName string
			err = row.Scan(&resultName)

			t.Assert().Equal(err, errors.New("sql: no rows in result set"))
		})
	}
}
