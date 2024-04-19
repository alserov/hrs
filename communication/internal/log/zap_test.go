package log

import (
	"context"
	"github.com/alserov/hrs/communication/internal/config"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestContextLogger(t *testing.T) {
	ctx := context.Background()
	l := NewLogger(config.Local)

	ctx = WithLogger(ctx, l)

	require.Equal(t, l, GetLogger(ctx))
}
