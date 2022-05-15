package main

import (
	"final-project-2/database"
	"final-project-2/router"
)

func main(){
	database.MulaiDB()
	r := router.MulaiApp()
	r.Run(":8080")
}