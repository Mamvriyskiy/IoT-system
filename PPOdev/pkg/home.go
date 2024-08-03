package pkg

type Home struct {
	Name    string `db:"name"    json:"name"`
	OwnerID int    `db:"ownerid" json:"ownerId"`
	ID      int    `db:"homeid"  json:"-"`
}
