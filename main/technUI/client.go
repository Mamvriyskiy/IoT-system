package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	pkg "git.iu7.bmstu.ru/mis21u869/PPO/-/tree/lab3/pkg"
)

type UserResponse struct {
	Token    string `json:"token"`
	HomeName string `json:"name"`
	ID       int    `json:"id"`
	HomeID   int    `json:"homeId"`
}

func menu() {
	fmt.Fprintf(os.Stdout, `
		1) Создать дом
		2) Удалить дом
		3) Обновить имя дома
		4) Добавить участника
		5) Удалить участника
		6) Добавить устройство
		7) Удалить устройство
		8) Просмотреть статистику
		9) Запустить устройство
		0) Завершить работу
`)
}

func checkStat(token string) error {
	var device string
	reader := bufio.NewReader(os.Stdin)
	fmt.Fprintf(os.Stdout, "Введите имя устройства: ")
	_, err := fmt.Fscanf(reader, "%s\n", &device)
	if err != nil {
		return err
	}

	data := pkg.AddHistory{
		Name: device,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = requestServer("GET", "/api/history", jsonData, token)
	if err != nil {
		return err
	}

	return nil
}

func runDevice(token string) error {
	var device string
	reader := bufio.NewReader(os.Stdin)
	fmt.Fprintf(os.Stdout, "Введите имя устройства: ")
	_, err := fmt.Fscanf(reader, "%s\n", &device)
	if err != nil {
		return err
	}

	data := pkg.AddHistory{
		Name: device,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = requestServer("POST", "/api/history", jsonData, token)
	if err != nil {
		return err
	}

	return nil
}

func deleteUser(token string) error {
	var email string
	reader := bufio.NewReader(os.Stdin)
	fmt.Fprintf(os.Stdout, "Введите почту пользователя: ")
	_, err := fmt.Fscanf(reader, "%s\n", &email)
	if err != nil {
		return err
	}

	data := pkg.AddUserHome{
		Email: email,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = requestServer("POST", "/api/access", jsonData, token)
	if err != nil {
		return err
	}

	return nil
}

func addUser(token string) error {
	var email string
	reader := bufio.NewReader(os.Stdin)
	fmt.Fprintf(os.Stdout, "Введите почту пользователя: ")
	_, err := fmt.Fscanf(reader, "%s\n", &email)
	if err != nil {
		return err
	}

	var level int
	fmt.Fprintf(os.Stdout, "Введите уровень доступа: ")
	_, err = fmt.Fscanf(reader, "%d\n", &level)
	if err != nil {
		return err
	}

	data := pkg.AddUserHome{
		Email:       email,
		AccessLevel: level,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = requestServer("POST", "/api/access", jsonData, token)
	if err != nil {
		return err
	}

	return nil
}

func deleteDevice(token string) error {
	var deviceName string
	reader := bufio.NewReader(os.Stdin)
	fmt.Fprintf(os.Stdout, "Введите имя устройства: ")
	_, err := fmt.Fscanf(reader, "%s\n", &deviceName)
	if err != nil {
		return err
	}

	data := pkg.Devices{
		Name: deviceName,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = requestServer("DELETE", "/api/device", jsonData, token)
	if err != nil {
		return err
	}

	return nil
}

func addDevice(token string) error {
	var deviceName string
	reader := bufio.NewReader(os.Stdin)
	fmt.Fprintf(os.Stdout, "Введите имя устройства: ")
	_, err := fmt.Fscanf(reader, "%s\n", &deviceName)
	if err != nil {
		return err
	}

	data := pkg.Devices{
		Name:             deviceName,
		TypeDevice:       "vacuum cleaner",
		Brand:            "apple",
		Status:           "wait",
		PowerConsumption: 100,
		MinParameter:     10,
		MaxParameter:     30,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = requestServer("POST", "/api/device", jsonData, token)
	if err != nil {
		return err
	}

	return nil
}

func registerUser() error {
	var password, username, email string
	reader := bufio.NewReader(os.Stdin)
	fmt.Fprintf(os.Stdout, "Введите имя пользователя: ")
	_, err := fmt.Fscanf(reader, "%s\n", &username)
	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stdout, "Введите почту: ")
	_, err = fmt.Fscanf(reader, "%s\n", &email)
	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stdout, "Введите пароль: ")
	_, err = fmt.Fscanf(reader, "%s\n", &password)
	if err != nil {
		return err
	}

	data := map[string]string{
		"password": password,
		"login":    username,
		"email":    email,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = requestServer("POST", "/auth/sign-up", jsonData, "")
	if err != nil {
		return err
	}

	return nil
}

func auth() (string, error) {
	var password, username string
	reader := bufio.NewReader(os.Stdin)
	fmt.Fprintf(os.Stdout, "Введите имя пользователя: ")
	_, err := fmt.Fscanf(reader, "%s\n", &username)
	if err != nil {
		return "", err
	}

	fmt.Fprintf(os.Stdout, "Введите пароль: ")
	_, err = fmt.Fscanf(reader, "%s\n", &password)
	if err != nil {
		return "", err
	}

	data := map[string]string{
		"password": password,
		"login":    username,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	result, err := requestServer("POST", "/auth/sign-in", jsonData, "")
	if err != nil {
		return "", err
	}

	return result, nil
}

func createHome(token string) error {
	var nameHome string
	reader := bufio.NewReader(os.Stdin)
	fmt.Fprintf(os.Stdout, "Введите имя дома: ")
	_, err := fmt.Fscanf(reader, "%s\n", &nameHome)
	if err != nil {
		return err
	}

	data := map[string]string{
		"name": nameHome,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = requestServer("POST", "/api/home", jsonData, token)
	if err != nil {
		return err
	}

	return nil
}

func deleteHome(token string) error {
	data := map[string]string{
		// "homeId": string(user.ID),
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = requestServer("DELETE", "/api/home", jsonData, token)
	if err != nil {
		return err
	}

	return nil
}

func updateHome(token string) error {
	var newHomeName string
	reader := bufio.NewReader(os.Stdin)
	fmt.Fprintf(os.Stdout, "Введите имя дома: ")
	fmt.Fscan(reader, &newHomeName)

	data := map[string]string{
		"name": newHomeName,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = requestServer("PUT", "/api/home", jsonData, token)
	if err != nil {
		return err
	}

	return nil
}

// func getListHomeByUser() error {
// 	data := map[string]string{}

// 	jsonData, err := json.Marshal(data)
// 	if err != nil {
// 		return err
// 	}

// 	_, err = requestServer("GET", "/api/home", jsonData)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func requestServer(typeReq, path string, jsonData []byte, token string) (string, error) {
	ctx := context.Background()
	req, err := http.NewRequestWithContext(ctx, typeReq,
		"http://localhost:8000"+path, bytes.NewBuffer(jsonData))

	// req, err := http.NewRequest(typeReq, "http://localhost:8000"+path, bytes.NewBuffer(jsonData))
	req.Header.Set("Authorization", "Bearer "+token)
	if err != nil {
		return "", err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var responseBody []byte
	if resp.StatusCode == http.StatusOK {
		responseBody, err = io.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		fmt.Fprint(os.Stdout, "Тело ответа:", string(responseBody))
	} else {
		fmt.Fprint(os.Stdout, "Сервер вернул статус:", resp.Status)
	}

	var userResponse UserResponse
	if len(responseBody) > 0 {
		if err := json.Unmarshal(responseBody, &userResponse); err != nil {
			return "", err
		}
	}

	return userResponse.Token, nil
}

func clientGO(token string) {
	for {
		menu()

		var num int
		reader := bufio.NewReader(os.Stdin)
		fmt.Fprint(os.Stdout, "Выберите номер: ")
		_, err := fmt.Fscanf(reader, "%d\n", &num)
		if err != nil {
			continue
		}
		handleOption(num, token)
	}
}

func handleOption(num int, token string) {
	switch {
	case num < 4 && num > 0:
		handleHome(num, token)
	case num >= 6 && num <= 9:
		handleDevice(num, token)
	case num == 4 || num == 5:
		handleAccess(num, token)
	default:
		break
	}
}

func handleAccess(num int, token string) {
	switch num {
	case 4:
		handleAction(addUser, token)
	case 5:
		handleAction(deleteUser, token)
	}
}

func handleDevice(num int, token string) {
	switch num {
	case 6:
		handleAction(addDevice, token)
	case 7:
		handleAction(deleteDevice, token)
	case 8:
		handleAction(checkStat, token)
	case 9:
		handleAction(runDevice, token)
	}
}

func handleHome(num int, token string) {
	switch num {
	case 1:
		handleAction(createHome, token)
	case 2:
		handleAction(deleteHome, token)
	case 3:
		handleAction(updateHome, token)
	}
}

func handleAction(action func(string) error, token string) {
	if err := action(token); err != nil {
		return
	}
}

func main() {
	var err error
	var token string
	fmt.Fprintln(os.Stdout, `
		Меню:
		1)Регистрация
		2)Войти
	`)
	var num int
	fmt.Scan(&num)
	switch num {
	case 1:
		err = registerUser()
		if err != nil {
			return
		}
	case 2:
		token, err = auth()
		if err != nil {
			return
		}
	}

	if err == nil {
		clientGO(token)
	} else {
		return
	}
}
