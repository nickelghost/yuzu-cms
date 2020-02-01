package handlers

import (
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

type APILoginRequest struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

func (hs *Handlers) APILogin(c echo.Context) error {
	req := new(APILoginRequest)
	err := c.Bind(req)
	if err != nil {
		return nil
	}
	if req.User != os.Getenv("APP_USER_NAME") {
		return c.String(http.StatusUnauthorized, "invalid credentials")
	}
	log.Info(os.Getenv("APP_USER_PASS"))
	err = bcrypt.CompareHashAndPassword(
		[]byte(os.Getenv("APP_USER_PASS")),
		[]byte(req.Password),
	)
	if err != nil {
		return c.String(http.StatusUnauthorized, "invalid credentials")
	}
	token := jwt.New(jwt.SigningMethodHS256)
	t, err := token.SignedString([]byte(os.Getenv("APP_SECRET")))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{"token": t})
}
