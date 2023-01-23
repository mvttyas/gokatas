// Myprint shows how come fmt.Print can print (any number of) arguments of any
// type. It uses reflection. See research.swtch.com/gotour, from go.dev/doc, for
// more.
//
// Level: intermediate
// Topics: reflect, strconv, switch, interfaces
package main

import (
	"os"
	"reflect"
	"strconv"
)

func main() {
	myPrint("hello", 42)
}

func myPrint(args ...any) {
	for _, arg := range args {
		switch v := reflect.ValueOf(arg); v.Kind() {
		case reflect.String:
			os.Stdout.WriteString(v.String())
		case reflect.Int:
			os.Stdout.WriteString(strconv.Itoa(10))
		}
	}
}
