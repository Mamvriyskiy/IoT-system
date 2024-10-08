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

func (s *DeviceService) GetListDevices(homeID string) ([]pkg.DevicesInfo, error) {
	return s.repo.GetListDevices(homeID)
}

func (s *DeviceService) CreateDevice(homeID string, device pkg.Devices) (pkg.Devices, error) {
	deviceTypes := []string{"Type1", "Type2", "Type3"}
	brands := []string{"Brand1", "Brand2", "Brand3"}

	device.TypeDevice = deviceTypes[UseCryptoRandIntn(len(deviceTypes))]
	device.Status = "Inactive"
	device.Brand = brands[UseCryptoRandIntn(len(brands))]

	character := pkg.DeviceCharacteristics{
		Values: 123.45,
	}

	typeCharacter := pkg.TypeCharacter{
		Type: "weight",
		UnitMeasure: "kg",
	}

	var err error
	device.DeviceID, err = s.repo.CreateDevice(homeID, device, character, typeCharacter)

	return device, err
}

func (s *DeviceService) DeleteDevice(deviceID string) error {
	return s.repo.DeleteDevice(deviceID)
}

func (s *DeviceService) GetDeviceByID(deviceID string) (pkg.Devices, error) {
	return s.repo.GetDeviceByID(deviceID)
}

func (s *DeviceService) GetInfoDevice(deviceID string) (pkg.Devices, error) {
	return s.repo.GetDeviceByID(deviceID)
}
