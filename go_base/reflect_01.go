package main

import (
	"fmt"
	"reflect"
)

func main() {
	elemStudy()

}

// reflect.Value.Elem()  和 reflect.Type.Elem() 理解
func elemStudy() {

	// 声明一个空结构体
	type cat struct {
	}

	// 创建cat的实例
	ins := &cat{}

	// 获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(ins)

	// 显示反射类型对象的名称和种类
	fmt.Printf("name:'%v' kind:'%v'\n", typeOfCat.Name(), typeOfCat.Kind())

	// 取类型的元素
	typeOfCat = typeOfCat.Elem()

	// 显示反射类型对象的名称和种类
	fmt.Printf("element name: '%v', element kind: '%v'\n", typeOfCat.Name(), typeOfCat.Kind())

}
