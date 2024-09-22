package factory 

import (
	"github.com/Mamvriyskiy/database_course/main/testdatabase" testdb
	"crypto/rand"
	"math/big"
)

const (
	charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	lengthPassword = 12
	lengthUserName = 6
	lengtDomainLeft= 5
	lengtDomainRight = 3
	lengthEmail = 10 
)

type InvoiceUser struct {
	User testdb.TestUser
}

func NewUser() pkg.User {
	var b InvoiceUser

	return b.BuilderUser()
}

func (b *InvoiceUser) BuilderUser() pkg.User {
	b.generatePassword()
	b.generateUserName()
	b.generateEmail()

	return b.User
}

func (b *InvoiceUser) generatePassword() {
	password := make([]byte, lengthPassword)
	for i := range password {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		password[i] = charset[n.Int64()]
	}

	b.Password = password
}

func (b *InvoiceUser) generateUserName() {
	userName := make([]byte, lengthUserName)
	for i := range userName {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		userName[i] = charset[n.Int64()]
	}

	b.Username = userName
}

func (b *InvoiceUser) generateEmail() {
	email := make([]byte, lengthEmail + lengtDomainLeft + lengtDomainRight + 2)
	i := 0
	for ; i < lengthEmail; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		email[i] = charset[n.Int64()]
	}

	email[i] = '@'
	i++

	for ; i < i + lengtDomainLeft; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		email[i] = charset[n.Int64()]
	}

	email[i] = '.'
	i++

	for ; i < i + lengtDomainRight; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		email[i] = charset[n.Int64()]
	}
	
	b.Email = emil
}
