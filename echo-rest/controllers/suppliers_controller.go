package controllers

import (
	"net/http"
	"pnp/echo-rest/models"

	"github.com/labstack/echo"
)

//FetchAllSuppliers ...
func FetchAllSuppliers(c echo.Context) (err error) {

	result, err := models.FetchSuppliers()

	return c.JSON(http.StatusOK, result)
}

//AddSupplier ...
func AddSupplier(c echo.Context) (err error) {

	result, err := models.AddSupplier(c)

	return c.JSON(http.StatusOK, result)
}

//UpdateSuppliers ...
func UpdateSuppliers(c echo.Context) (err error) {

	result, err := models.UpdateSupplier(c)

	return c.JSON(http.StatusOK, result)
}

//DeleteSuppliers ...
func DeleteSuppliers(c echo.Context) (err error) {

	result, err := models.DeleteSupplier(c)

	return c.JSON(http.StatusOK, result)
}
