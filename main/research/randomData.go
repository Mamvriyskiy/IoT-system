package main

import (
	"bufio"
	"os"
)

const (
	START = 1000
	MAXSIZE = 10000
	STEP = 1000
)

// w := bufio.NewWriter(os.Stdout)
// r := bufio.NewReader(os.Stdin)

// name, err := r.ReadString('\\\\n') // считывание до первого переноса строки
// if err != nil {
// 	log.Fatal(err)
// }
// w.WriteString("Привет, " + name) // запись строки в объект w
// w.Flush()

func createRandomData() error {
	for i := START; i <= MAXSIZE; i += STEP {
		nameFile := fmt.Sprintf("research_data_%s", i)
		file, err := os.Create(nameFile)
	}
	//w := bufio.NewWriter()
}


