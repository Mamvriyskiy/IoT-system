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

func (s *DeviceService) GetListDevices(homeID string) ([]pkg.DevicesData, error) {
	return s.repo.GetListDevices(homeID)
}

func (s *DeviceService) CreateDevice(homeID string, device pkg.DevicesHandler) (pkg.DevicesData, error) {
	deviceTypes := []string{"Type1", "Type2", "Type3"}
	brands := []string{"Brand1", "Brand2", "Brand3"}

	deviceService := pkg.DevicesService{
		Devices: device.Devices,
		DevicesInfo: pkg.DevicesInfo{
			TypeDevice: deviceTypes[UseCryptoRandIntn(len(deviceTypes))],
			Status: "Inactive", 
			Brand: brands[UseCryptoRandIntn(len(brands))],
		},
	}

	character := pkg.DeviceCharacteristicsService{
		Values: 123.45,
	}

	typeCharacter := pkg.TypeCharacterService{
		Type: "weight",
		UnitMeasure: "kg",
	}

	var err error
	deviceID, err := s.repo.CreateDevice(homeID, deviceService, character, typeCharacter)
	if err != nil {
		return pkg.DevicesData{}, err
	}
	
	return s.repo.GetDeviceByID(deviceID)
}

func (s *DeviceService) DeleteDevice(deviceID string) error {
	return s.repo.DeleteDevice(deviceID)
}

func (s *DeviceService) GetDeviceByID(deviceID string) (pkg.DevicesData, error) {
	return s.repo.GetDeviceByID(deviceID)
}

func (s *DeviceService) GetInfoDevice(deviceID string) (pkg.DevicesData, error) {
	return s.repo.GetDeviceByID(deviceID)
}
