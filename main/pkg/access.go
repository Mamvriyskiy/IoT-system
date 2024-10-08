package pkg

type AccessHome struct {
	AccessStatus string `json:"status"`
	ID           string    `db:"accessID" json:"-"`
	AccessLevel  int    `json:"level"`
}

type Access struct {
	ID           string    `db:"accessid"`
	Home         string `json:"home" db:"home"`
	Email        string `json:"email" db:"email"`
	AccessLevel  int    `json:"accesslevel" db:"accesslevel"`
	AccessStatus string `json:"accessstatus" db:"accessstatus"`
	ClientID     string
	HomeID       string
	Login        string `db:"login" json:"login"`
}
