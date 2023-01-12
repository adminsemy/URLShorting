package model

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/samber/mo"
)

var (
	ErrNotFound         = errors.New("not found")
	ErrorIdenfierExists = errors.New("identidier already exists")
	ErrUserIsNotMember  = errors.New("user is not member of organization")
)

type Shortening struct {
	Identidier  string    `json:"identidier,omitempty"`
	OriginalURL string    `json:"original_url,omitempty"`
	Visits      int64     `json:"visits,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type ShortenInput struct {
	RawURL     string
	Identifier mo.Option[string]
	CreatedBy  string
}

type User struct {
	IsActivity      bool      `json:"is_activity,omitempty"`
	GithubLogin     string    `json:"github_login,omitempty"`
	GithubAccessKey string    `json:"github_access_key,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
}

type UserClaims struct {
	jwt.RegisteredClaims
	User `json:"user_data"`
}
