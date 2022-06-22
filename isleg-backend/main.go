package main

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/routes"
	"log"
	"os"
)

func init() {
	config.ConnDB()
}

func main() {

	r := routes.Routes()

	os.Mkdir("./uploads", os.ModePerm)
	r.Static("/image", "./uploads")

	if err := r.Run(":2406"); err != nil {
		log.Fatal(err)
	}

}
