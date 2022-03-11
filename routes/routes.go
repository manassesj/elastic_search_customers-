package routes

import (
	"elastic_search_customers/handlers"
	"github.com/labstack/echo/v4"
)

func Routes() *echo.Echo {
	e := echo.New()
	e.GET("/customers/:clientId/products", handlers.GetCustomers)
	return e
}
