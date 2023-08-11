package main

import (
	pg "github.com/habx/pg-commands"
	"os"
	"fmt"
)

func main() {
	fmt.Println("T")
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

	dump.SetFileName("postgres_backup_01_01_2023_19_54")
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