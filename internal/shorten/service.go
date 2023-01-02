package shorten

import (
	"context"

	"github.com/adminsemy/URLShorting/internal/model"
	"github.com/google/uuid"
)

type Service struct {
}

func (s *Service) Shorten(ctx context.Context, input model.ShortenInput) (*model.Shortening, error) {
	var (
		id         = uuid.New().ID()
		identifier = input.Identifier.OrElse(Shorten(id))
	)

	return &model.Shortening{
		Identidier:  identifier,
		OriginalURL: input.RawURL,
	}, nil

}
