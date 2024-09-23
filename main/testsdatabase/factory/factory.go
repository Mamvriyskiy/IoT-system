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
	default:
		return nil
	} 
}

func (u TestUser) InsertObject() {

}

func (u TestUser) DeleteObject() {

}
