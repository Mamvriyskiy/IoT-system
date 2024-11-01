package pkg

type Devices struct {
	Name string `db:"name" json:"name"`
}

type DevicesHandler struct {
	Devices
}

type DevicesInfo struct {
	TypeDevice string `db:"typedevice" json:"typedevice"`
	Status     string `db:"status" json:"status"`
	Brand      string `db:"brand" json:"brand"`
}

type DevicesService struct {
	Devices
	DevicesInfo
}

type DevicesData struct {
	Devices
	ID     string `db:"deviceid"`
	HomeID string `db:"homeid"`
	DevicesInfo
}
