package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func FuzzDefaultLimit(f *testing.F) {
	testcases := []int{0, 1, 2}
	for _, tc := range testcases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, orig int) {
		limit := defaultLimit(orig)
		if orig == 0 {
			assert.Equal(t, defaultLimitValue, limit)
		} else {
			assert.Equal(t, orig, limit)
		}
	})
}
