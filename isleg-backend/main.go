package main

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// func init() {
// 	// initialize database connection
// 	db, err := config.ConnDB()
// 	if err != nil {
// 		db.Close()
// 		log.Fatal(err)
// 	}

// 	if err := godotenv.Load(); err != nil {
// 		log.Fatal(err)
// 	}
// }

func main() {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		db.Close()
		log.Fatal(err)
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	r := routes.Routes()

	// static file
	os.Mkdir("./uploads", os.ModePerm)
	r.Static("/uploads", "./uploads")

	// test
	// stringDate := "12:00"
	// date, err := time.Parse("2006-01-02 00:00:00 +0000 UTC", stringDate)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(date)

	// run routes
	if err := r.Run(":2406"); err != nil {
		log.Fatal(err)
	}

}
