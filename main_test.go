package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestName(t *testing.T) {
	require.NotPanics(t, func() {
		main()
	})
}
