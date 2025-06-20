package support

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type StructSupport struct {
}

func NewStructSupport() *StructSupport {
	return &StructSupport{}
}

func (s *StructSupport) Fill(in any, data map[string]string, tagName string) error {
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		panic(fmt.Errorf("ToMap only accepts structs; got %T", v))
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// gets us a StructField
		fi := typ.Field(i)
		tagv := fi.Tag.Get(tagName)
		if strings.Contains(tagv, ",") {
			tagv = strings.TrimSpace(strings.Split(tagv, ",")[0])
		}

		switch tagv {
		case "-":
			continue
		case "":
			tagv = fi.Name
		default:
			if val, ok := data[tagv]; ok {
				switch fi.Type.Kind() {
				case reflect.String:
					v.FieldByName(fi.Name).SetString(val)
					break
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					valInt, err := strconv.ParseInt(val, 10, 64)
					if err != nil {
						return err
					}

					v.FieldByName(fi.Name).SetInt(valInt)
					break
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					valInt, err := strconv.ParseUint(val, 10, 64)
					if err != nil {
						return err
					}

					v.FieldByName(fi.Name).SetUint(valInt)
					break
				case reflect.Float32, reflect.Float64:
					valFloat, err := strconv.ParseFloat(val, 64)
					if err != nil {
						return err
					}

					v.FieldByName(fi.Name).SetFloat(valFloat)
					break
				case reflect.Bool:
					valBool, err := strconv.ParseBool(val)
					if err != nil {
						return err
					}

					v.FieldByName(fi.Name).SetBool(valBool)
					break
				default:
					break
				}
			}
		}
	}

	return nil
}
