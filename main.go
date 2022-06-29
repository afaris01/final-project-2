package main

import (
	"final-project-2/database"
	"final-project-2/router"
	"os"
)

func main() {
	database.MulaiDB()
	r := router.MulaiApp()
	port := os.Getenv("DB_PORT")
	r.Run(":" + port)
}
