package route

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-gorm/config"
	"go-fiber-gorm/handler"
	"go-fiber-gorm/middleware"
)

func RouteInit(r *fiber.App) {
	r.Static("/public", config.ProjectRootPath+"/public/asset")

	r.Post("/login", handler.LoginHandler)

	r.Get("/user", middleware.AuthMiddleware, handler.UserHandlerGetAll)
	r.Get("/user/:id", handler.UserHandlerGetById)
	r.Post("/user", handler.UserHandlerCreate)
	r.Put("/user/:id", handler.UserhandlerUpdate)
	r.Put("/user/:id/update-email", handler.UserhandlerUpdateEmail)
	r.Delete("/user/:id", handler.UserHandlerDelete)

	r.Get("/product", handler.ProductHandlerGetAll)
	r.Post("/product", handler.ProductHandlerCreate)
	r.Get("/product/:id", handler.ProductHandlerGetById)
	r.Put("/product/:id", handler.ProducthandlerUpdate)
	r.Delete("/product/:id", handler.ProductHandlerDelete)

}
