package id

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenShortID(t *testing.T) {
	shortID := GenShortID()
	fmt.Println(shortID)
	assert.NotEqual(t, "", shortID)
	assert.Equal(t, 6, len(shortID))
}

func BenchmarkGenShortID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenShortID()
	}
}

func BenchmarkGenShortIDTimeConsuming(b *testing.B) {
	b.StopTimer()

	shortId := GenShortID()
	if shortId == "" {
		b.Error("Failed to generate short id")
	}

	b.StartTimer() //重新开始时间

	for i := 0; i < b.N; i++ {
		GenShortID()
	}
}
