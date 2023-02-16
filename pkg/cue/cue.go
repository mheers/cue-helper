package cue

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"github.com/mheers/cue-helper/pkg/value"
)

func Exists(data, path string) bool {
	v := cueValue(data)
	return value.Exists(v, path)
}

func Set(data, path string, x interface{}) (string, error) {
	v := cueValue(data)
	result, err := value.Set(v, path, x)
	if err != nil {
		return "", err
	}

	return cueString(result), nil
}

func Get(data, path string, result interface{}) error {
	v := cueValue(data)
	return value.Get(v, path, result)
}

func Replace(data, path string, x interface{}) (string, error) {
	v := cueValue(data)
	result, err := value.Replace(v, path, x)
	if err != nil {
		return "", err
	}

	return cueString(result), nil
}

func cueString(v cue.Value) string {
	return fmt.Sprintf("%.2v", v)
}

func cueValue(data string) cue.Value {
	ctx := cuecontext.New()
	return ctx.CompileString(data)
}
