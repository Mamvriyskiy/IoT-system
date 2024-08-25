package testsdatabase

import (
	"testing"
	//"github.com/jmoiron/sqlx"
	"github.com/Mamvriyskiy/database_course/main/pkg/repository"
	"github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/stretchr/testify/assert"
)

func TestCreateHome(t *testing.T) {
	tests := []struct {
		nameTest string
		home     pkg.Home
	}{
		{
			nameTest: "Test1",
			home: pkg.Home{
				Name: "home1",
				GeographCoords: 12345,
			},
		},
		{
			nameTest: "Test2",
			home: pkg.Home{
				Name: "home2",
				GeographCoords: 23456,
			},
		},
		{
			nameTest: "Test3",
			home: pkg.Home{
				Name: "home3",
				GeographCoords: 34567,
			},
		},
	}

	repos := repository.NewRepository(connDB)

	for _, test := range tests {
		t.Run(test.nameTest, func(t *testing.T) {
			homeID, err := repos.IHomeRepo.CreateHome(test.home)
			if err != nil {
				assert.NoError(t, err, "")
			}
			test.home.ID = homeID

			query := `SELECT name, coords FROM home WHERE homeid = $1`
			row := connDB.QueryRow(query, homeID)

			retrievedHome:= pkg.Home{
				ID: homeID,
			}
		
			err = row.Scan(&retrievedHome.Name, &retrievedHome.GeographCoords)
			if err != nil {
				assert.NoError(t, err, "")
			}

			assert.Equal(t, test.home, retrievedHome, "The passwords should be the same.")
		})
	}
}
