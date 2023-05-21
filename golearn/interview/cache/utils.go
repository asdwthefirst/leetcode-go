package cache

import (
	"reflect"
)

func getObjectSize(val interface{}) int64 {
	return int64(reflect.TypeOf(val).Size())
}
