package main

import (
	"fmt"

	i "github.com/bdn/jeker/db"
	r "github.com/bdn/jeker/routes"
	"github.com/gofiber/fiber"
)

func main() {
	db, err := i.InitDB()
	sqlDb, err := db.DB()
	defer sqlDb.Close()
	if err != nil {
		fmt.Println(fmt.Sprintf("[ERROR] Failed To Connect To DB: %s", err.Error()))
		sqlDb.Close()
		panic("failed to connect database")
	}

	app := fiber.New()
	r.UserRoutes(app)
	r.ProfileRoutes(app)
	r.BookRoutes(app)
	app.Listen(":3000")
}
