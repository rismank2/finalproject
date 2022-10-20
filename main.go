package main

import (
	"finalproject/database"
	"finalproject/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":4828")
}
