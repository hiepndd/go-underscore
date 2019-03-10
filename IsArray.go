package underscore

import (
	"reflect"
)

// IsList is function which check whether argument passed is list(array or slices) or not
func IsList(list interface{}) bool {
	result := reflect.ValueOf(list)
	return result.Kind() == reflect.Array || result.Kind() == reflect.Slice
}
