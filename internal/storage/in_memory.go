package storage

import (
	"context"
	"sync"
	"time"

	"github.com/adminsemy/URLShorting/internal/model"
)

type inMemory struct {
	m sync.Map
}

func NewInMemory() *inMemory {
	return &inMemory{}
}

func (s *inMemory) Put(_ context.Context, shortening model.Shortening) (*model.Shortening, error) {
	if _, exists := s.m.Load(shortening.Identidier); exists {
		return nil, model.ErrorIdenfierExists
	}
	shortening.CreatedAt = time.Now().UTC()

	s.m.Store(shortening.Identidier, shortening)

	return &shortening, nil
}

func (s *inMemory) Get(_ context.Context, identifier string) (*model.Shortening, error) {
	v, ok := s.m.Load(identifier)
	if !ok {
		return nil, model.ErrNotFound
	}

	shortering := v.(model.Shortening)

	return &shortering, nil
}

func (s *inMemory) IncrementVisitsd(_ context.Context, identifier string) error {
	v, ok := s.m.Load(identifier)
	if !ok {
		return model.ErrNotFound
	}

	shortering := v.(model.Shortening)
	shortering.Visits++

	s.m.Store(identifier, shortering)

	return nil

}
