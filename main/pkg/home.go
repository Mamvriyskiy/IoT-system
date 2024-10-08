package pkg

type Home struct {
	Name      string  `db:"name"    json:"name"`
	Latitude  float64 `db:"latitude" json:"latitude"`
	Longitude float64 `db:"longitude" json:"longitude"`
	ID        string     `db:"homeid"  json:"homeid"`
}
