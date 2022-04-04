/*
反射相关
*/
package language

import (
	"fmt"
	"reflect"
)

//打印配置文件的结构体,看配置是不是搞好了
func PrintStruct(obj interface{}) {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})

	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	for i, v := range data {
		fmt.Printf("%v : %v\n", i, v)
	}
}

//反射修改结构体的一个例子
func priorityConfig(nochange interface{}, willChange interface{}) {
	master_t := reflect.TypeOf(nochange).Elem()
	slave_t := reflect.TypeOf(willChange).Elem()
	master_v := reflect.ValueOf(nochange).Elem()
	slave_v := reflect.ValueOf(willChange).Elem()
	store_map := make(map[string]interface{})
	for i := 0; i < master_t.NumField(); i++ {
		store_map[master_t.Field(i).Name] = master_v.Field(i).Interface()
	}
	//两份配置文件有些许不同
	if v, ok := store_map["SYSTEM_NAME"]; ok {
		switch v {
		case "an":
			store_map["SYSTEM_NAME"] = "analysis"
		case "th":
			store_map["SYSTEM_NAME"] = "probe"
		}
	}

	for i := 0; i < slave_t.NumField(); i++ {
		if value, ok := store_map[slave_t.Field(i).Name]; ok {
			switch slave_t.Field(i).Type.String() {
			case "string":
				if value == "" {
					continue
				}
				slave_v.Field(i).SetString(value.(string))
			case "int":
				if value == 0 {
					continue
				}
				slave_v.Field(i).SetInt(int64(value.(int)))
			}
		}
	}
}
