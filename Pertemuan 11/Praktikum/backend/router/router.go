package router

import (
	"inibackend/config/middleware"
	"inibackend/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", middleware.Middlewares("admin"))
	api.Get("/", handler.Homepage)
	api.Get("/mahasiswa", handler.GetAllMahasiswa)
	api.Get("/mahasiswa/:npm", handler.GetMahasiswaByNPM)
	api.Post("/mahasiswa", handler.CreateMahasiswa)
	api.Put("/mahasiswa/:npm", handler.UpdateMahasiswa)
	api.Delete("/mahasiswa/:npm", handler.DeleteMahasiswa)
	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)
}
