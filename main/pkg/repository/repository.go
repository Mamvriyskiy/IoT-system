package repository

import (
	"github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/jmoiron/sqlx"
)

//go:generate mockgen -source=repository.go -destination=mocks/mocks.go

type IUserRepo interface {
	CreateUser(user pkg.User) (string, error)
	GetUser(email, password string) (pkg.User, error)
	ChangePassword(password, token string) error
	GetCode(token string) (string, error)
	AddCode(email pkg.Email) error
	GetUserByEmail(email string) (int, error)
	GetAccessLevel(userID, homeID string) (int, error)
}

type IHomeRepo interface {
	CreateHome(home pkg.Home) (string, error)
	DeleteHome(homeID string) error
	UpdateHome(homeID, name string) error
	GetHomeByID(homeID string) (pkg.Home, error)
	ListUserHome(userID string) ([]pkg.Home, error)
}

type IAccessHomeRepo interface {
	AddUser(homeID string, access pkg.Access) (string, error)
	DeleteUser(accessID string) error
	UpdateLevel(accessID string, updateAccess pkg.Access) error
	UpdateStatus(userID string, access pkg.AccessHome) error
	GetListUserHome(homeID string) ([]pkg.ClientHome, error)
	AddOwner(userID, homeID string) (string, error)
	GetInfoAccessByID(accessID string) (pkg.Access, error)
}

type IDeviceRepo interface {
	CreateDevice(homeID string, device pkg.Devices, 
		character pkg.DeviceCharacteristics, typeCharacter pkg.TypeCharacter) (string, error)
	DeleteDevice(deviceID string) error
	GetDeviceByID(deviceID string) (pkg.Devices, error)
	GetListDevices(homeID string) ([]pkg.DevicesInfo, error)
}

type IHistoryDeviceRepo interface {
	CreateDeviceHistory(deviceID string, history pkg.AddHistory) (string, error)
	GetDeviceHistory(deviceID string) ([]pkg.DevicesHistory, error)
}

type Repository struct {
	IUserRepo
	IHomeRepo
	IAccessHomeRepo
	IDeviceRepo
	IHistoryDeviceRepo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		IUserRepo:          NewUserPostgres(db),
		IHomeRepo:          NewHomePostgres(db),
		IAccessHomeRepo:    NewAccessHomePostgres(db),
		IDeviceRepo:        NewDevicePostgres(db),
		IHistoryDeviceRepo: NewDeviceHistoryPostgres(db),
	}
}
