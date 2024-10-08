package main

import (
	"fmt"
	"net/smtp"
)

func main() {
	// Настройки MailHog SMTP
	smtpServer := "localhost"
	smtpPort := "1025"
	from := "from@example.com"
	to := []string{"to@example.com"}
	subject := "Тестовое сообщение"
	body := "Привет, это тестовое сообщение!"

	// Соединяемся с сервером SMTP MailHog
	auth := smtp.PlainAuth("", "", "", smtpServer)

	// Формируем сообщение
	message := []byte("To: " + to[0] + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	// Отправляем сообщение
	err := smtp.SendMail(smtpServer+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Сообщение успешно отправлено!")
}
