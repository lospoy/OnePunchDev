package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/lospoy/onepunchdev/models"
	"github.com/lospoy/onepunchdev/storage"
	"gorm.io/gorm"
)

type Session struct {
	Code			string		`json:"code"`
	Workout		string		`json:"workout"`
	Rest			string		`json:"rest"`
}

// talks to the database
type Repository struct {
	DB *gorm.DB
}

// CREATE SESSION
func (r *Repository) CreateSession(context *fiber.Ctx) error {
	session := Session{}

	err := context.BodyParser(&session)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message":"create session request failed",
		})
			return err
	}

	err = r.DB.Create(&session).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message":"could not create session",
		})
			return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message":"session has been created",
	})
	return nil
}

// GET ALL SESSIONS
func (r *Repository) GetAllSessions(context *fiber.Ctx) error {
	sessionModels := &[]models.Sessions{}

	err := r.DB.Find(sessionModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message":"could not find all sessions",
		})
			return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message":"all sessions fetched successfully",
		"data": sessionModels,
	})
	return nil
}

// DELETE SESSION
func (r *Repository) DeleteSession(context *fiber.Ctx) error {
	sessionModel := models.Sessions{}
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message":"id cannot be empty",
		})
		return nil
	}

	err := r.DB.Delete(sessionModel, id).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message":"could not delete session",
		})
			return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message":"session deleted successfully",
	})
	return nil
}

// GET SESSION BY ID
func (r *Repository) GetSessionByID(context *fiber.Ctx) error {
	id := context.Params("id")
	sessionModel := &models.Sessions{}

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message":"id cannot be empty",
		})
		return nil
	}

	fmt.Println("the ID is ", id)

	err := r.DB.Where("id = ?", id).First(sessionModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message":"could not get the session",
		})
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message":"session fetched successfully",
		"data": sessionModel,
	})
	return nil
}

// ROUTES
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

	config := &storage.Config{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		User: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName: os.Getenv("DB_DBNAME"),
		SSLMode: os.Getenv("DB_SSLMODE"),
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("Could not load database")
	}

	err = models.MigrateSessions(db)
	if err != nil {
		log.Fatal("Could not migrate db")
	}

	// initialized repository
	r := Repository {
		DB: db, 
	}

	app := fiber.New()
	r.SetupRoutes(app)
	log.Fatal(app.Listen(":4000"))

	
}