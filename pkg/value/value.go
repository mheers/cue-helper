package value

import (
	"fmt"
	"strings"

	"cuelang.org/go/cue"
	"github.com/kubevela/workflow/pkg/cue/model/sets"
)

// Exists returns true if the value at the given path exists.
func Exists(v cue.Value, path string) bool {
	p := cue.ParsePath(path)
	ex := v.LookupPath(p)
	return ex.Exists()
}

// Set sets the value at the given path to the given value.
func Set(v cue.Value, path string, value interface{}) (cue.Value, error) {
	p := cue.ParsePath(path)
	ex := v.LookupPath(p)
	if ex.Exists() {
		return cue.Value{}, fmt.Errorf("value for %s already exists", path)
	}

	v = v.FillPath(p, value)
	err := v.Err()
	if err != nil {
		return cue.Value{}, err
	}

	return v, nil
}

// Get gets the value at the given path and decodes it into the given result.
func Get(v cue.Value, path string, result interface{}) error {
	p := cue.ParsePath(path)
	ex := v.LookupPath(p)
	if !ex.Exists() {
		return fmt.Errorf("path %s not found", path)
	}

	err := ex.Decode(result)
	if err != nil {
		return err
	}

	return nil
}

// Replace replaces the value at the given path with the given value.
func Replace(v cue.Value, path string, value interface{}) (cue.Value, error) {
	p := cue.ParsePath(path)

	ex := v.LookupPath(p)
	if !ex.Exists() {
		return Set(v, path, value)
	}

	emptyBase := v.Context().CompileString(fmt.Sprintf(`{ %s: [...] }`, strings.ReplaceAll(path, ".", ":")))
	n := emptyBase.FillPath(p, value)

	ret, err := sets.StrategyUnify(v, n, sets.UnifyByJSONMergePatch{})
	if err != nil {
		return cue.Value{}, err
	}

	return ret, nil
}
