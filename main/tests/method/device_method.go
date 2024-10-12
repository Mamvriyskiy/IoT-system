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
	length = 6
)

type TestDevice struct {
	pkg.DevicesData
}

func NewDevice() *TestDevice {
	var b TestDevice

	return b.BuilderAccess()
}

func (b *TestDevice) BuilderAccess() *TestDevice {
	b.Name = b.generateDev()
	b.TypeDevice = b.generateDev()
	b.Status = "inactive"
	b.Brand = b.generateDev()

	return b
}

func (b *TestDevice) generateDev() string {
	dev := make([]byte, length)
	for j := 0; j < length; j++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		dev[j] = charset[n.Int64()]
	}

	return string(dev)
}

func (tu TestDevice) InsertObject(connDB *sqlx.DB) (string, error) {
	deviceID := uuid.New()

	query := fmt.Sprintf(`INSERT INTO %s (name, TypeDevice, Status, Brand, homeID, deviceID)
			values ($1, $2, $3, $4, $5, $6) RETURNING deviceID`, "device")

	row := connDB.QueryRow(query, tu.Name, tu.TypeDevice,
		tu.Status, tu.Brand, uuid.New(), deviceID)
		// fmt.Println(reflect.TypeOf(row))
	
	// if tu.HomeID == "" {
	// 	row = connDB.QueryRow(query, tu.Name, tu.TypeDevice,
	// 		tu.Status, tu.Brand, uuid.New(), deviceID)
	// } else {
	// 	row = connDB.QueryRow(query, tu.Name, tu.TypeDevice,
	// 		tu.Status, tu.Brand, tu.HomeID, deviceID)
	// }

	var id string
	err := row.Scan(&id)

	return id, err
}
