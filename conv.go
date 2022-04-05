package clitable

import (
	"fmt"
	"reflect"

	"github.com/go-yaml/yaml"
)

func toString(v interface{}) string {
	if v == nil {
		return "nil"
	}
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Map, reflect.Slice:
		line, err := yaml.Marshal(v)
		if err != nil {
			return ""
		}
		return string(line)
	default:
		return fmt.Sprintf("%v", v)
	}
}
