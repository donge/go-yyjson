package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"unsafe"
)

// #include <yyjson.h>
import "C"

type Data struct {
	Name string `json:"name"`
	Star int    `json:"star"`
	Hits []int  `json:"hits"`
}

func convertCValueToMap(cVal *C.yyjson_val) map[string]interface{} {
	result := make(map[string]interface{})

	iter := C.yyjson_obj_iter{}
	C.yyjson_obj_iter_init(cVal, &iter)

	for {
		key := C.yyjson_obj_iter_next(&iter)
		if key == nil {
			break
		}

		val := C.yyjson_obj_iter_get_val(key)
		keyStr := C.GoString(C.yyjson_get_str(key))
		result[keyStr] = convertCValueToInterface(val)
	}

	return result
}

func convertCValueToInterface(cVal *C.yyjson_val) interface{} {
	switch C.yyjson_get_type(cVal) {
	case C.YYJSON_TYPE_NULL:
		return nil
	case C.YYJSON_TYPE_BOOL:
		return C.yyjson_get_bool(cVal) != false
	case C.YYJSON_TYPE_NUM:
		if C.yyjson_is_int(cVal) {
			return int(C.yyjson_get_int(cVal))
		} else {
			return float64(C.yyjson_get_real(cVal))
		}
	case C.YYJSON_TYPE_STR:
		return C.GoString(C.yyjson_get_str(cVal))
	case C.YYJSON_TYPE_ARR:
		return convertCArrayToSlice(cVal)
	case C.YYJSON_TYPE_OBJ:
		return convertCValueToMap(cVal)
	default:
		return nil
	}
}

func convertCArrayToSlice(cVal *C.yyjson_val) []interface{} {
	result := make([]interface{}, 0)
	iter := C.yyjson_arr_iter{}
	C.yyjson_arr_iter_init(cVal, &iter)

	for {
		val := C.yyjson_arr_iter_next(&iter)
		if val == nil {
			break
		}
		result = append(result, convertCValueToInterface(val))
	}

	return result
}

func Unmarshal(data []byte, v *map[string]interface{}) error {
	cJsonStr := C.CString(string(data))
	defer C.free(unsafe.Pointer(cJsonStr))

	cDoc := C.yyjson_read(cJsonStr, C.size_t(len(data)), 0)
	if cDoc == nil {
		return fmt.Errorf("failed to read JSON")
	}
	defer C.yyjson_doc_free(cDoc)

	cRoot := C.yyjson_doc_get_root(cDoc)
	if cRoot == nil {
		return fmt.Errorf("failed to get root node")
	}
	r := convertCValueToMap(cRoot)
	*v = r
	return nil
}

func main() {
	jsonStr := []byte(`{"name":"Mash","star":4,"hits":[1,2,3,4], "a": { "b": {"c":"d"}}}`)
	var x map[string]interface{}
	var y map[string]interface{}

	// unmarshal 函数
	err := Unmarshal(jsonStr, &x)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("yyjson: %v\n", x)
	err = json.Unmarshal(jsonStr, &y)
	fmt.Printf("stdjson: %v\n", y)

	fmt.Printf("result equal %v", reflect.DeepEqual(x, y))
}
