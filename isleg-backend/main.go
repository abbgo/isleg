package main

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/routes"
	"log"
	"os"
)

func init() {
	// initialize database connection
	config.ConnDB()
}

func main() {

	r := routes.Routes()

	// static file
	os.Mkdir("./uploads", os.ModePerm)
	r.Static("/image", "./uploads")

	// run routes
	if err := r.Run(":2406"); err != nil {
		log.Fatal(err)
	}

}
