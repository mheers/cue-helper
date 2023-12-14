package cue

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const v = `
{
	certManager: {
		name: "certManager"
		rootCertificates: [
			{
				name: "test1"
				cert: "test1"
			}
		]
	}
	user: "demo"
	age: 35
	temperature: 25.7
}
`

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

func TestExists(t *testing.T) {
	require.True(t, Exists(v, "certManager.rootCertificates"))
	require.False(t, Exists(v, "test.rootCertificates"))
}

func TestGetString(t *testing.T) {
	var result string
	err := Get(v, "user", &result)
	require.NoError(t, err)
	require.Equal(t, "demo", result)
}
func TestGetInt(t *testing.T) {
	var result int
	err := Get(v, "age", &result)
	require.NoError(t, err)
	require.Equal(t, 35, result)
}

func TestGetFloat(t *testing.T) {
	var result float64
	err := Get(v, "temperature", &result)
	require.NoError(t, err)
	require.Equal(t, 25.7, result)
}

func TestReplaceString(t *testing.T) {
	v2, err := Replace(v, "user", "new")
	require.NoError(t, err)

	var result string
	err = Get(v2, "user", &result)
	require.NoError(t, err)
	require.Equal(t, "new", result)
}

func TestReplaceList(t *testing.T) {
	v2, err := Replace(v, "certManager.rootCertificates", []string{"123", "456"})
	require.NoError(t, err)

	var result []string
	err = Get(v2, "certManager.rootCertificates", &result)
	require.NoError(t, err)
	require.Equal(t, "123", result[0])
	require.Equal(t, "456", result[1])
}

func TestReplaceInt(t *testing.T) {
	v2, err := Replace(v, "age", 36)
	require.NoError(t, err)

	var result int
	err = Get(v2, "age", &result)
	require.NoError(t, err)
	require.Equal(t, 36, result)
}

func TestReplaceFloat(t *testing.T) {
	v2, err := Replace(v, "temperature", 12.4)
	require.NoError(t, err)

	var result float64
	err = Get(v2, "temperature", &result)
	require.NoError(t, err)
	require.Equal(t, 12.4, result)
}
