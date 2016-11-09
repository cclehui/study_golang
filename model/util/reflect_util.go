package util

import (
    "fmt"
    "reflect"
    //"strconv"
)

type ReflectUtil struct {
    
}

func (a *ReflectUtil) Test() {

    fmt.Println("model/util -> reflect util test func");

}

func (a *ReflectUtil) ShowObject( obj interface{}) {
        value_method := reflect.ValueOf(obj);
        obj_type := value_method.Type();

        fmt.Printf("输出对象的属性和方法\t%v\n", obj);

        fmt.Println("\tMethods...");

        for i := 0;i < value_method.NumMethod(); i++ {
            fmt.Printf("\t%d\t%s\n", i, obj_type.Method(i).Name);
        }

        value_element := reflect.ValueOf(obj).Elem();
        obj_element := value_element.Type();

        fmt.Println("\tFields...");
        for i := 0;i < value_element.NumField(); i++ {
            fmt.Printf("\t%d\t%s\n", i, obj_element.Field(i).Name);

        }
}
