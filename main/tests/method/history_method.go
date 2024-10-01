package method 

import (
	"crypto/rand"
	"math/big"
	"github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/jmoiron/sqlx"
	"fmt"
)


type TestHistory struct {
	pkg.AddHistory
}

func NewHistory() *TestHistory {
	var b TestHistory

	return b.BuilderAccess()
}

func (b *TestHistory) BuilderAccess() *TestHistory {
	b.TimeWork = b.generateValues()
	b.AverageIndicator = float64(b.generateValues())
	b.EnergyConsumed = b.generateValues()

	return b
}

func (b *TestHistory) generateValues() int {
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(100)))

	return int(n.Int64())
}

func (tu TestHistory) InsertObject(connDB *sqlx.DB) (int, error) {
	var id int
	query := fmt.Sprintf(`INSERT INTO %s 
		(timeWork, AverageIndicator, EnergyConsumed) 
			values ($1, $2, $3) RETURNING historyDevID`, "historyDev")
	row := connDB.QueryRow(query, tu.TimeWork, tu.AverageIndicator, tu.EnergyConsumed)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	query = fmt.Sprintf("INSERT INTO %s (deviceID, historydevID) VALUES ($1, $2)", "historydevice")
	_, err = connDB.Exec(query, tu.DeviceID, id)
	if err != nil {
		return 0, err
	}

	return 0, nil
}
