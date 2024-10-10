package pkg

type Home struct {
	Name      string  `db:"name"    json:"name"`
	Latitude  float64 `db:"latitude" json:"latitude"`
	Longitude float64 `db:"longitude" json:"longitude"`
}

type HomeHandler struct {
	Home
}

type HomeService struct {
	Home
}

type HomeData struct {
	Home
	ID string `db:"homeid"`
}
