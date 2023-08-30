package main

import (
	pg "github.com/habx/pg-commands"
	"os"
	"fmt"
	"time"
	"github.com/joho/godotenv"
)

func main() {
	currentTime := time.Now()
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading environment file")
		os.Exit(1)
	}

	postgresConfig := pg.Postgres {
		Host: os.Getenv("POSTGRES_HOST"),
		Port: os.Getenv("POSTGRES_PORT"),
		DB: os.Getenv("POSTGRES_DATABASE"),
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		EnvPassword: os.Getenv("POSTGRES_ENV_PASSWORD"),
	}

	if postgresConfig.Host == "" || postgresConfig.Port == 0 || postgresConfig.DB == "" || postgresConfig.Username == "" || postgresConfig.EnvPassword == "" {
		fmt.Println("Valid postgres config not found")
		os.Exit(2)
	}

	dump, err := pg.NewDump(&postgresConfig)
	
	if err != nil {
		fmt.Println("Error creating new db dump: " + string(err.Error()))
		os.Exit(3)
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