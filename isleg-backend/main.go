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

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	r := routes.Routes()

	// static file
	os.Mkdir("./uploads", os.ModePerm)
	r.Static("/uploads", "./uploads")

	// run routes
	if err := r.Run(":2406"); err != nil {
		log.Fatal(err)
	}

}
