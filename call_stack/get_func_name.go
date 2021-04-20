package main

import (
	"fmt"
	"reflect"
	"runtime"
	"runtime/debug"
	"strings"
)

func getFuncName(i interface{}, seps ...rune) string {
	fn := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()

	fields := strings.FieldsFunc(fn, func(sep rune) bool {
		for _, s := range seps {
			if sep == s {
				return true
			}
		}

		return false
	})

	fmt.Println(fields)
	if size := len(fields); size > 0 {
		return fields[size-1]
	}

	return ""
}

func foo() {}

func main() {
	fmt.Println("name:", getFuncName(foo))
	fmt.Println(getFuncName(debug.FreeOSMemory))
	fmt.Println(getFuncName(debug.FreeOSMemory, '.'))
	fmt.Println(getFuncName(debug.FreeOSMemory, '/', '.'))
}