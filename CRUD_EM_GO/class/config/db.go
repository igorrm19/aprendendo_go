package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func DBconfig() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar dontenv")
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbDatabase := os.Getenv("DB_DATABASE")
	dbPort := os.Getenv("DB_PORT")

	//connetionStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbDatabase)
	connectionStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbDatabase,
	)

	//db, err := sql.Open("pgx", "postgres://pgx_md5:secret@localhost:5432/pgx_test?sslmode=disable")
	dbconnetion, err := sql.Open("pgx", connectionStr)

	if err != nil {
		log.Fatal(err)
	}

	err = dbconnetion.Ping()

	if err != nil {
		log.Fatal(err)
	}

	return dbconnetion

}
