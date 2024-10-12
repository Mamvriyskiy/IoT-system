package pkg

type DeviceCharacteristicsService struct {
	ID              string     `db:"id"       json:"-"`
	DeviceID        string     `db:"deviceid"         json:"deviceid"`
	Values          float64 `db:"valueschar" json:"values"`
	TypeCharacterID int     `db:"typecharacterid"   json:"typecharacterid"`
}

type TypeCharacterService struct {
	ID          string     `db:"id"       json:"-"`
	Type        string     `db:"typecharacter"         json:"type"`
	UnitMeasure string`db:"unitmeasure" json:"unitmeasure"`
}

