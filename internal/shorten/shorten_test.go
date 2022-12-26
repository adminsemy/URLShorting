package shorten_test

import (
	"testing"

	"github.com/adminsemy/URLShorting/internal/shorten"
	"github.com/stretchr/testify/assert"
)

func TestShorten(t *testing.T) {
	t.Run("Возвращаем идентификатор", func(t *testing.T) {
		type testCase struct {
			id       uint32
			expected string
		}
		testCases := []testCase{
			{
				id:       1024,
				expected: "Mv",
			},
			{
				id:       0,
				expected: "",
			},
		}
		for _, tc := range testCases {
			actual := shorten.Shorten(tc.id)
			assert.Equal(t, tc.expected, actual)
		}
	})
}
