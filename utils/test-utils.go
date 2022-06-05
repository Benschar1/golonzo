package utils

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func AsrtCallEq(t *testing.T, FuncName string) func(Want any, actual any, args ...any) {
	return func(Want any, actual any, args ...any) {
		if !reflect.DeepEqual(Want, actual) {
			t.Errorf("Want %s to be %#v but found %#v", CallStr(FuncName, args), Want, actual)
		}
	}
}

type FuncTc struct {
	Func   any
	Inputs []any
	Want   []any
}

func MakeFtc1(Func any, Want any, Inputs ...any) FuncTc {
	return FuncTc{Func, Inputs, []any{Want}}
}

type FuncUnitTests struct {
	Cases map[string]FuncTc
}

func (fut FuncUnitTests) Execute(t *testing.T) {
	for name, tcase := range fut.Cases {
		t.Run(name, func(t *testing.T) {

			curResult, restArgs, numIn := []any{tcase.Func}, tcase.Inputs, 0
			var theseArgs []any

			for len(restArgs) != 0 {
				numIn = reflect.TypeOf(curResult[0]).NumIn()

				theseArgs = restArgs[:numIn]
				restArgs = restArgs[numIn:]

				curResult = apply(curResult[0], theseArgs)
			}

			if !reflect.DeepEqual(curResult, tcase.Want) {
				t.Errorf("\nwant: %s\ngot : %s", funcOutputStr(tcase.Want), funcOutputStr(curResult))
			}

		})
	}
}

func funcOutputStr(arr []any) string {
	switch len(arr) {
	case 0:
		return "()"
	case 1:
		return fmt.Sprint(arr[0])
	default:
		strs := make([]string, len(arr))
		for i, el := range arr {
			strs[i] = fmt.Sprint(el)
		}
		return fmt.Sprintf("(%s)", strings.Join(strs, ", "))
	}
}

func anysToVals(anys []any) []reflect.Value {
	vals := make([]reflect.Value, len(anys))
	for i, el := range anys {
		vals[i] = reflect.ValueOf(el)
	}
	return vals
}

func valsToAnys(vals []reflect.Value) []any {
	anys := make([]any, len(vals))
	for i, val := range vals {
		anys[i] = val.Interface()
	}
	return anys
}

func apply(f any, ins []any) []any {
	fVal := reflect.ValueOf(f)
	inVals := anysToVals(ins)
	outVals := fVal.Call(inVals)
	return valsToAnys(outVals)
}
