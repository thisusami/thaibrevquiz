package handlers

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/joho/godotenv"
	"github.com/thisusami/thaibrevquiz/controller"
	"github.com/thisusami/thaibrevquiz/db"
	"github.com/thisusami/thaibrevquiz/repositories"
	"github.com/thisusami/thaibrevquiz/services"
	"github.com/thisusami/thaibrevquiz/utils"
)

var JWTSecret []byte

func init() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using fallback values")
	}

	// Set JWTSecret from environment variable
	JWTSecret = []byte(getEnv("JWT_SECRET"))
}
func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return ""
}
func NewHandler() {
	mongo, err := db.NewMongoDbProperty(getEnv("CONNECTION_STRING"), getEnv("DB_NAME"))
	if err != nil {
		panic(err)
	}
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))
	repo := repositories.NewRepository(mongo)
	services := services.NewService(repo)
	controller := controller.NewController(app, services)
	app.Post("/register", controller.RegisterRoutes())
	app.Post("/login", controller.LoginRoutes(JWTSecret))
	app.Use("/api", jwtware.New(jwtware.Config{
		SigningKey: JWTSecret,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"code":    fiber.StatusUnauthorized,
				"error":   "Unauthorized",
				"message": "Invalid or expired JWT",
			})
		},
	}))
	app.Get("/api/get", utils.IsAuthenticated(), controller.GetRoutes())
	log.Printf("app listen to port%v", getEnv("PORT"))
	app.Listen(getEnv("PORT"))
}
