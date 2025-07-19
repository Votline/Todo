package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	auth "todo/auth-service/pb"
	"todo-service/internal/models"
)

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
			http.StatusInternalServerError, err.Error())
	}
	user.PdHash = hashRes.Hash

	if err = h.Trs.AddUser(user); err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError, err.Error())
	}

	token, err := h.AuthClient.GenJWT(
		ctx, &auth.JWTReq{UserID: user.Id})
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": token.Token,
	})
}
