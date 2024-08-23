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
		
        // return t.tv_sec * 1000LL + t.tv_nsec / 1000000LL;
        return t.tv_sec * 1000000LL + t.tv_nsec / 1000LL;
	}
*/
import "C"
import (
	"fmt"
	"context"
	//"time"
	"github.com/jmoiron/sqlx"
	"github.com/Mamvriyskiy/database_course/main/containers"
)

const (
	START = 10000
	MAXSIZE = 100000
	STEP = 10000
	SEARCHLOGIN = "gfkdnikald"
)

func main() {
	dbTestContainers, connDB, err := containers.SetupTestDataBase()

	if err != nil {
		panic(err)
	}
	defer dbTestContainers.Terminate(context.Background())

	err = CreateRandomDataClient()
	if err != nil {
		panic(err)
	}

	err = goResearch(connDB)
	if err != nil {
		panic(err)
	}
}

func goResearch(connDB *sqlx.DB) error {
	for i := START; i <= MAXSIZE; i += STEP {
		dataFile := fmt.Sprintf("/mnt/research_data_%d.csv", i)
		err := MigrationsResearchDataBaseUp(connDB, dataFile, "./migrations")
		if err != nil {
			panic(err)
		}

		fmt.Println("Test Size:", i)
		searchWithoutIndex(connDB)

		err = MigrationsResearchDataBaseUp(connDB, dataFile, "./migrations/indx")
		if err != nil {
			panic(err)
		}

		searchIndex(connDB)
		fmt.Println()

		err = MigrationsResearchDataBaseDown(connDB)
		if err != nil {
			panic(err)
		}
	}

	return nil
}

func searchWithoutIndex(connDB *sqlx.DB) error {
	var result int64
	for i := 0; i < 30; i++ {
		start := C.getThreadCpuTimeNs()
		row := connDB.QueryRow(`SELECT login FROM client WHERE login = $1`, SEARCHLOGIN)
		finish := C.getThreadCpuTimeNs()
		result += int64(finish - start)
		
		var searchLogin string
		if err := row.Scan(&searchLogin); err != nil {
			return err
		}

		if SEARCHLOGIN != searchLogin {
			return fmt.Errorf("Negative result", nil)
		}
	}

	fmt.Println("Result time:", result / 30)
	return nil
}

func searchIndex(connDB *sqlx.DB) error {
	var result int64
	for i := 0; i < 100; i++ {
		start := C.getThreadCpuTimeNs()
		row := connDB.QueryRow(`SELECT login FROM client WHERE login = $1`, SEARCHLOGIN)
		finish := C.getThreadCpuTimeNs()
		result += int64(finish - start)
		
		var searchLogin string
		if err := row.Scan(&searchLogin); err != nil {
			return err
		}

		if SEARCHLOGIN != searchLogin {
			return fmt.Errorf("Negative result", nil)
		}
	}

	fmt.Println("Result time Index:", result / 100)
	return nil
}
