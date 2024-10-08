package endtoendtests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"crypto/sha256"
	"encoding/hex"

	"github.com/gin-gonic/gin"
	"github.com/Mamvriyskiy/database_course/main/pkg/handler"
	"github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/Mamvriyskiy/database_course/main/pkg/service"
	"github.com/Mamvriyskiy/database_course/main/pkg/repository"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/Mamvriyskiy/database_course/main/tests/factory"
	method "github.com/Mamvriyskiy/database_course/main/tests/method"
	//"fmt"
)

const (
	salt = "hfdjmaxckdk20"
)

type Response struct {
    ID string `json:"id"`
}

func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return hex.EncodeToString(hash.Sum([]byte(salt)))
}

func (s *MyEtoESuite) TestSignUp(t provider.T) {
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

	gin.SetMode(gin.TestMode)
	router := gin.New()

	repos := repository.NewRepository(connDB)
	services := service.NewServicesPsql(repos)
	handlers := handler.NewHandler(services)

	auth := router.Group("/auth")
	auth.POST("/sign-up", handlers.SignUp)

	for _, test := range tests {
		t.Run(test.nameTest, func(t provider.T) {
			newUser := test.user.(*method.TestUser)
			
			body, err := json.Marshal(newUser.User)
			t.Require().NoError(err)

			req, err := http.NewRequest(http.MethodPost, "/auth/sign-up", bytes.NewBuffer(body))
			t.Require().NoError(err)

			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			var response Response
			err = json.NewDecoder(w.Body).Decode(&response)
			t.Require().NoError(err)


			newUser.User.ID = response.ID
			query := `SELECT password, login, email FROM client WHERE clientid = $1`
			row := connDB.QueryRow(query, response.ID)

			retrievedUser := pkg.User{
				ID: response.ID,
			}

			newUser.Password = generatePasswordHash(newUser.Password)
			err = row.Scan(&retrievedUser.Password, &retrievedUser.Username, &retrievedUser.Email)
			t.Require().NoError(err)
			t.Assert().Equal(newUser.User, retrievedUser)
		})
	}
}

