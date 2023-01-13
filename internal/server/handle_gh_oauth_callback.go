package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/adminsemy/URLShorting/internal/model"
	"github.com/labstack/echo/v4"
)

type callBackProvider interface {
	GitHubAuthCallback(context.Context, string) (*model.User, string, error)
}

func HandleGitHubAuthCallback(callBackProvider callBackProvider) echo.HandlerFunc {
	return func(c echo.Context) error {
		sessionCode := c.QueryParam("code")
		if sessionCode == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "missing code")
		}

		_, jwt, err := callBackProvider.GitHubAuthCallback(c.Request().Context(), sessionCode)
		if err != nil {
			log.Printf("error handling github auth callback: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		redirectURL := fmt.Sprintf("http://localhost:8080/auth/token.html?token=%s", jwt)

		return c.Redirect(http.StatusMovedPermanently, redirectURL)
	}
}
