package main

import (
	"fmt"
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// func init() {
// 	// initialize database connection
// 	config.ConnDB()

// 	if err := godotenv.Load(); err != nil {
// 		log.Fatal(err)
// 	}
// }

func main() {
	// initialize database connection
	config.ConnDB()
	defer config.ConnDB().Close()

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	r := routes.Routes()

	// static file
	os.Mkdir("./uploads", os.ModePerm)
	r.Static("/uploads", "./uploads")

	fmt.Println(config.ConnDB().Stats())

	// run routes
	if err := r.Run(":2406"); err != nil {
		log.Fatal(err)
	}

}
