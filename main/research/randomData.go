package main

import (
	"bufio"
	"os"
	"fmt"
	"math/rand"
	"time"
)

func CreateRandomDataClient() error {
	for i := START; i <= MAXSIZE; i += STEP {
		nameFile := fmt.Sprintf("./mnt/research_data_%d.csv", i)
		file, err := os.Create(nameFile)
		if err != nil {
			return err
		}

		w := bufio.NewWriter(file)
		w.WriteString("password,login,email\n")
		fmt.Println(nameFile, "created")
		for j := 0; j < i - 1; j++ {
			password := randStringRunes(15)
			login := randStringRunes(10)
			email := randStringRunes(10)
			w.WriteString(fmt.Sprintf("%s,%s,%s\n", password, login, email))
		}	
		
		w.WriteString(fmt.Sprintf("%s,%s,%s\n", "sfhkhtflkhbnghj", SEARCHLOGIN, "emadinaarl"))
		fmt.Println(nameFile, "filled")

		w.Flush()
	}
	
	return nil
}

func init() {
    rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}

