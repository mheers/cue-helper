package value

import (
	"fmt"
	"testing"

	"cuelang.org/go/cue/cuecontext"
	"github.com/stretchr/testify/require"
)

type RootCertificate struct {
	Name    string `json:"name"`
	CertB64 string `json:"cert"`
}

func TestAppend(t *testing.T) {
	require := require.New(t)

	ctx := cuecontext.New()
	v := ctx.CompileString(`
		{
			rootCertificates: [
				{
					name: "test1"
					cert: "test1"
				},
			]
		}
`)

	path := "rootCertificates"

	nRootCerts := []RootCertificate{
		{
			Name:    "test1",
			CertB64: "test1",
		},
		{
			Name:    "test2",
			CertB64: "test2",
		},
	}

	ret, err := Replace(v, path, nRootCerts)
	require.NoError(err)
	require.NotNil(ret)

	result := fmt.Sprintf("%v", ret)
	require.Equal(`{
	rootCertificates: [{
		cert: "test1"
		name: "test1"
	}, {
		cert: "test2"
		name: "test2"
	}, ...]
}`, result)
}

func TestAppendComplexPath(t *testing.T) {
	require := require.New(t)

	ctx := cuecontext.New()
	v := ctx.CompileString(`
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
		}
`)

	path := "certManager.rootCertificates"

	nRootCerts := []RootCertificate{
		{
			Name:    "test1",
			CertB64: "test1",
		},
		{
			Name:    "test2",
			CertB64: "test2",
		},
	}

	ret, err := Replace(v, path, nRootCerts)
	require.NoError(err)
	require.NotNil(ret)

	result := fmt.Sprintf("%v", ret)
	require.Equal(`{
	certManager: {
		name: "certManager"
		rootCertificates: [{
			cert: "test1"
			name: "test1"
		}, {
			cert: "test2"
			name: "test2"
		}, ...]
	}
	user: "demo"
}`, result)
}

func TestClear(t *testing.T) {
	require := require.New(t)

	ctx := cuecontext.New()
	v := ctx.CompileString(`
		{
			rootCertificates: [
				{
					name: "test1"
					cert: "test1"
				},
			]
		}
`)

	path := "rootCertificates"

	nRootCerts := []RootCertificate{}

	ret, err := Replace(v, path, nRootCerts)
	require.NoError(err)
	require.NotNil(ret)

	result := fmt.Sprintf("%v", ret)
	require.Equal(`{
	rootCertificates: [...]
}`, result)
}
