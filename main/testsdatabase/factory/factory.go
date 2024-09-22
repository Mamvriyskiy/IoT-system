package factory

import (

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
