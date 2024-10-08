package service

import (
	"github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/Mamvriyskiy/database_course/main/pkg/repository"
)

type IUser interface {
	CreateUser(user pkg.User) (string, error)
	CheckUser(user pkg.User) (string, error)
	GenerateToken(login, password string) (pkg.User, string, error)
	ParseToken(token string) (string, error)
	ChangePassword(password, token string) error
	CheckCode(code, token string) error
	SendCode(email pkg.Email) error
	GetUserByEmail(email string) (int, error)
	GetAccessLevel(userID, homeID string) (int, error)
}

type IHome interface {
	CreateHome(home pkg.Home) (string, error)
	DeleteHome(homeID string) error
	UpdateHome(homeID, name string) error
	GetHomeByID(homeID string) (pkg.Home, error)
	ListUserHome(userID string) ([]pkg.Home, error)
}

type IAccessHome interface {
	AddUser(homeID string, access pkg.Access) (string, error)
	DeleteUser(accessID string) error
	UpdateLevel(accessID string, updateAccess pkg.Access) error
	UpdateStatus(userID string, access pkg.AccessHome) error
	GetListUserHome(homeID string) ([]pkg.ClientHome, error)
	AddOwner(userID, homeID string) (string, error)
	GetInfoAccessByID(accessID string) (pkg.Access, error)
}

type IDevice interface {
	CreateDevice(homeID string, device pkg.Devices) (pkg.Devices, error)
	DeleteDevice(deviceID string) error
	GetDeviceByID(deviceID string) (pkg.Devices, error)
	GetListDevices(homeID string) ([]pkg.DevicesInfo, error)
	GetInfoDevice(deviceID string) (pkg.Devices, error)
}

type IHistoryDevice interface {
	CreateDeviceHistory(deviceID string) (string, error)
	GetDeviceHistory(deviceID string) ([]pkg.DevicesHistory, error)
}

type Services struct {
	IUser
	IHome
	IAccessHome
	IDevice
	IHistoryDevice
}

func NewServicesPsql(repo *repository.Repository) *Services {
	return &Services{
		IUser:          NewUserService(repo.IUserRepo),
		IHome:          NewHomeService(repo.IHomeRepo),
		IAccessHome:    NewAccessHomeService(repo.IAccessHomeRepo),
		IDevice:        NewDeviceService(repo.IDeviceRepo),
		IHistoryDevice: NewHistoryDeviceService(repo.IHistoryDeviceRepo),
	}
}
