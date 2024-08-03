package repository

import (
	pkg "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"github.com/jmoiron/sqlx"
)

//go:generate mockgen -source=repository.go -destination=mocks/mocks.go

type IUserRepo interface {
	CreateUser(user pkg.User) (int, error)
	GetUser(login, password string) (pkg.User, error)
	ChangePassword(password, token string) error
	GetCode(token string) (string, error)
	AddCode(email pkg.Email) error
}

type IHomeRepo interface {
	CreateHome(idUser int, home pkg.Home) (int, error)
	DeleteHome(homeID int) error
	UpdateHome(home pkg.Home) error
	GetHomeByID(homeID int) (pkg.Home, error)
	ListUserHome(userID int) ([]pkg.Home, error)
}

type IAccessHomeRepo interface {
	AddUser(userID, accessLevel int, email string) (int, error)
	DeleteUser(idUser int, email string) error
	UpdateLevel(idUser int, access pkg.AddUserHome) error
	UpdateStatus(idUser int, access pkg.AccessHome) error
	GetListUserHome(idHome int) ([]pkg.ClientHome, error)
	AddOwner(userID, homeID int) (int, error)
}

type IDeviceRepo interface {
	CreateDevice(homeID int, device *pkg.Devices) (int, error)
	DeleteDevice(idDevice int, name string) error
	GetDeviceByID(deviceID int) (pkg.Devices, error)
}

type IHistoryDeviceRepo interface {
	CreateDeviceHistory(userID int, history pkg.AddHistory) (int, error)
	GetDeviceHistory(userID int, name string) ([]pkg.DevicesHistory, error)
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
