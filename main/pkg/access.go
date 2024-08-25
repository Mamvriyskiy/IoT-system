package pkg

type AccessHome struct {
	AccessStatus string `json:"status"`
	ID           int    `db:"accessID" json:"-"`
	AccessLevel  int    `json:"level"`
}

type Access struct {
	Home        string `json:"home"`
	Email       string `json:"email"`
	AccessLevel int `json:"accesslevel"`
}
