package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
)

//利用json转化返回map
func ToMapByJson(in interface{}) (map[string]interface{}, error) {

	b, _ := json.Marshal(&in)
	var m map[string]interface{}
	err := json.Unmarshal(b, &m)
	//for k, v := range m{
	//	fmt.Printf("key:%v value:%v value type:%T\n", k, v, v)
	//}
	return m, err
}

//利用反射返回单层map
func ToMapByReflect1(in interface{}, tagName string) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct { // 非结构体返回错误提示
		return nil, fmt.Errorf("ToMap only accepts struct or struct pointer; got %T", v)
	}

	t := v.Type()
	// 遍历结构体字段
	// 指定tagName值为map中key;字段值为map中value
	for i := 0; i < v.NumField(); i++ {
		fi := t.Field(i)
		if tagValue := fi.Tag.Get(tagName); tagValue != "" {
			out[tagValue] = v.Field(i).Interface()
		}
	}
	return out, nil
}

// ToMap2 使用反射将结构体转为单层map
func ToMapByReflect2(in interface{}, tag string) (map[string]interface{}, error) {

	// 当前函数只接收struct类型
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr { // 结构体指针
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("ToMap only accepts struct or struct pointer; got %T", v)
	}

	out := make(map[string]interface{})
	queue := make([]interface{}, 0, 1)
	queue = append(queue, in)

	for len(queue) > 0 {
		v := reflect.ValueOf(queue[0])
		if v.Kind() == reflect.Ptr { // 结构体指针
			v = v.Elem()
		}
		queue = queue[1:]
		t := v.Type()
		for i := 0; i < v.NumField(); i++ {
			vi := v.Field(i)
			if vi.Kind() == reflect.Ptr { // 内嵌指针
				vi = vi.Elem()
				if vi.Kind() == reflect.Struct { // 结构体
					queue = append(queue, vi.Interface())
				} else {
					ti := t.Field(i)
					if tagValue := ti.Tag.Get(tag); tagValue != "" {
						// 存入map
						out[tagValue] = vi.Interface()
					}
				}
				break
			}
			if vi.Kind() == reflect.Struct { // 内嵌结构体
				queue = append(queue, vi.Interface())
				break
			}
			// 一般字段
			ti := t.Field(i)
			if tagValue := ti.Tag.Get(tag); tagValue != "" {
				// 存入map
				out[tagValue] = vi.Interface()
			}
		}
	}
	return out, nil
}
