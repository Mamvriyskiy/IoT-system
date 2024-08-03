package pkg

type Devices struct {
	Name             string `db:"name"             json:"name"`
	TypeDevice       string `db:"typedevice"       json:"typeDevice"`
	Status           string `db:"status"           json:"status"`
	Brand            string `db:"brand"            json:"brand"`
	DeviceID         int    `db:"deviceid"         json:"-"`
	PowerConsumption int    `db:"powerconsumption" json:"powerConsumption"`
	MinParameter     int    `db:"minparametr"      json:"minParametr"`
	MaxParameter     int    `db:"maxparametr"      json:"maxParametr"`
}
