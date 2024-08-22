package testsdatabase

import (
	"testing"
	"context"
	"os"
	"github.com/jmoiron/sqlx"
	//"github.com/testcontainers/testcontainers-go"
	"github.com/Mamvriyskiy/database_course/main/migrations"
	"github.com/Mamvriyskiy/database_course/main/containers"
	"github.com/Mamvriyskiy/database_course/main/pkg/repository"
	"github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/stretchr/testify/assert"
)

var connDB *sqlx.DB

func TestMain(m *testing.M) {
	dbTestContainers, db, err := containers.SetupTestDataBase()

	if err != nil {
		panic(err)
	}
	defer dbTestContainers.Terminate(context.Background())

	connDB = db
	err = migrations.MigrationsTestDataBase(connDB)
	if err != nil {
		panic(err)
	}

	code := m.Run()

	os.Exit(code)
}

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
				ID: 1,
			},
		},
		{
			nameTest: "Test2",
			user: pkg.User{
				Password: "pswrd2",
				Username: "user2",
				Email:    "email2",
				ID: 2,
			},
		},
		{
			nameTest: "Test3",
			user: pkg.User{
				Password: "pswrd3",
				Username: "user3",
				Email:    "email3",
				ID: 3,
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
		nameTest string
		user     pkg.User
	}{
		{
			nameTest: "Test1",
			user: pkg.User{
				Password: "pswrd4",
				Username: "user4",
				Email:    "email4",
			},
		},
		{
			nameTest: "Test2",
			user: pkg.User{
				Password: "pswrd5",
				Username: "user5",
				Email:    "email5",			
			},
		},
		{
			nameTest: "Test3",
			user: pkg.User{
				Password: "pswrd6",
				Username: "user6",
				Email:    "email6",
			},
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

			resultUser, err := repos.IUserRepo.GetUser(test.user.Username, test.user.Password)
			assert.NoError(t, err, "")
			assert.Equal(t, clientID, resultUser.ID, "The clientIDs should be the same.")
		})
	}
}
