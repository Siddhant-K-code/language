package errors

import (
	"errors"
	"fmt"
)

var ErrUnsupportedJSON = errors.New("json not currently supported by dsl")

func UnsupportedDSLNestingError(typeName string, relationName string) error {
	return fmt.Errorf("the '%s' relation under the '%s' type has mixed operators which are not supported by the OpenFGA DSL syntax yet", relationName, typeName)
}
