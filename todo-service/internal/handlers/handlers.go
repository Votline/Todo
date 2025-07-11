package handlers

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"

	"todo-service/internal/repo"
	"todo-service/internal/models"
	
	auth "todo/auth-service/pb"
)

type Handler struct {
	AuthClient auth.AuthServiceClient
	Tdb *sqlx.DB
	Trs *repo.TodoRepoSql
}

func (h *Handler) AddUser(c echo.Context) error {
	ctx := c.Request().Context()
	var user models.User
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest, "Invalid fields")
	}

	hashRes, err := h.AuthClient.HashPd(ctx, &auth.HashReq{Password: user.PdHash})
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"Hash password error\n" + err.Error())
	}
	user.PdHash = hashRes.Hash

	if err = h.Trs.AddUser(user); err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"Error adding a user to the db\n" + err.Error())
	}

	token, err := h.AuthClient.GenJWT(
		ctx, &auth.JWTReq{UserID: user.Id})
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"Error generating the JWT token\n" + err.Error())
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Successfully added user",
		"token": token.Token,
	})
}
