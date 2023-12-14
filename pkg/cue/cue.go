package cue

import (
	"bytes"
	"context"
	"fmt"

	"cuelang.org/go/cmd/cue/cmd"
	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"github.com/mheers/cue-helper/helpers"
	"github.com/mheers/cue-helper/pkg/value"
)

func Exists(data, path string) bool {
	v, err := StringToCueValue(data)
	if err != nil {
		return false
	}
	return value.Exists(v, path)
}

func Set(data, path string, x interface{}) (string, error) {
	v, err := StringToCueValue(data)
	if err != nil {
		return "", err
	}
	result, err := value.Set(v, path, x)
	if err != nil {
		return "", err
	}

	return CueValueToString(result), nil
}

func Get(data, path string, result interface{}) error {
	v, err := StringToCueValue(data)
	if err != nil {
		return err
	}
	return value.Get(v, path, result)
}

func Replace(data, path string, x interface{}) (string, error) {
	v, err := StringToCueValue(data)
	if err != nil {
		return "", err
	}
	result, err := value.Replace(v, path, x)
	if err != nil {
		return "", err
	}

	return CueValueToString(result), nil
}

func Format(data string) (string, error) {
	v, err := StringToCueValue(data)
	if err != nil {
		return "", err
	}
	return CueValueToString(v), nil
}

func CueValueToString(v cue.Value) string {
	return fmt.Sprintf("%.2v", v)
}

func StringToCueValue(data string) (cue.Value, error) {
	ctx := cuecontext.New()
	v := ctx.CompileString(data)
	if err := v.Err(); err != nil {
		return cue.Value{}, err
	}

	return v, nil
}

func Render(path string) (string, error) {
	cueCmd, err := cmd.New([]string{"export", "--out", "cue"})
	if err != nil {
		return "", err
	}

	var stdOut bytes.Buffer

	cueCmd.SetOutput(&stdOut)

	err = cueCmd.Run(context.Background())
	if err != nil {
		return "", err
	}

	return stdOut.String(), nil
}

func RenderBin(path string) (string, error) {
	stdout, stderr, err := helpers.RunResultDir("../../example/", "cue", "export", "--out", "cue")
	if err != nil {
		return "", fmt.Errorf("failed to run cue: %w", err)
	}
	if stderr != "" {
		return "", fmt.Errorf("cue failed: %s", stderr)
	}
	return stdout, nil
}
