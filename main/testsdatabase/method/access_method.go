package method 

import (
	"crypto/rand"
	"math/big"
	"github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/jmoiron/sqlx"
	"fmt"
)

const (
	lengthName = 6
	lengthGeographCoords = 8
)

type TestAccess struct {
	pkg.Access
}

func NewAccess() *TestAccess {
	var b TestAccess

	return b.BuilderAccess()
}

func (b *TestAccess) BuilderAccess() *TestAccess {
	b.generateLevel()

	return b
}

func (b *TestAccess) generateLevel() {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(5)))
	if err != nil {
		b.AccessLevel = 1
	}

	b.AccessLevel =  int(n.Int64())
}

func (tu TestAccess) InsertObject(connDB *sqlx.DB) (int, error) {
	query := fmt.Sprintf(`INSERT INTO %s (accessStatus, accessLevel, HomeID, clientid) 
		values ($1, $2, $3, $4) RETURNING accessID`, "access")
	_ = connDB.QueryRow(query, "active", tu.AccessLevel, tu.HomeID, tu.ClientID)

	return 0, nil
}
