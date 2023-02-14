package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"golang.org/x/mod/sumdb/storage"
	"gorm.io/gorm"
)

type Timespan struct {
	Start			int64		`json:"startTime"`
	End				int64		`json:"endTime"`
	Duration	int64		`json:"duration"`
}

type Code struct {
	Time	[]Timespan
}

type Workout struct {
	Time	[]Timespan
}

type Rest struct {
	Time	[]Timespan
}

type Session struct {
	ID				int				`json:"id"`
	Code			[]Code		`json:"code"`
	Workout		[]Workout	`json:"workout"`
	Rest			[]Rest		`json:"rest"`
}

type Repository struct {
	DB *gorm.DB
}

func(r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_session", r.CreateSession)
	api.Delete("/delete_session/:id", r.DeleteSession)
	api.Get("/get_session/:id", r.GetSessionByID)
	api.Get("/sessions", r.GetAllSessions)
}

func main(){
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	db, err := storage.NewConnection(config.Storage)

	if err != nil {
		log.Fatal("Could not load database")
	}

	r := Repository {
		DB: db, 
	}

	app := fiber.New()
	r.SetupRoutes(app)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	log.Fatal(app.Listen(":4000"))

	
}