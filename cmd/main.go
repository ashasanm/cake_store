package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"cake_store/routes"
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
)

func connectDB() *sql.DB {
	var dataSourceName = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	var driverName = os.Getenv("DB_DRIVER")

	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connected to database!")
	return db
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	// Initiate database
	db := connectDB()
	defer db.Close()

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	// use this gin new to use custom middleware logger
	// Uncomment gin default to user default gin recover and logger and comment gin new, and user
	// router := gin.New()
	// router.Use(gin.Recovery())                 // to recover gin automatically
	// router.Use(handler.JsonLoggerMiddleware())

	routes.AddCakeRoutes(router)
	router.Run() // listen and serve on 0.0.0.0:8080
}
