package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/adminsemy/URLShorting/internal/model"
	"github.com/adminsemy/URLShorting/internal/shorten"
	"github.com/labstack/echo/v4"
	"github.com/samber/mo"
)

type shortener interface {
	Shorten(context.Context, model.ShortenInput) (*model.Shortening, error)
}

type shortenRequest struct {
	URL        string `json:"url" validate:"required,url"`
	Identifier string `json:"identifier,omitempty"`
}

type shortenResponse struct {
	ShortURL string `json:"short_url,omitempty"`
	Message  string `json:"message,omitempty"`
}

func HandleShorten(shortener shortener) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req shortenRequest
		if err := c.Bind(&req); err != nil {
			return err
		}

		if err := c.Validate(req); err != nil {
			return err
		}

		identifier := mo.None[string]()

		if strings.TrimSpace(req.Identifier) != "" {
			identifier = mo.Some(req.Identifier)
		}

		input := model.ShortenInput{
			RawURL:     req.URL,
			Identifier: identifier,
		}

		shortening, err := shortener.Shorten(c.Request().Context(), input)

		if err != nil {
			if errors.Is(err, model.ErrorIdenfierExists) {
				return echo.NewHTTPError(http.StatusConflict, err.Error())
			}

			log.Printf("shortener.Shorten: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		shortURL, err := shorten.PrependBaseURL("http://localhost:8080", shortening.Identidier)

		if err != nil {
			log.Printf("shorten.PrependBaseURL: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, shortenResponse{ShortURL: shortURL})
	}
}
