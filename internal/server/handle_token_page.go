package server

import (
	"bytes"
	"embed"
	"html/template"
	"log"
	"net/http"

	"github.com/adminsemy/URLShorting/internal/model"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/go-github/v48/github"
	"github.com/labstack/echo/v4"
)

var static embed.FS

func HandleTokenPage() echo.HandlerFunc {
	tmpl, err := template.ParseFS(static, "../../public/static/token.html")
	if err != nil {
		log.Fatalf("error parsing token.html template file: %v", err)
	}

	type templateData struct {
		Token                   string
		TelegramContactUserName string
		GithubUserName          string
		GithubAvatarURL         string
	}

	type request struct {
		Token string `query:"token"`
	}

	return func(c echo.Context) error {
		var req request
		if err := c.Bind(&req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
		}

		tokenObj, err := jwt.ParseWithClaims(
			req.Token,
			&model.UserClaims{},
			func(t *jwt.Token) (interface{}, error) {
				return []byte("JWT"), nil
			},
		)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		if !tokenObj.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		claims, ok := tokenObj.Claims.(*model.UserClaims)

		if !ok {
			log.Printf("error asserting claims to *model.UserClaims")
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		ghClient := github.NewClient(nil)
		ghUser, _, err := ghClient.Users.Get(c.Request().Context(), claims.User.GithubLogin)
		if err != nil {
			log.Printf("error getting github user: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		var (
			buf  bytes.Buffer
			data = templateData{
				Token:           req.Token,
				GithubUserName:  claims.User.GithubLogin,
				GithubAvatarURL: ghUser.GetAvatarURL(),
			}
		)
		if err := tmpl.Execute(&buf, data); err != nil {
			log.Printf("error executing token.html template: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return c.HTML(http.StatusOK, buf.String())
	}
}

func HandleStatic() echo.HandlerFunc {
	return echo.WrapHandler(http.FileServer(http.FS(static)))
}
