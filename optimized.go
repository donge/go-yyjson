package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"unsafe"
)

// #include <yyjson.h>
import "C"

// 优化版本：避免字符串转换，直接使用字节数据
func UnmarshalOptimized(data []byte, v *map[string]interface{}) error {
	// 直接使用字节数据，避免CString转换
	cData := (*C.char)(unsafe.Pointer(&data[0]))
	cDoc := C.yyjson_read(cData, C.size_t(len(data)), 0)
	if cDoc == nil {
		return fmt.Errorf("failed to read JSON")
	}
	defer C.yyjson_doc_free(cDoc)

	cRoot := C.yyjson_doc_get_root(cDoc)
	if cRoot == nil {
		return fmt.Errorf("failed to get root node")
	}

	// 预分配map容量
	result := make(map[string]interface{}, 8)
	convertCValueToMapOptimized(cRoot, result)
	*v = result
	return nil
}

// 优化版本：减少CGO调用，批量处理
func convertCValueToMapOptimized(cVal *C.yyjson_val, result map[string]interface{}) {
	iter := C.yyjson_obj_iter{}
	C.yyjson_obj_iter_init(cVal, &iter)

	// 预估容量减少扩容
	if C.yyjson_get_type(cVal) == C.YYJSON_TYPE_OBJ {
		size := C.yyjson_obj_size(cVal)
		if size > 0 && len(result) == 0 {
			// 确保result有足够容量
			newResult := make(map[string]interface{}, int(size))
			result = newResult
		}
	}

	for {
		key := C.yyjson_obj_iter_next(&iter)
		if key == nil {
			break
		}

		val := C.yyjson_obj_iter_get_val(key)
		// 优化字符串转换：避免GoString调用
		var keyStr string
		if C.yyjson_get_type(key) == C.YYJSON_TYPE_STR {
			keyStr = cStringToString(key)
		} else {
			keyStr = C.GoString(C.yyjson_get_str(key))
		}

		result[keyStr] = convertCValueToInterfaceOptimized(val)
	}
}

// 优化的字符串转换：减少内存分配
func cStringToString(cVal *C.yyjson_val) string {
	str := C.yyjson_get_str(cVal)
	len := C.yyjson_get_len(cVal)
	if len == 0 {
		return ""
	}

	// 直接转换，避免GoString的开销
	return C.GoStringN(str, C.int(len))
}

func convertCValueToInterfaceOptimized(cVal *C.yyjson_val) interface{} {
	valType := C.yyjson_get_type(cVal)

	switch valType {
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
		return cStringToString(cVal)
	case C.YYJSON_TYPE_ARR:
		return convertCArrayToSliceOptimized(cVal)
	case C.YYJSON_TYPE_OBJ:
		result := make(map[string]interface{}, 4)
		convertCValueToMapOptimized(cVal, result)
		return result
	default:
		return nil
	}
}

func convertCArrayToSliceOptimized(cVal *C.yyjson_val) []interface{} {
	// 预分配容量
	size := C.yyjson_arr_size(cVal)
	result := make([]interface{}, 0, int(size))

	iter := C.yyjson_arr_iter{}
	C.yyjson_arr_iter_init(cVal, &iter)

	for {
		val := C.yyjson_arr_iter_next(&iter)
		if val == nil {
			break
		}
		result = append(result, convertCValueToInterfaceOptimized(val))
	}

	return result
}

func testOptimized() {
	jsonStr := []byte(`{"name":"Mash","star":4,"hits":[1,2,3,4], "a": { "b": {"c":"d"}}}`)
	var z map[string]interface{}
	var y map[string]interface{}

	// 优化版本
	err := UnmarshalOptimized(jsonStr, &z)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("yyjson optimized: %v\n", z)

	// 标准库
	err = json.Unmarshal(jsonStr, &y)
	fmt.Printf("stdjson: %v\n", y)

	fmt.Printf("optimized == std: %v\n", reflect.DeepEqual(z, y))
}
