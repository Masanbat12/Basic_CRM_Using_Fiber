package main

import (
	"fmt"

	"github.com/akhil/CRM-Fiber/database"
	"github.com/akhil/CRM-Fiber/lead"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setupRoutes(app *fiber.App){
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatbase(){
	var err error
	database.DBCon, err = gorm.Open("sqlite3", "leads.db")
	//golang will handle on every single operation that taken place
	if err != nil{
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}
	fmt.Println("Connection opened to database")
	database.DBCon.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatbase()
	setupRoutes(app)
	app.Listen(3000)
	defer database.DBCon.Close()
}
