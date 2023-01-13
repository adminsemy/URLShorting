package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type githubAuthLinkProvider interface {
	GithubAuthLinkProvider() string
}

func HandleGetGithubAuthLinkProvider(provider githubAuthLinkProvider) echo.HandlerFunc {
	return func(c echo.Context) error {
		link := provider.GithubAuthLinkProvider()
		return c.JSON(http.StatusOK, map[string]string{"link": link})
	}
}
