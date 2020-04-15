package utils

import (
	"fmt"
	"github.com/spf13/cast"
	"os"
	"reflect"
)

const tagKey = "env"

func ParseStruct(count int, t reflect.Type, v reflect.Value) (bool, error) {
	var envNameArr []string
	for i := 0; i < t.NumField(); i++ {
		name := t.Field(i).Tag.Get(tagKey)
		if count >= 0 {
			name = fmt.Sprintf("%s_%d", name, count)
		}
		envNameArr = append(envNameArr, name)
	}

	for index, key := range envNameArr {
		value, flag := os.LookupEnv(key)
		if !flag && index != 0 {
			return false, fmt.Errorf("key `%s` is absent", key)
		} else if !flag && index == 0 {
			return false, nil
		}

		switch t.Field(index).Type.Kind() {
		case reflect.Int:
			v.FieldByName(t.Field(index).Name).SetInt(cast.ToInt64(value))
		case reflect.String:
			v.FieldByName(t.Field(index).Name).SetString(value)
		}
	}

	return true, nil
}
