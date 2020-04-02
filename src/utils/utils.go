package utils

import (
	"flag"
	"fmt"
	"os"
	"path"
	"reflect"
	"runtime"
)

func PANIC_CHECK(e error) {
	if e != nil {
		panic(e)
	}
}

func FlagInit(required []string) {
	flag.Parse()
	seen := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) {
		seen[f.Name] = true
	})
	rv := true
	for _, req := range required {
		if !seen[req] {
			_, _ = fmt.Fprintf(os.Stderr, "missing required argument: -%s\n", req)
			rv = false
		}
	}
	if rv == false {
		os.Exit(2)
	}
}

func RecursiveStatField(a interface{}, indent int) {
	t := reflect.TypeOf(a)
	v := reflect.ValueOf(a)
	for i := 0; i < t.NumField(); i++ {
		fmt.Print(string(make([]byte, indent)))
		if t.Field(i).Type.Kind() != reflect.Struct {
			switch t.Field(i).Type.Kind() {
			case reflect.Uint64:
				fmt.Printf("%-15v: 0x%016x\n", t.Field(i).Name, v.Field(i))
			default:
				fmt.Printf("%-15v: %v\n", t.Field(i).Name, v.Field(i))
			}
		} else {
			fmt.Println(t.Field(i).Name)
			if v.Field(i).CanInterface() {
				RecursiveStatField(v.Field(i).Interface(), indent+4)
			}
		}
	}
}

func GetCurrentPath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

func Filter(arr interface{}, cond func(interface{}) bool) interface{} {
	contentType := reflect.TypeOf(arr)
	contentValue := reflect.ValueOf(arr)
	newContent := reflect.MakeSlice(contentType, 0, 0)
	for i := 0; i < contentValue.Len(); i++ {
		if content := contentValue.Index(i); cond(content.Interface()) {
			newContent = reflect.Append(newContent, content)
		}
	}
	return newContent.Interface()
}
