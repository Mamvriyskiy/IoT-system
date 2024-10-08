package pkg

type ClientHome struct {
	Home         string `db:"name" json:"name"`
	Username     string `db:"login" json:"login"`
	Email        string `db:"email" json:"email"`
	AccessStatus string `db:"accessstatus" json:"accessstatus"`
	AccessLevel  int    `db:"accesslevel" json:"accesslevel"`
}

type AddUserHome struct {
	Email       string `json:"email"`
	AccessLevel int    `json:"accessLevel"`
}

type AddHistory struct {
	Name             string  `db:"name"             json:"name"`
	TimeWork         int     `db:"timework"         json:"timework"`
	AverageIndicator float64 `db:"averageindicator" json:"average"`
	EnergyConsumed   int     `db:"energyconsumed"   json:"energy"`
	Home             string  `db:"home"   			json:"home"`
	DeviceID		 string     `db:"deviceid"   		json:"deviceid"`
}

type Email struct {
	Email string `db:"email" json:"email"`
	Code  int    `db:"code" json:"code"`
	Token string `db:"token"`
}
