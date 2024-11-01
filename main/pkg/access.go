package pkg

type Access struct {
	AccessLevel  int    `json:"accesslevel" db:"accesslevel"`
	AccessStatus string `json:"accessstatus" db:"accessstatus"`
	Email        string `json:"email" db:"email"`
}

type AccessInfo struct {
	ID    string `db:"accessid"`
	Home  string `json:"home" db:"name"`
	Login string `db:"login" json:"login"`
	Access
}

type AccessHandler struct {
	Access
}

type AccessService struct {
	Access
}

type AccessInfoData struct {
	AccessInfo
	ClientID string `db:"clientid"`
	HomeID   string `db:"homeid"`
}
