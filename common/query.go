package common

import (
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

func EncodeStruct(i interface{}, v *url.Values) error {
	val := reflect.ValueOf(i)
	return encodeStructWithPrefix("", val, v)
}

func encodeStructWithPrefix(prefix string, val reflect.Value, v *url.Values) error {
	switch val.Kind() {
	case reflect.Struct:
		{
			typ := val.Type()
			for index := 0; index < val.NumField(); index++ {
				tag, opts := parseTag(typ.Field(index).Tag.Get("qcloud_arg"))
				fieldVal := val.Field(index)
				if fieldVal.Kind() == reflect.Ptr {
					if fieldVal.IsNil() {
						if opts.Contains("required") {
							return errors.New(fmt.Sprintf("field %s of %s should not be nil", tag, typ))
						}
						continue
					}
				}
				p := strings.Join([]string{prefix, tag}, ".")
				if err := encodeStructWithPrefix(p, fieldVal, v); err != nil {
					return err
				}
			}
		}
	case reflect.Array, reflect.Slice:
		{
			for index := 0; index < val.Len(); index++ {
				p := strings.Join([]string{prefix, fmt.Sprint(index)}, ".")
				if err := encodeStructWithPrefix(p, val.Index(index), v); err != nil {
					return err
				}
			}
		}
	case reflect.Ptr, reflect.Interface:
		if err := encodeStructWithPrefix(prefix, val.Elem(), v); err != nil {
			return err
		}
	case reflect.String:
		v.Set(strings.TrimLeft(prefix, "."), val.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v.Set(strings.TrimLeft(prefix, "."), fmt.Sprint(val))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		v.Set(strings.TrimLeft(prefix, "."), fmt.Sprint(val))
	case reflect.Float32, reflect.Float64:
		v.Set(strings.TrimLeft(prefix, "."), fmt.Sprint(val))
	case reflect.Bool:
		v.Set(strings.TrimLeft(prefix, "."), fmt.Sprint(val))
	}
	return nil
}

type tagOptions []string

func (tOpts tagOptions) Contains(opt string) bool {
	for _, o := range tOpts {
		if o == opt {
			return true
		}
	}
	return false
}

func parseTag(tag string) (string, tagOptions) {
	parts := strings.Split(tag, ",")
	return parts[0], parts[1:]
}
