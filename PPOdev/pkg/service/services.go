package service

import (
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg/repository"
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg/repositoryCH"
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
	DeleteHome(homeID int) error
	UpdateHome(home pkg.Home) error
	GetHomeByID(homeID int) (pkg.Home, error)
	ListUserHome(userID int) ([]pkg.Home, error)
}

type IAccessHome interface {
	AddUser(userID, accessLevel int, email string) (int, error)
	DeleteUser(idUser int, email string) error
	UpdateLevel(idUser int, access pkg.AddUserHome) error
	UpdateStatus(idUser int, access pkg.AccessHome) error
	GetListUserHome(homeID int) ([]pkg.ClientHome, error)
	AddOwner(userID, homeID int) (int, error)
}

type IDevice interface {
	CreateDevice(homeID int, device *pkg.Devices) (int, error)
	DeleteDevice(idDevice int, name string) error
	GetDeviceByID(deviceID int) (pkg.Devices, error)
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

func NewServicesCH(repo *repositoryCH.Repository) *Services {
	return &Services{
		IUser:          NewUserService(repo.IUserRepo),
		IHome:          NewHomeService(repo.IHomeRepo),
		IAccessHome:    NewAccessHomeService(repo.IAccessHomeRepo),
		IDevice:        NewDeviceService(repo.IDeviceRepo),
		IHistoryDevice: NewHistoryDeviceService(repo.IHistoryDeviceRepo),
	}
}
