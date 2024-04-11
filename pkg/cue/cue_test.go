package cue

import (
	"fmt"
	"os"
	"testing"

	"cuelang.org/go/cue"
	"github.com/AsaiYusuke/jsonpath"
	"github.com/kubevela/workflow/pkg/cue/model/sets"
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
func TestGetIndex(t *testing.T) {
	var result string
	err := Get(v, "certManager.rootCertificates[0].name", &result)
	require.NoError(t, err)
	require.Equal(t, "test1", result)
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

func TestReplaceValueInList(t *testing.T) {
	v2, err := Replace(v, "certManager.rootCertificates[0].name", "test2")
	require.NoError(t, err)

	var result string
	err = Get(v2, "certManager.rootCertificates[0].name", &result)
	require.NoError(t, err)
	require.Equal(t, "test2", result)
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

func TestDev(t *testing.T) {
	path := "certManager.rootCertificates[0].name"
	p := cue.ParsePath(path)

	v1, err := StringToCueValue(string(v))
	require.NoError(t, err)
	jsV, err := v1.MarshalJSON()
	require.NoError(t, err)

	var value = "marcel"

	empytValue := `string`

	v1.LookupPath(p)

	// s := fmt.Sprintf(`{ %s: %s }`, strings.ReplaceAll(path, ".", ":"), empytValue)
	s := fmt.Sprintf(`{ certManager:rootCertificates:[{name: %s}] }`, empytValue)
	fmt.Println(s)
	emptyBase := v1.Context().CompileString(s)
	n := emptyBase.FillPath(p, value)

	js, err := n.MarshalJSON()
	require.NoError(t, err)
	fmt.Println(string(js))

	ret, err := sets.StrategyUnify(v1, n, sets.UnifyByJSONMergePatch{})
	require.NoError(t, err)
	jsRet, err := ret.MarshalJSON()
	require.NoError(t, err)
	fmt.Println(string(jsV))
	fmt.Println(string(jsRet))
}

func TestJSONPath(t *testing.T) {
	f, err := jsonpath.Parse("certManager.rootCertificates[0].name")
	require.NoError(t, err)
	require.NotNil(t, f)
}
