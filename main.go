package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Starting program")
	godotenv.Load()

	/*
	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("Connected to database")
 	*/

	http.HandleFunc("/", Handler)
	envPort := os.Getenv("API_PORT")
	if envPort == "" {
		envPort = "3000"
	}
	port := fmt.Sprintf(":%s", envPort)
	fmt.Println("Listening on port", port)
	http.ListenAndServe(port, nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World3!")
}

func connectDB() (*sqlx.DB, error) {
	dbName := os.Getenv("MINHA_CPA_DB_NAME")
	user := os.Getenv("MINHA_CPA_DB_USER")
	pass := os.Getenv("MINHA_CPA_DB_PASS")
	host := os.Getenv("MINHA_CPA_DB_HOST")
	port := os.Getenv("MINHA_CPA_DB_PORT")

	// dbConn := "user=postgres dbname=yourdatabase sslmode=disable password=yourpassword host=localhost port=5432"

	dbConn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s", user, dbName, pass, host, port)

	if os.Getenv("MINHA_CPA_IS_LOCAL") == "true" {
		dbConn += " sslmode=disable"
	}

	db, err := sqlx.Connect("postgres", dbConn)
	if err != nil {
		return nil, err
	}

	// Test the connection to the database
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully Connected")
	}

	return db, nil
}
