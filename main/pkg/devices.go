package pkg

type Devices struct {
	Name string `db:"name" json:"Name"`
}

type DevicesHandler struct {
	Devices
}

type DevicesInfo struct {
	TypeDevice string `db:"typedevice" json:"TypeDevice"`
	Status     string `db:"status" json:"Status"`
	Brand      string `db:"brand" json:"Brand"`
}

type DevicesData struct {
	Devices
	ID     string `db:"deviceid"`
	HomeID string `db:"homeid"`
	DevicesInfo
}

type DevicesService struct {
	Devices
	DevicesInfo
}

