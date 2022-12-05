package contextmanager

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitEnv(fileName string) error {
	err := godotenv.Load(fileName)
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
		return err
	}
	return nil
}

func OpenDBConnection() *sql.DB {
	// load env variable
	err := InitEnv(".env")
	if err != nil {
		panic(err)
	}

	var dataSourceName = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	var driverName = os.Getenv("DB_DRIVER")

	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}
	return db
}
