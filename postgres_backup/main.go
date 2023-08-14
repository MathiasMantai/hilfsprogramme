package main

import (
	pg "github.com/habx/pg-commands"
	"os"
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()

	postgresConfig := pg.Postgres {
		Host: "192.168.178.6",
		Port: 5432,
		DB: " ",
		Username: "postgres",
		Password: "postgres",
		EnvPassword: "postgres",
	}

	if postgresConfig.Host == "" || postgresConfig.Port == 0 || postgresConfig.DB == "" || postgresConfig.Username == "" || postgresConfig.EnvPassword == "" {
		fmt.Println("Valid postgres config not found")
		os.Exit(1)
	}

	dump, err := pg.NewDump(&postgresConfig)
	
	if err != nil {
		fmt.Println("Error creating new db dump: " + string(err.Error()))
		os.Exit(2)
	}

	dumpFileName := currentTime.Format("2006-01-02") + "_" + currentTIme.Hour() + "_" + currentTime.Minute() + "_" + currentTime.Second()
	dump.SetFileName(dumpFileName)
	dump.SetPath("./")
	//execute dump
	options := pg.ExecOptions{
		StreamPrint: true,
	}

	res := dump.Exec(options)

	if res.Error != nil {
		fmt.Println("Error executing dump: " + string(err.Error()))
		os.Exit(3)
	} else {
		fmt.Println(res.Output)
	}
}