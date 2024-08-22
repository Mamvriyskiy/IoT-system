package service

import (
	"github.com/Mamvriyskiy/DBCourse/main/pkg"
	"github.com/Mamvriyskiy/DBCourse/main/pkg/repository"
)

type IUser interface {
	CreateUser(user pkg.User) (int, error)
	CheckUser(user pkg.User) (int, error)
	GenerateToken(login, password string) (string, error)
	ParseToken(token string) (int, error)
	ChangePassword(password, token string) error
	CheckCode(code, token string) error
	SendCode(email pkg.Email) error
}

type IHome interface {
	CreateHome(idUser int, home pkg.Home) (int, error)
	DeleteHome(homeID int, homeName string) error
	UpdateHome(home pkg.UpdateNameHome) error
	GetHomeByID(homeID int) (pkg.Home, error)
	ListUserHome(userID int) ([]pkg.Home, error)
}

type IAccessHome interface {
	AddUser(userID int, access pkg.Access) (int, error)
	DeleteUser(idUser int, access pkg.Access) error
	UpdateLevel(idUser int, updateAccess pkg.Access) error
	UpdateStatus(idUser int, access pkg.AccessHome) error
	GetListUserHome(homeID int) ([]pkg.ClientHome, error)
	AddOwner(userID, homeID int) (int, error)
}

type IDevice interface {
	CreateDevice(homeID int, device *pkg.Devices) (int, error)
	DeleteDevice(idDevice int, name string) error
	GetDeviceByID(deviceID int) (pkg.Devices, error)
	GetListDevices(userID int) ([]pkg.Devices, error)
}

type IHistoryDevice interface {
	CreateDeviceHistory(deviceID int, history pkg.AddHistory) (int, error)
	GetDeviceHistory(userID int, name string) ([]pkg.DevicesHistory, error)
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
