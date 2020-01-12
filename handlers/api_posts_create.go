package handlers

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/nickelghost/cms/db"
)

// APIPostsCreateRequest represents a request for creating a post
type APIPostsCreateRequest struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"min=10"`
	IsDraft bool   `json:"is_draft"`
}

// APIPostsCreateResponse represents a response of a created post
type APIPostsCreateResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	IsDraft   bool      `json:"is_draft"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// APIPostsCreate adds a new post for the website
func APIPostsCreate(c echo.Context) error {
	req := new(APIPostsCreateRequest)
	err := c.Bind(req)
	if err != nil {
		return err
	}
	err = validator.New().Struct(req)
	if err != nil {
		return err
	}
	sql, err := ioutil.ReadFile("queries/api_posts_create.sql")
	if err != nil {
		return err
	}
	res := new(APIPostsCreateResponse)
	err = db.Conn.QueryRow(
		string(sql),
		req.Title,
		req.Content,
		req.IsDraft,
	).Scan(
		&res.ID,
		&res.Title,
		&res.Content,
		&res.IsDraft,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}
