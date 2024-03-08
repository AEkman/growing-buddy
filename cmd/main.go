package main

import (
	"com.andreasekman/AEkman/growing-buddy/internal/infrastructure"
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"

	_ "github.com/lib/pq"
)

var version = "dev"

func main() {
	fmt.Println("Growing Buddy application started!")
	fmt.Printf("Version: %s\n", version)

	config := infrastructure.NewConfig()

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.PostgresHost(), config.PostgresPort(), config.PostgresUser(), config.PostgresPassword(), config.PostgresDbname())

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Unable to ping database: %v\n", err)
	}

	err = createChildTable(db)
	if err != nil {
		log.Fatalf("Unable to create table: %v\n", err)
	}

	fmt.Println("Child table created successfully.")

	app := fiber.New()

	SetupRoutes(app)

	app.Listen(`:` + config.GrowingBuddyPort())
}

func createChildTable(db *sql.DB) error {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS children (
            id SERIAL PRIMARY KEY,
            first_name TEXT,
            last_name TEXT,
            gender TEXT,
            birth_date DATE
        )
    `)
	return err
}

func SetupRoutes(app *fiber.App) fiber.Router {
	return app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Growing Buddy says hello!")
	})
}
