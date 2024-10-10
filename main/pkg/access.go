package pkg

type Access struct {
	AccessLevel int    `json:"accesslevel" db:"accesslevel"`
	AccessStatus string `json:"accessstatus" db:"accessstatus"`
}

type AccessInfo struct {
	ID           string `db:"accessid"`
	Home         string `json:"home" db:"name"`
	Email        string `json:"email" db:"email"`
	Login        string `db:"login" json:"login"`
	Access
}

type AccessHandler struct {
	Access
	Email       string `json:"email" db:"email"`
}

type AccessService struct {
	Access
	Email       string `json:"email" db:"email"`
	ClientID           string `db:"clientid"`
	HomeID           string `db:"homeid"`
}

type AccessInfoData struct {
	AccessInfo
}
