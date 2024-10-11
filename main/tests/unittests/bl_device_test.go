package unittests

import (
	"github.com/Mamvriyskiy/database_course/main/pkg"
	mocks_service "github.com/Mamvriyskiy/database_course/main/pkg/repository/mocks"
	"github.com/Mamvriyskiy/database_course/main/pkg/service"
	"github.com/Mamvriyskiy/database_course/main/tests/factory"
	method "github.com/Mamvriyskiy/database_course/main/tests/method"
	"github.com/golang/mock/gomock"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

func (s *MyUnitTestsSuite) TestCreateDeviceBL(t provider.T) {
	tests := []struct {
		nameTest string
		device   factory.ObjectSystem
		homeID   string
		deviceID string
	}{
		{
			nameTest: "Test1",
			device:   factory.New("device", ""),
			homeID:   "1",
			deviceID: "4",
		},
		{
			nameTest: "Test2",
			device:   factory.New("device", ""),
			homeID:   "2",
			deviceID: "5",
		},
		{
			nameTest: "Test3",
			device:   factory.New("device", ""),
			homeID:   "3",
			deviceID: "6",
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks_service.NewMockIDeviceRepo(ctrl)

	for _, test := range tests {
		t.Run(test.nameTest, func(t provider.T) {
			newDevice := test.device.(*method.TestDevice)
			newDevice.ID = test.deviceID

			mockRepo.EXPECT().CreateDevice(test.homeID, gomock.Any(), gomock.Any(), gomock.Any()).Return(test.deviceID, nil)
			mockRepo.EXPECT().GetDeviceByID(test.deviceID).Return(newDevice.DevicesData, nil)

			userService := service.NewDeviceService(mockRepo)

			devicesHandler := pkg.DevicesHandler{
				Devices: newDevice.Devices,
			}

			devices, err := userService.CreateDevice(test.homeID, devicesHandler)
			
			t.Assert().NoError(err)
			t.Assert().Equal(newDevice.DevicesData, devices)
		})
	}
}

func (s *MyUnitTestsSuite) TestDeleteDeviceBL(t provider.T) {
	tests := []struct {
		nameTest   string
		deviceID   string
		nameDevice string
		homeName   string
	}{
		{
			nameTest:   "Test1",
			deviceID:   "1",
			nameDevice: "dev1",
			homeName:   "home1",
		},
		{
			nameTest:   "Test2",
			deviceID:   "2",
			nameDevice: "dev2",
			homeName:   "home2",
		},
		{
			nameTest:   "Test3",
			deviceID:   "3",
			nameDevice: "dev3",
			homeName:   "home3",
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks_service.NewMockIDeviceRepo(ctrl)

	for _, test := range tests {
		t.Run(test.nameTest, func(t provider.T) {
			mockRepo.EXPECT().DeleteDevice(test.deviceID).Return(nil)

			userService := service.NewDeviceService(mockRepo)

			err := userService.DeleteDevice(test.deviceID)

			t.Assert().NoError(err)
		})
	}
}
