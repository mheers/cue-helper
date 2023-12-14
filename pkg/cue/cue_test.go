package cue

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStringToCueValue(t *testing.T) {
	data, err := os.ReadFile("../../example/config.example.cue")
	require.NoError(t, err)
	require.NotEmpty(t, data)

	v, err := StringToCueValue(string(data))
	require.NoError(t, err)
	require.NotNil(t, v)
}

func TestRenderBin(t *testing.T) {
	dir := "../../example/"
	data, err := RenderBin(dir)
	require.NoError(t, err)
	require.NotEmpty(t, data)
}
