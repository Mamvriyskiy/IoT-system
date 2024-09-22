package factory

import (
	"github.com/Mamvriyskiy/database_course/main/testdatabase" testdb
)

type ObjectSystem interface {
	InsertObject()
	DeleteObject()
}

func New(typeObject string) ObjectSystem {
	switch typeObject {
	case "user":
		return NewUser()
	}
}



func (u *testdb.TestUser) InsertObject() {

}

func (u *testdb.TestUser) DeleteObject() {

}
