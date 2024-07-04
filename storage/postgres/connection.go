package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Config() {
	err := godotenv.Load("storage/postgres/.env")
	if err != nil {
		log.Fatal("Xatolik .env fileda?", err)
	}
}

func Connection() (*sql.DB, error) {
	conn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s sslmode=disable", os.Getenv("HOST"), os.Getenv("USER"), os.Getenv("DBNAME"), os.Getenv("PORT"), os.Getenv("PASSWORD"))

	fmt.Println( os.Getenv("HOST"), os.Getenv("USER"), os.Getenv("DBNAME"), os.Getenv("PORT"), os.Getenv("PASSWORD"))

	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal("Postgres connectionda xatolik?", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Xatolik PingPong da?", err)
		return nil, err
	}

	return db, nil
}