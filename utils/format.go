package utils

import (
	"fmt"
	"strings"
)

//"Expected Maybe value %v to be of type Some or None, but its type is %T", m, v

func BadSumTypeConstructor(val any, typeName string, typeConstructors ...string) string {
	return fmt.Sprintf("expected %v to be of type %s :: %s, but its type is %T", val, typeName, strings.Join(typeConstructors, " | "), val)
}

func CallStr(funcName string, args []any) string {
	argStrs := make([]string, 0, len(args))

	for _, arg := range args {
		argStrs = append(argStrs, fmt.Sprintf("%v", arg))
	}

	return fmt.Sprintf("%s(%s)", funcName, strings.Join(argStrs, ", "))
}
