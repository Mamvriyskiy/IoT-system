package pkg

type AccessHome struct {
	AccessStatus string `json:"status"`
	ID           int    `db:"accessID" json:"-"`
	AccessLevel  int    `json:"level"`
}

type Access struct {
	ID           int    `db:"accessid"`
	Home         string `json:"home" db:"home"`
	Email        string `json:"email" db:"email"`
	AccessLevel  int    `json:"accesslevel" db:"accesslevel"`
	AccessStatus string `json:"accessstatus" db:"accessstatus"`
	ClientID     int
	HomeID       int
	Login        string `db:"login" json:"login"`
}
