package handlers

import (
	"elastic_search_customers/repository"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetCustomers(c echo.Context) error {
	clientId := c.Param("clientId")
	startDate := c.QueryParam("start")
	endDate := c.QueryParam("end")
	t, err := repository.GetCustomerBetweenDates(clientId, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, t)
}
