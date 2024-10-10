package pkg

type Devices struct {
	Name string `db:"name" json:"name"`
}

type DevicesHandler struct {
	Devices
}

type DevicesService struct {
	Devices
	TypeDevice string
	Status     string
	Brand      string
}

type DevicesData struct {
	Devices
	ID         string `db:"deviceid"`
	HomeID         string `db:"homeid"`
	TypeDevice string `db:"typedevice" json:"typedevice"`
	Status     string `db:"status" json:"status"`
	Brand      string `db:"brand" json:"brand"`
}
