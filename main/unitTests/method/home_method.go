package method 

import (
	"crypto/rand"
	"math/big"
	"github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/jmoiron/sqlx"
	"fmt"
)

const (

)

type TestHome struct {
	pkg.Home
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
		b.GeographCoords = 11111111
	}

	b.GeographCoords = int(n.Int64())
}

func (tu TestHome) InsertObject(connDB *sqlx.DB) (int, error) {
	var homeID int
	query := fmt.Sprintf("INSERT INTO %s (coords, name) values ($1, $2) RETURNING homeID", "home")
	row := connDB.QueryRow(query, tu.GeographCoords, tu.Name)
	if err := row.Scan(&homeID); err != nil {
		return 0, err
	}

	return homeID, nil
}
