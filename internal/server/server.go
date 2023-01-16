package server

import (
	"context"
	"net/http"

	"github.com/adminsemy/URLShorting/internal/auth"
	"github.com/adminsemy/URLShorting/internal/model"
	"github.com/adminsemy/URLShorting/internal/shorten"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CloseFunc func(context.Context) error

type Server struct {
	e         *echo.Echo
	shortener *shorten.Service
	closers   []CloseFunc
	auth      *auth.Service
}

func NewServer(shortener *shorten.Service, auth *auth.Service) *Server {
	s := &Server{
		shortener: shortener,
		auth:      auth,
	}
	s.setupRouter()

	return s
}

func (s *Server) AddCloser(closer CloseFunc) {
	s.closers = append(s.closers, closer)

}

func (s *Server) setupRouter() {
	s.e = echo.New()
	s.e.HideBanner = true
	s.e.Validator = NewValidator()

	s.e.Pre(middleware.RemoveTrailingSlash())
	s.e.Use(middleware.RequestID())

	s.e.GET("/auth/oauth/github/link", HandleGetGithubAuthLinkProvider(s.auth))
	s.e.GET("/auth/oauth/github/callback", HandleGitHubAuthCallback(s.auth))
	s.e.GET("/auth/token.html", HandleTokenPage())
	s.e.GET("/static", HandleStatic())

	restricted := s.e.Group("/api")
	{
		//restricted.Use(middleware.JWTWithConfig(makeJWTConfig()))
		restricted.POST("/shorten", HandleShorten(s.shortener))
	}
	s.e.GET("/:identifier", HandleRedirect(s.shortener))

	s.AddCloser(s.e.Shutdown)

}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.e.ServeHTTP(w, r)
}

func (s *Server) Shutdown(ctx context.Context) error {
	for _, fn := range s.closers {
		if err := fn(ctx); err != nil {
			return err
		}
	}
	return nil
}

func makeJWTConfig() middleware.JWTConfig {
	return middleware.JWTConfig{
		SigningKey: []byte("JWT"),
		Claims:     &model.UserClaims{},
		ErrorHandler: func(err error) error {
			return echo.NewHTTPError(http.StatusUnauthorized)
		},
	}
}
