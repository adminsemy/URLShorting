package user

import (
	"context"
	"sync"

	"github.com/adminsemy/URLShorting/internal/model"
)

type inMemory struct {
	m sync.Map
}

func NewUserInMemory() *inMemory {
	return &inMemory{}
}

func (i *inMemory) CreateOrUpdate(_ context.Context, user model.User) (*model.User, error) {
	i.m.Store(user.GithubLogin, user)
	return &user, nil
}

func (i *inMemory) Update(_ context.Context, user model.User) error {
	i.m.Store(user.GithubLogin, user)
	return nil
}

func (i *inMemory) GetByGithubLogin(_ context.Context, login string) (*model.User, error) {
	if user, ok := i.m.Load(login); ok {
		return user.(*model.User), nil
	}
	return nil, model.ErrNotFound
}

func (i *inMemory) Deactivate(_ context.Context, login string) error {
	if user, ok := i.m.Load(login); ok {
		user.(*model.User).IsActivity = false
		i.m.Store(login, user)
	}
	return nil
}
