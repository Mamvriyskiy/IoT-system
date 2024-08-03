package pkg

type DeviceCharacteristics struct {
	ID              int     `db:"id"       json:"-"`
	DeviceID        int     `db:"deviceid"         json:"deviceid"`
	Values          float64 `db:"valueschar" json:"values"`
	TypeCharacterID int     `db:"typecharacterid"   json:"typecharacterid"`
}

type TypeCharacter struct {
	ID          int     `db:"id"       json:"-"`
	Type        string     `db:"typecharacter"         json:"type"`
	UnitMeasure string`db:"unitmeasure" json:"unitmeasure"`
}
