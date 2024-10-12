package repository

import (
	"github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/jmoiron/sqlx"
)

//go:generate mockgen -source=repository.go -destination=mocks/mocks.go

type IUserRepo interface {
	CreateUser(user pkg.UserService) (string, error)
	GetUser(email, password string) (pkg.UserData, error)
	ChangePassword(password, token string) error
	GetCode(token string) (string, error)
	AddCode(email pkg.EmailService) error
	GetUserByEmail(email string) (int, error)
	GetAccessLevel(userID, homeID string) (int, error)
}

type IHomeRepo interface {
	CreateHome(home pkg.HomeService) (string, error)
	DeleteHome(homeID string) error
	UpdateHome(homeID, name string) error
	GetHomeByID(homeID string) (pkg.HomeData, error)
	ListUserHome(userID string) ([]pkg.HomeData, error)
}

type IAccessHomeRepo interface {
	AddUser(homeID string, access pkg.AccessService) (string, error)
	DeleteUser(accessID string) error
	UpdateLevel(accessID string, updateAccess pkg.AccessService) error
	UpdateStatus(userID string, access pkg.AccessService) error
	GetListUserHome(homeID string) ([]pkg.AccessInfoData, error)
	AddOwner(userID, homeID string) (string, error)
	GetInfoAccessByID(accessID string) (pkg.AccessInfoData, error)
}

type IDeviceRepo interface {
	CreateDevice(homeID string, device pkg.DevicesService, 
		character pkg.DeviceCharacteristicsService, typeCharacter pkg.TypeCharacterService) (string, error)
	DeleteDevice(deviceID string) error
	GetDeviceByID(deviceID string) (pkg.DevicesData, error)
	GetListDevices(homeID string) ([]pkg.DevicesData, error)
}

type IHistoryDeviceRepo interface {
	CreateDeviceHistory(deviceID string, history pkg.HistoryService) (string, error)
	GetDeviceHistory(deviceID string) ([]pkg.DevicesHistoryData, error)
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
