package method 

import (
	"crypto/rand"
	"math/big"
	"github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/jmoiron/sqlx"
)

const (
	lengthToken = 25
	lengthCode = 6
)

type TestEmail struct {
	pkg.EmailService
}

func NewEmailDB(email string) *TestEmail {
	var b TestEmail

	return b.BuilderEmail(email)
}

func (b *TestEmail) BuilderEmail(email string) *TestEmail {
	b.generateCode()
	b.generateToken()
	if email == "" {
		b.generateEmail()
	} else {
		b.Email = email
	}

	return b
}

func (b *TestEmail) generateCode() {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(lengthCode)))
	if err != nil {
		b.Code = 111111
	}
	b.Code = int(n.Int64())
}

func (b *TestEmail) generateEmail() {
	b.Email = createEmail(lengthEmail, lengtDomainLeft, lengtDomainRight)
}

func (b *TestEmail) generateToken() {
	token := make([]byte, lengthToken)
	for j := 0; j < lengthToken; j++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		token[j] = charset[n.Int64()]
	}

	b.Token = string(token)
}

func (tu TestEmail) InsertObject(connDB *sqlx.DB) (string, error) {
	return "", nil
}
