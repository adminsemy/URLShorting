package model

import (
	"errors"
	"time"

	"github.com/samber/mo"
)

var (
	ErrNotFound         = errors.New("not found")
	ErrorIdenfierExists = errors.New("identidier already exists")
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
