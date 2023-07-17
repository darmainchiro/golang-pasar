package routes

import (
	"pasar/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) *echo.Echo {
	e.GET("/pasar", controllers.GetPasar)
	e.GET("/pasar/:id", controllers.DetailPasar)
	e.POST("/pasar", controllers.PostPasar)
	e.PUT("/pasar/:id", controllers.UpdatePasar)
	e.DELETE("/pasar/:id", controllers.DeletePasar)

	return e
}
