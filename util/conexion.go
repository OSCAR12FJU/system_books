package util

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func CreateConnection() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("No detectado")
	}

	// dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// fmt.Printf("DB_HOST: %s\n", dbHost)
	fmt.Printf("DB_PORT: %s\n", dbPort)
	fmt.Printf("DB_USER: %s\n", dbUser)
	fmt.Printf("DB_PASSWORD: %s\n", dbPassword)
	fmt.Printf("DB_NAME: %s\n", dbName)

	portAsNumber, _ := strconv.Atoi(dbPort)
	conectStrin := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", portAsNumber, dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", conectStrin)
	if err != nil {
		log.Fatal("No se conecto", err)
	}
	err = db.Ping()
	if err != nil {
		log.Printf("Error de conexión %v", err)
	}

	fmt.Println("Conexion establecida")
	return db, nil
}
