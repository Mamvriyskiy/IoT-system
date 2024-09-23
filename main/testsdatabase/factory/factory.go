package factory

import (
	"github.com/Mamvriyskiy/database_course/main/testsdatabase/builder"
)

type ObjectSystem interface {
	InsertObject()
	DeleteObject()
}

func New(typeObject string) ObjectSystem {
	switch typeObject {
	case "user":
		return builder.NewUser()
	default:
		return nil
	} 
}
