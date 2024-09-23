package factory 

import (
	"crypto/rand"
	"math/big"
	"github.com/Mamvriyskiy/database_course/main/pkg"
)

const (
	charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	lengthPassword = 12
	lengthUserName = 6
	lengtDomainLeft= 5
	lengtDomainRight = 3
	lengthEmail = 10 
)

type TestUser struct {
	pkg.User
}

func NewUser() *TestUser {
	var b TestUser

	return b.BuilderUser()
}

func (b *TestUser) BuilderUser() *TestUser {
	b.generatePassword()
	b.generateUserName()
	b.generateEmail()

	return b
}

func (b *TestUser) generatePassword() {
	password := make([]byte, lengthPassword)
	for i := range password {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		password[i] = charset[n.Int64()]
	}

	b.Password = string(password)
}

func (b *TestUser) generateUserName() {
	userName := make([]byte, lengthUserName)
	for i := range userName {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		userName[i] = charset[n.Int64()]
	}

	b.Username = string(userName)
}

func (b *TestUser) generateEmail() {
	email := make([]byte, lengthEmail + lengtDomainLeft + lengtDomainRight + 2)
	i := 0
	for j := 0; j < lengthEmail; j++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		email[i] = charset[n.Int64()]
		i++
	}

	email[i] = '@'
	i++

	for j := 0; j < lengtDomainLeft; j++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		email[i] = charset[n.Int64()]
		i++
	}

	email[i] = '.'
	i++

	for j := 0; j < lengtDomainRight; j++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		email[i] = charset[n.Int64()]
		i++
	}
	
	b.Email = string(email)
}

func (u TestUser) InsertObject() {

}

func (u TestUser) DeleteObject() {

}
