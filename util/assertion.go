package util

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

func AnyToInt(t1 interface{}) (t2 int) {
	switch t1.(type) {
	case uint:
		t2 = int(t1.(uint))
		break
	case int8:
		t2 = int(t1.(int8))
		break
	case uint8:
		t2 = int(t1.(uint8))
		break
	case int16:
		t2 = int(t1.(int16))
		break
	case uint16:
		t2 = int(t1.(uint16))
		break
	case int32:
		t2 = int(t1.(int32))
		break
	case uint32:
		t2 = int(t1.(uint32))
		break
	case int64:
		t2 = int(t1.(int64))
		break
	case uint64:
		t2 = int(t1.(uint64))
		break
	case float32:
		t2 = int(t1.(float32))
		break
	case float64:
		t2 = int(t1.(float64))
		break
	case string:
		t2, _ = strconv.Atoi(t1.(string))
		break
	default:
		t2 = t1.(int)
		break
	}
	return t2
}

func AnyToUint(t1 interface{}) (t2 uint) {
	switch t1.(type) {
	case int8:
		t2 = uint(t1.(int8))
		break
	case uint8:
		t2 = uint(t1.(uint8))
		break
	case int16:
		t2 = uint(t1.(int16))
		break
	case uint16:
		t2 = uint(t1.(uint16))
		break
	case int32:
		t2 = uint(t1.(int32))
		break
	case uint32:
		t2 = uint(t1.(uint32))
		break
	case int64:
		t2 = uint(t1.(int64))
		break
	case uint64:
		t2 = uint(t1.(uint64))
		break
	case float32:
		t2 = uint(t1.(float32))
		break
	case float64:
		t2 = uint(t1.(float64))
		break
	case string:
		t, _ := strconv.ParseUint(t1.(string), 10, 64)
		t2 = uint(t)
		break
	default:
		t2 = t1.(uint)
		break
	}
	return t2
}

// AnyToStr 任意类型数据转string
func AnyToStr(i interface{}) (string, error) {
	if i == nil {
		return "", nil
	}

	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return "", nil
		}
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.String:
		return v.String(), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10), nil
	case reflect.Float32:
		return strconv.FormatFloat(v.Float(), 'f', -1, 32), nil
	case reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64), nil
	case reflect.Complex64:
		return fmt.Sprintf("(%g+%gi)", real(v.Complex()), imag(v.Complex())), nil
	case reflect.Complex128:
		return fmt.Sprintf("(%g+%gi)", real(v.Complex()), imag(v.Complex())), nil
	case reflect.Bool:
		return strconv.FormatBool(v.Bool()), nil
	case reflect.Slice, reflect.Map, reflect.Struct, reflect.Array:
		str, _ := json.Marshal(i)
		return string(str), nil
	default:
		return "", fmt.Errorf("unable to cast %#v of type %T to string", i, i)
	}
}

func Decimal(num float64) float64 {
	num, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", num), 64)
	return num
}
