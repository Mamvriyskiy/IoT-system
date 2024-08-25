package pkg

type Home struct {
	Name           string `db:"name"    json:"name"`
	GeographCoords int    `db:"coords" json:"coords"`
	ID             int    `db:"homeid"  json:"-"`
	UserID 		   int
}

type UpdateNameHome struct {
	LastName string `json:"lastname"`
	NewName string ` json:"newname"`
	UserID int		`json:"-"`
}
