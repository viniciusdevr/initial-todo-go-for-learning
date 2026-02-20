package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func SetupDB() *sql.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error to connect with .env")
	}

	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DB_USER")
	cfg.Passwd = os.Getenv("DB_PASSWORD")
	cfg.Net = "tcp"
	cfg.Addr = fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
	cfg.DBName = os.Getenv("DB_NAME")

	connectionString := cfg.FormatDSN()

	dbConnection, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal("Database connection error")
	}

	dbConnection.SetMaxOpenConns(10)
	dbConnection.SetMaxIdleConns(10)
	dbConnection.SetConnMaxLifetime(time.Minute * 3)

	err = dbConnection.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Sucessfully connected do database")
	return dbConnection
}
