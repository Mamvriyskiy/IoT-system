package service

import (
	"github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/Mamvriyskiy/database_course/main/pkg/repository"
)

type IUser interface {
	CreateUser(user pkg.UserHandler) (string, error)
	CheckUser(user pkg.UserHandler) (pkg.UserData, error)
	GenerateToken(login, password string) (pkg.UserData, string, error)
	ParseToken(token string) (string, error)
	ChangePassword(password, token string) error
	CheckCode(code, token string) error
	SendCode(email pkg.EmailHandler) error
	GetUserByEmail(email string) (int, error)
	GetAccessLevel(userID, homeID string) (int, error)
}

type IHome interface {
	CreateHome(home pkg.HomeHandler) (pkg.HomeData, error)
	DeleteHome(homeID string) error
	UpdateHome(homeID, name string) error
	GetHomeByID(homeID string) (pkg.HomeData, error)
	ListUserHome(userID string) ([]pkg.HomeData, error)
}

type IAccessHome interface {
	AddUser(homeID string, access pkg.AccessHandler) (string, error)
	DeleteUser(accessID string) error
	UpdateLevel(accessID string, updateAccess pkg.AccessHandler) error
	UpdateStatus(userID string, access pkg.AccessHandler) error
	GetListUserHome(homeID string) ([]pkg.AccessInfoData, error)
	AddOwner(userID, homeID string) (string, error)
	GetInfoAccessByID(accessID string) (pkg.AccessInfoData, error)
}

type IDevice interface {
	CreateDevice(homeID string, device pkg.DevicesHandler) (pkg.DevicesData, error)
	DeleteDevice(deviceID string) error
	GetDeviceByID(deviceID string) (pkg.DevicesData, error)
	GetListDevices(homeID string) ([]pkg.DevicesData, error)
	GetInfoDevice(deviceID string) (pkg.DevicesData, error)
}

type IHistoryDevice interface {
	CreateDeviceHistory(deviceID string) (string, error)
	GetDeviceHistory(deviceID string) ([]pkg.DevicesHistoryData, error)
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
