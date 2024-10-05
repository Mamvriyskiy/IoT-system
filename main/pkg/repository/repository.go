package repository

import (
	"github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/jmoiron/sqlx"
)

//go:generate mockgen -source=repository.go -destination=mocks/mocks.go

type IUserRepo interface {
	CreateUser(user pkg.User) (int, error)
	GetUser(email, password string) (pkg.User, error)
	ChangePassword(password, token string) error
	GetCode(token string) (string, error)
	AddCode(email pkg.Email) error
	GetUserByEmail(email string) (int, error)
	GetAccessLevel(userID int, homeID string) (int, error)
}

type IHomeRepo interface {
	CreateHome(home pkg.Home) (int, error)
	DeleteHome(homeID string) error
	UpdateHome(homeID, name string) error
	GetHomeByID(homeID string) (pkg.Home, error)
	ListUserHome(userID int) ([]pkg.Home, error)
}

type IAccessHomeRepo interface {
	AddUser(homeID string, access pkg.Access) (int, error)
	DeleteUser(accessID string) error
	UpdateLevel(accessID string, updateAccess pkg.Access) error
	UpdateStatus(userID int, access pkg.AccessHome) error
	GetListUserHome(homeID string) ([]pkg.ClientHome, error)
	AddOwner(userID int, homeID string) (int, error)
	GetInfoAccessByID(accessID string) (pkg.Access, error)
}

type IDeviceRepo interface {
	CreateDevice(homeID string, device pkg.Devices, 
		character pkg.DeviceCharacteristics, typeCharacter pkg.TypeCharacter) (int, error)
	DeleteDevice(deviceID string) error
	GetDeviceByID(deviceID string) (pkg.Devices, error)
	GetListDevices(homeID string) ([]pkg.DevicesInfo, error)
}

type IHistoryDeviceRepo interface {
	CreateDeviceHistory(deviceID string, history pkg.AddHistory) (int, error)
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
