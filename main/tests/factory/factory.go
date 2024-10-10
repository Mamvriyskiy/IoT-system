package factory

import (
	method "github.com/Mamvriyskiy/database_course/main/tests/method"
	"github.com/jmoiron/sqlx"
)

type ObjectSystem interface {
	InsertObject(connDB *sqlx.DB) (string, error)
	//DeleteObject()
}

func New(typeObject, keyСharacter, typeСharacter string) ObjectSystem {
	switch typeObject {
	case "service":
		return NewService(typeObject, keyСharacter)
	case "DB":
		return NewDB(typeObject, keyСharacter)
	default:
		return nil
	} 
}

func NewDB(typeObject, keyСharacter string) {
	switch typeObject {
	case "user":
		return method.NewUserDB(keyСharacter)
	case "email":
		return method.NewEmailDB(keyСharacter)
	case "home":
		return method.NewHomeDB()
	case "access":
		return method.NewAccessDB()
	case "device":
		return method.NewDeviceDB()
	case "character":
		return method.NewCharacterDB()
	case "history":
		return method.NewHistoryDB()
	default:
		return nil
	} 
}
 
func NewService(typeObject, keyСharacter string) {
	switch typeObject {
	case "user":
		return method.NewUserService(keyСharacter)
	case "email":
		return method.NewEmailService(keyСharacter)
	case "home":
		return method.NewHomeService()
	case "access":
		return method.NewAccessService()
	case "device":
		return method.NewDeviceService()
	case "character":
		return method.NewCharacterService()
	case "history":
		return method.NewHistoryService()
	default:
		return nil
	} 
}
