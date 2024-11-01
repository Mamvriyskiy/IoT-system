package pkg

type EmailUser struct {
	Email string `db:"email" json:"email"`
}

type EmailHandler struct {
	EmailUser
}

type EmailService struct {
	EmailUser
	Code  int    `db:"code" json:"code"`
	Token string `db:"token"`
}
