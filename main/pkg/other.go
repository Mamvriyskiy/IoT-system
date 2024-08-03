package pkg

type ClientHome struct {
	Username     string `db:"login"`
	AccessStatus string `db:"accessstatus"`
	AccessLevel  int    `db:"accesslevel"`
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
}

type Email struct {
	Email string `db:"email" json:"email"`
	Code  int    `db:"code" json:"code"`
	Token string `db:"token"`
}
