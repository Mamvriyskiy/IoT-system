package method 

import (
	"crypto/rand"
	"math/big"
	"github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/jmoiron/sqlx"
	"github.com/google/uuid"
	"fmt"
)

const (

)

type TestHome struct {
	pkg.HomeData
}

func NewHome() *TestHome {
	var b TestHome

	return b.BuilderHome()
}

func (b *TestHome) BuilderHome() *TestHome {
	b.generateName()
	b.generateGeographCoords()

	return b
}

func (b *TestHome) generateName() {
	name := make([]byte, lengthName)
	for j := 0; j < lengthName; j++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		name[j] = charset[n.Int64()]
	}

	b.Name = string(name)
}

func (b *TestHome) generateGeographCoords() {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(lengthGeographCoords)))
	if err != nil {
		b.Latitude = 11111111
		b.Longitude = 11111111
	}

	b.Latitude = float64(n.Int64())
	b.Longitude = float64(n.Int64())
}

func (tu TestHome) InsertObject(connDB *sqlx.DB) (string, error) {
	homeID := uuid.New()
	var id string
	query := fmt.Sprintf("INSERT INTO %s (latitude, longitude, name, homeID) values ($1, $2, $3, $4) RETURNING homeID", "home")
	row := connDB.QueryRow(query, tu.Latitude, tu.Longitude, tu.Name, homeID)

	err := row.Scan(&id)

	return id, err
}
