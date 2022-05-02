package utils

import (
	"reflect"
	"testing"
)

func AsrtCallEq(t *testing.T, funcName string) func(expected any, actual any, args ...any) {
	return func(expected any, actual any, args ...any) {
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("expected %s to be %#v but found %#v", CallStr(funcName, args), expected, actual)
		}
	}
}
