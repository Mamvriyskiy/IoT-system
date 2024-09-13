package main

/*
	#include <pthread.h>
	#include <time.h>
	#include <stdio.h>

	static long long getThreadCpuTimeNs() {
		struct timespec t;
		if (clock_gettime(CLOCK_MONOTONIC, &t)) {
			perror("clock_gettime");
			return 0;
		}
		
        //return t.tv_sec * 1000LL + t.tv_nsec / 1000000LL;
        return t.tv_sec * 1000000LL + t.tv_nsec / 1000LL;
	}
*/
import "C"
import (
	"fmt"
	"context"
	"encoding/json"
	"os/exec"
	"github.com/jmoiron/sqlx"
	"github.com/Mamvriyskiy/database_course/main/containers"
)

const (
	START = 10000
	MAXSIZE = 100000
	STEP = 10000
	SEARCHLOGIN = "gfkdnikald"
	REAPEAT = 2000
)

func main() {
	//Запуск контейнера
	dbTestContainers, connDB, err := containers.SetupTestDataBase()

	if err != nil {
		panic(err)
	}
	defer dbTestContainers.Terminate(context.Background())

	//Инициализация файлов 
	err = CreateRandomDataClient()
	if err != nil {
		panic(err)
	}

	//Замер времени
	result, err := goResearch(connDB)
	if err != nil {
		panic(err)
	}

	//Преобразование результатов замера в json
	jsonData, err := json.Marshal(result)
    if err != nil {
        fmt.Println("Ошибка при преобразовании в JSON: %v", err)
    }

	//Запуск скрипта
	cmd := exec.Command("python3", "graph.py", string(jsonData))
    output, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Println("Ошибка при выполнении скрипта: %v", err)
    }

    fmt.Println(string(output))

    fmt.Println("Скрипт успешно запущен и выполнен.")
}

func goResearch(connDB *sqlx.DB) ([3][]int64, error) {
	var mtrResult [3][]int64
	for i := 0; i < 3; i++ {
		mtrResult[i] = make([]int64, MAXSIZE / STEP)
	}

	k := 0
	for i := START; i <= MAXSIZE; i += STEP {
		mtrResult[0][k] = int64(i)
		dataFile := fmt.Sprintf("/mnt/research_data_%d.csv", i)
		err := MigrationsResearchDataBaseUp(connDB, dataFile, "./migrations")
		if err != nil {
			panic(err)
		}

		fmt.Println("Test Size:", i)
		num, _ := searchWithoutIndex(connDB)
		mtrResult[1][k] = num


		err = MigrationsResearchDataBaseUp(connDB, dataFile, "./migrations/indx")
		if err != nil {
			panic(err)
		}

		num, _ = searchIndex(connDB)
		mtrResult[2][k] = num
		fmt.Println()

		err = MigrationsResearchDataBaseDown(connDB)
		if err != nil {
			panic(err)
		}
		k++
	}

	return mtrResult, nil
}

func searchWithoutIndex(connDB *sqlx.DB) (int64, error) {
	var result int64
	for i := 0; i < REAPEAT; i++ {
		start := C.getThreadCpuTimeNs()
		row := connDB.QueryRow(`SELECT login FROM client WHERE login = $1`, SEARCHLOGIN)
		finish := C.getThreadCpuTimeNs()
		result += int64(finish - start)
		
		var searchLogin string
		if err := row.Scan(&searchLogin); err != nil {
			return 0, err
		}

		if SEARCHLOGIN != searchLogin {
			return 0, fmt.Errorf("Negative result", nil)
		}
	}

	fmt.Println("Result time:", result / REAPEAT)
	return result / REAPEAT, nil
}

func searchIndex(connDB *sqlx.DB) (int64, error) {
	var result int64
	for i := 0; i < REAPEAT; i++ {
		start := C.getThreadCpuTimeNs()
		row := connDB.QueryRow(`SELECT login FROM client WHERE login = $1`, SEARCHLOGIN)
		finish := C.getThreadCpuTimeNs()
		result += int64(finish - start)
		
		var searchLogin string
		if err := row.Scan(&searchLogin); err != nil {
			return 0, err
		}

		if SEARCHLOGIN != searchLogin {
			return 0, fmt.Errorf("Negative result", nil)
		}
	}

	fmt.Println("Result time Index:", result / REAPEAT)
	return result / REAPEAT, nil
}
