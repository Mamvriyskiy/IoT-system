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
}

type IHomeRepo interface {
	CreateHome(home pkg.Home) (int, error)
	DeleteHome(homeID int, homeName string) error
	UpdateHome(home pkg.UpdateNameHome) error
	GetHomeByID(homeID int) (pkg.Home, error)
	ListUserHome(userID int) ([]pkg.Home, error)
}

type IAccessHomeRepo interface {
	AddUser(userID int, access pkg.Access) (int, error)
	DeleteUser(idUser int, access pkg.Access) error
	UpdateLevel(userID int, updateAccess pkg.Access) error
	UpdateStatus(idUser int, access pkg.AccessHome) error
	GetListUserHome(userID int) ([]pkg.ClientHome, error)
	AddOwner(userID, homeID int) (int, error)
}

type IDeviceRepo interface {
	CreateDevice(userID int, device *pkg.Devices, 
		character pkg.DeviceCharacteristics, typeCharacter pkg.TypeCharacter) (int, error)
	DeleteDevice(idDevice int, name, home string) error
	GetDeviceByID(deviceID int) (pkg.Devices, error)
	GetListDevices(userID int) ([]pkg.Devices, error)
}

type IHistoryDeviceRepo interface {
	CreateDeviceHistory(userID int, history pkg.AddHistory) (int, error)
	GetDeviceHistory(userID int, name, home string) ([]pkg.DevicesHistory, error)
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
