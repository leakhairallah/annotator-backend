package annotation

import "github.com/labstack/echo/v4"

func MapAnnotationRoutes(commGroup *echo.Group, handlers Handlers) {
	commGroup.POST("", handlers.Create())
	commGroup.GET("", handlers.GetAll())
	commGroup.PUT("/:id", handlers.Update())
	commGroup.DELETE("/:id", handlers.Delete())
}
