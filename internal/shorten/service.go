package shorten

import (
	"context"

	"github.com/adminsemy/URLShorting/internal/model"
	"github.com/google/uuid"
)

type Storage interface {
	Put(ctx context.Context, shortening model.Shortening) (*model.Shortening, error)
	Get(ctx context.Context, input string) (*model.Shortening, error)
	IncrementVisits(ctx context.Context, identifier string) error
}
type Service struct {
	storage Storage
}

func NewService(storage Storage) *Service {
	return &Service{storage: storage}
}

func (s *Service) Shorten(ctx context.Context, input model.ShortenInput) (*model.Shortening, error) {
	var (
		id         = uuid.New().ID()
		identifier = input.Identifier.OrElse(Shorten(id))
	)

	inputShortening := model.Shortening{
		Identidier:  identifier,
		OriginalURL: input.RawURL,
	}

	shortening, err := s.storage.Put(ctx, inputShortening)
	if err != nil {
		return nil, err
	}
	return shortening, nil

}

func (s *Service) Redirect(ctx context.Context, identifier string) (string, error) {
	shortening, err := s.storage.Get(ctx, identifier)
	if err != nil {
		return "", nil
	}

	if err := s.storage.IncrementVisits(ctx, identifier); err != nil {
		return "", nil
	}

	return shortening.OriginalURL, nil
}
