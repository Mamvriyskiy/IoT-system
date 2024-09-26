package pkg

type User struct {
	Password string `db:"password" json:"password"`
	Username string `db:"login"    json:"login"`
	Email    string `db:"email"    json:"email"`
	ID       int    `db:"clientid" json:"-"`
}
