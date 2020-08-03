package handlers

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

// APILogin logs in a user to the admin panel
func (hs *Handlers) APILogin(c echo.Context) error {
	type Request struct {
		User     string `json:"user"`
		Password string `json:"password"`
	}
	req := new(Request)
	err := c.Bind(req)
	if err != nil {
		return nil
	}
	if req.User != hs.Config.AppUserName {
		return c.String(http.StatusUnauthorized, "invalid credentials")
	}
	err = bcrypt.CompareHashAndPassword(
		hs.Config.AppUserPassword,
		[]byte(req.Password),
	)
	if err != nil {
		return c.String(http.StatusUnauthorized, "invalid credentials")
	}
	token := jwt.New(jwt.SigningMethodHS256)
	t, err := token.SignedString(hs.Config.AppSecret)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{"token": t})
}
