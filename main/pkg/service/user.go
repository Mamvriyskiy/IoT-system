package service

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
	"errors"
	"fmt"
	"net/smtp"

	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/logger"
	pkg "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
	"git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg/repository"
	jwt "github.com/dgrijalva/jwt-go"
)

const (
	salt       = "hfdjmaxckdk20"
	signingKey = "jaskljfkdfndnznmckmdkaf3124kfdlsf"
)

type UserService struct {
	repo repository.IUserRepo
}

func NewUserService(repo repository.IUserRepo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) ChangePassword(password, token string) error {
	password = s.generatePasswordHash(password)
	return s.repo.ChangePassword(password, token)
}

func (s *UserService) CreateUser(user pkg.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *UserService) CheckUser(user pkg.User) (id int, err error) {
	if user.Email == "" {
		logger.Log("Error", "CheckUser", "Empty email:", nil, "")
		return -1, err
	}

	user, err = s.repo.GetUser(user.Username, user.Password)
	if err != nil {
		logger.Log("Error", "GetUser", "Error get user:", err, user.Username, user.Password)
		return id, err
	}

	return user.ID, err
}

type markerClaims struct {
	jwt.StandardClaims
}

func generateMarker() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &markerClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(3 * time.Minute).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	return token.SignedString([]byte(signingKey))
}

func (s *UserService) SendCode(email pkg.Email) error {
	safeNum := UseCryptoRandIntn(999999)
	email.Code = safeNum

	markerAccess, err := generateMarker()

	smtpServer := "localhost"
	smtpPort := "1025"
	from := "IoT-system@mail.ru"
	to := []string{email.Email}
	subject := "Код для восстановления пароля"
	body := fmt.Sprintf("%d\nhttp://localhost:8000/auth/checkcode?token=%s", safeNum, markerAccess)

	// Соединяемся с сервером SMTP MailHog
	auth := smtp.PlainAuth("", "", "", smtpServer)

	// Формируем сообщение
	message := []byte("To: " + to[0] + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	// Отправляем сообщение
	err = smtp.SendMail(smtpServer+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Письмо отправлено успешно!")
	email.Token = markerAccess
	err = s.repo.AddCode(email)

	return err
}

func (s *UserService) CheckCode(code, token string) error {
	verCode, err := s.repo.GetCode(token)
	fmt.Println(verCode, code)
	if code != verCode {
		err = errors.New("Negative code")
		logger.Log("Error", "CheckCode", "Negative codeID:", err)
		return err
	}

	return err
}

type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"userId"`
}

func (s *UserService) GenerateToken(login, password string) (string, error) {
	user, err := s.repo.GetUser(login, s.generatePasswordHash(password))
	if err != nil {
		logger.Log("Error", "GetUser", "Error get user:", err, login, "s.generatePasswordHash(password)")
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *UserService) generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return hex.EncodeToString(hash.Sum([]byte(salt)))
}

func (s *UserService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				logger.Log("Error", "token.Method.(*jwt.SigningMethodHMAC)", "Error jwt token:", nil, "")
				return 0, nil
			}

			return []byte(signingKey), nil
		})
	if err != nil {
		logger.Log("Error", "jwt.ParseWithClaims", "Error parse token:", err, accessToken)
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		logger.Log("Error", "token.Claims.(*tokenClaims)", "Error token:", nil)
		return 0, nil
	}

	return claims.UserID, nil
}
