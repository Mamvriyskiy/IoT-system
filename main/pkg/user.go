package pkg

type User struct {
	Password string `db:"password" json:"password"`
	Email    string `db:"email"    json:"email"`
	Username string `db:"login"    json:"login"`
}

type UserHandler struct {
	User
}

type UserService struct {
	User
}

type UserData struct {
	User
	ID string `db:"clientid"`
}
