package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"todo-service/internal/db"
//	"todo-service/internal/auth"
	"todo-service/internal/repo"
	"todo-service/internal/models"
)


var tdb = db.InitDB()
var trs = repo.NewTRS(tdb)

func AddUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest, "Invalid fields")
	}

	var err error
/*	if user.PdHash, err = auth.HashPd(user.PdHash); err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"Hash password error\n" + err.Error())
	}
*/
	if err = trs.AddUser(user); err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"Error adding a user to the db\n" + err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Successfully added user",
		"token": "JWT token",
	})
}
