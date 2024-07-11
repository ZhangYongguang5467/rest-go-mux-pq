package main

import (
	"fmt"
	"os"
)

var (
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbName     = os.Getenv("DB_NAME")
	dbHost     = os.Getenv("DB_HOST")
	dbPort     = os.Getenv("DB_PORT")
	dbSSLMode  = os.Getenv("DB_SSLMODE")
)

func main() {
	a := App{}

	fmt.Printf("DB_USER: %s\n", dbUser)
	fmt.Printf("DB_PASSWORD: %s\n", dbPassword)
	fmt.Printf("DB_NAME: %s\n", dbName)
	fmt.Printf("DB_HOST: %s\n", dbHost)
	fmt.Printf("DB_PORT: %s\n", dbPort)
	fmt.Printf("DB_SSLMODE: %s\n", dbSSLMode)
	// use db connection credentials stored in env vars
	a.Initialize(dbUser, dbPassword, dbName, dbHost, dbPort, dbSSLMode)

	a.Run(":8000")
}
