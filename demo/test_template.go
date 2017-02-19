package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "html/template"
    //"text/template"
)

func main() {

    var data = make(map[string] interface{});

    //var user_list = make([]map[string] interface{}, 10)
    //var user_list = make([]interface{}, 10)
    var user_list []interface{}

    user_1 := make(map[string] interface{})
    user_1["name"] = "cclehui"
    user_1["age"] = "28"
    user_1["birthday"] = "2010-02-26"

    user_list = append(user_list, user_1)

    var test_data = make(map[string] string)
    test_data["name"] = "11111111"
    fmt.Println(test_data)

    //模板数据
    data["user_list"] = user_list
    data["test_data"] = test_data

    /*
    for _, user := range user_list {
        if new_user, ok := user.(map[string] interface{});ok {
            fmt.Println(new_user["name"])    
        }
    }
    */

    //自定义模板函数
    var func_map = make(template.FuncMap);
    func_map["isRightType"] = IsRightType
    func_map["ccTest"] = Test

    //模板渲染
    var tmpl_name = "./tmpl/index.tpl"
    tmpl_content_bytes, er := ioutil.ReadFile(tmpl_name)
    if er != nil {
        os.Exit(0) 
    }

    tmpl := template.New("test").Funcs(func_map)//添加模板函数
    tmpl = template.Must(tmpl.Parse(string(tmpl_content_bytes)))
    tmpl.Execute(os.Stdout, data)

    for _, st := range tmpl.Templates() {
        fmt.Println("out template:template_name \t" + st.Name()) 
    }

    //tmpl.ExecuteTemplate(os.Stdout, "test", data)

    //tmpl, _ := template.ParseFiles("./tmpl/index.tpl")
    //tmpl.Execute(os.Stdout, data)
}

func Test(str string) string {
    return "input is " + str
}

func IsRightType(user interface{}) bool {
    if _, ok := user.(map[string] interface{});ok {
        return true
    }

    return false
}
