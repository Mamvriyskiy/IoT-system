package pkg

type AccessHome struct {
	AccessStatus string `json:"status"`
	ID           int    `db:"accessID" json:"-"`
	AccessLevel  int    `json:"level"`
}
