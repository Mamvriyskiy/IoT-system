package pkg

type UpdatePassword struct {
	Password string `db:"password" json:"password"`
	Token    string `db:"token" json:"token"`
}

type VerifyCode struct {
	Code  string `db:"code" json:"code"`
	Token string `db:"token" json:"token"`
}
