package service

import (
	"crypto/rand"
	"math/big"

	"github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/Mamvriyskiy/database_course/main/pkg/repository"
)

type DeviceService struct {
	repo repository.IDeviceRepo
}

func NewDeviceService(repo repository.IDeviceRepo) *DeviceService {
	return &DeviceService{repo: repo}
}

func UseCryptoRandIntn(max int) int {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		//TODO: error
		return 0
	}
	return int(n.Int64())
}

func (s *DeviceService) GetListDevices(userID int) ([]pkg.Devices, error) {
	return s.repo.GetListDevices(userID)
}

func (s *DeviceService) CreateDevice(homeID int, device *pkg.Devices) (int, error) {
	deviceTypes := []string{"Type1", "Type2", "Type3"}
	statusValues := []string{"Active", "Inactive"}
	brands := []string{"Brand1", "Brand2", "Brand3"}

	device.TypeDevice = deviceTypes[UseCryptoRandIntn(len(deviceTypes))]
	device.Status = statusValues[UseCryptoRandIntn(len(statusValues))]
	device.Brand = brands[UseCryptoRandIntn(len(brands))]

	character := pkg.DeviceCharacteristics{
		Values: 123.45,
	}

	typeCharacter := pkg.TypeCharacter{
		Type: "weight",
		UnitMeasure: "kg",
	}

	return s.repo.CreateDevice(homeID, device, character, typeCharacter)
}

func (s *DeviceService) DeleteDevice(idDevice int, name, home string) error {
	return s.repo.DeleteDevice(idDevice, name, home)
}

func (s *DeviceService) GetDeviceByID(deviceID int) (pkg.Devices, error) {
	return s.repo.GetDeviceByID(deviceID)
}
