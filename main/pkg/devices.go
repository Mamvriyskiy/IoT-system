package pkg

type Devices struct {
	Home       string `json:"home"`
	Name       string `db:"name"             json:"name"`
	TypeDevice string `db:"typedevice"       json:"typeDevice"`
	Status     string `db:"status"           json:"status"`
	Brand      string `db:"brand"            json:"brand"`
	DeviceID   string    `db:"deviceid"         json:"deviceID"`
	HomeID 	   string
}

type DevicesInfo struct {
	Name       string `db:"name"             json:"name"`
	Status     string `db:"status"           json:"status"`
	Brand      string `db:"brand"            json:"brand"`
	DeviceID   string    `db:"deviceid"         json:"deviceID"`
}
