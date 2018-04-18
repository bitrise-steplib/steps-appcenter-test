package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInstalledInPath(t *testing.T) {
	require.True(t, installedInPath("env"))
}
