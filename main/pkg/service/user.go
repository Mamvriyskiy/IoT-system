package service

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
	"errors"
	"fmt"
	"net/smtp"

	"github.com/Mamvriyskiy/database_course/main/logger"
	pkg "github.com/Mamvriyskiy/database_course/main/pkg"
	"github.com/Mamvriyskiy/database_course/main/pkg/repository"
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

func (s *UserService) CreateUser(user pkg.UserHandler) (string, error) {
	user.Password = s.generatePasswordHash(user.Password)

	userDB := pkg.UserService{
		User: user.User,
	}

	return s.repo.CreateUser(userDB)
}

func (s *UserService) GetAccessLevel(userID, homeID string) (int, error) {
	return s.repo.GetAccessLevel(userID, homeID)
}

func (s *UserService) CheckUser(user pkg.UserHandler) (pkg.UserData, error) {
	if user.Email == "" {
		logger.Log("Error", "CheckUser", "Empty email:", nil, "")
		return pkg.UserData{}, errors.New("Empty email")
	}

	userAuth, err := s.repo.GetUser(user.Email, user.Password)
	if err != nil {
		logger.Log("Error", "GetUser", "Error get user:", err, user.Username, user.Password)
		return userAuth, err
	}

	return userAuth, err
}

func (s *UserService) GetUserByEmail(email string) (int, error) {
	return s.repo.GetUserByEmail(email)
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

func (s *UserService) SendCode(email pkg.EmailHandler) error {
	safeNum := UseCryptoRandIntn(999999)

	markerAccess, err := generateMarker()

	smtpServer := "localhost"
	smtpPort := "1025"
	from := "IoT-system@mail.ru"
	to := []string{email.Email}
	subject := "Код для восстановления пароля"
	body := fmt.Sprintf("%d\nhttp://localhost:3000/auth/verification?token=%s", safeNum, markerAccess)

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

	//fmt.Println("Письмо отправлено успешно!")

	emailService := pkg.EmailService{
		EmailUser: email.EmailUser,
		Code: safeNum,
		Token: markerAccess,
	}

	err = s.repo.AddCode(emailService)

	return err
}

func (s *UserService) CheckCode(code, token string) error {
	verCode, err := s.repo.GetCode(token)
	// fmt.Println(verCode, code)
	if code != verCode {
		err = errors.New("Negative code")
		logger.Log("Error", "CheckCode", "Negative codeID:", err)
		return err
	}

	return err
}

type tokenClaims struct {
	jwt.StandardClaims
	UserID string `json:"userId"`
}

func (s *UserService) GenerateToken(login, password string) (pkg.UserData, string, error) {
	user, err := s.repo.GetUser(login, s.generatePasswordHash(password))
	if err != nil {
		logger.Log("Error", "GetUser", "Error get user:", err, login, "s.generatePasswordHash(password)")
		return pkg.UserData{}, "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	result, err := token.SignedString([]byte(signingKey))

	return user, result, err
}

func (s *UserService) generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return hex.EncodeToString(hash.Sum([]byte(salt)))
}

func (s *UserService) ParseToken(accessToken string) (string, error) {
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
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		logger.Log("Error", "token.Claims.(*tokenClaims)", "Error token:", nil)
		return "", nil
	}

	return claims.UserID, nil
}
