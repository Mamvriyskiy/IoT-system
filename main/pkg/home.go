package pkg

type Home struct {
	Name           string `db:"name"    json:"name"`
	GeographCoords int    `db:"coords" json:"coords"`
	ID             int    `db:"homeid"  json:"-"`
	UserID 		   int
}
