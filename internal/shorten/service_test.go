package shorten

import (
	"context"
	"testing"

	"github.com/adminsemy/URLShorting/internal/model"
	"github.com/adminsemy/URLShorting/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestService_Shorten(t *testing.T) {
	t.Run("generates shorting for a given URL ", func(t *testing.T) {
		svc := NewService(storage.NewInMemory())
		input := model.ShortenInput{RawURL: "https://google.com"}
		shortening, err := svc.Shorten(context.Background(), input)
		require.NoError(t, err)
		require.NotEmpty(t, shortening.Identidier)
		assert.Equal(t, input.RawURL, shortening.OriginalURL)
		assert.NotZero(t, shortening.CreatedAt)
	})
}
