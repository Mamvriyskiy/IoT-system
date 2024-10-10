package pkg

type History struct {
	TimeWork         int     `db:"timework"         json:"timework"`
	AverageIndicator float64 `db:"averageindicator" json:"average"`
	EnergyConsumed   int     `db:"energyconsumed"   json:"energy"`
}

type DevicesHistoryData struct {
	History
	ID               string     `db:"historydevid"       json:"-"`
}


type HistoryService struct {
	History
}

