package main

import (
    "fmt"
    "flag"
    //"os"
    //"reflect"
)

func main() {

    defer mainDefer();

    //输入参数解析 start
    /*
    sort_type := flag.Int("T", 1, "1:正序, -1:逆序");
    input_file := flag.String("i", "", "输入文件");
    output_file := flag.String("o", "", "输出文件");
    bool_flag := flag.Bool("b", false, "bool 类型测试");
    */

    flag.Parse();
    /*

    if (*input_file == "" || *output_file == "") {
        panic("输入参数错误哦");
    }
    */

    action := flag.Arg(0);
    //输入参数解析 end

    if action == "" {
        panic("输入参数错误哦");
    }

    switch action {
        case "flag":
            testFlag();
        default:
            fmt.Println("usage is :");
            flag.PrintDefaults();

    }

    /*
    fmt.Println("action:" + action);
    fmt.Println("sort_type:" + string(*sort_type));
    fmt.Println("input_file:" + *input_file);
    fmt.Println("output_file:" + *output_file);
    fmt.Printf("bool_flag:%v\n", *bool_flag);
    */


}

func testFlag() {
    fmt.Println("test flag function called");
    
}

/*
 * main func defer
 */
func mainDefer() {
    if ex := recover();ex != nil {
        /*
        err := fmt.Errorf("%v", ex);
        fmt.Println("catch error:" + err.Error());
        */
        if message, ok := ex.(string); ok {
            fmt.Println("catch error by assert:" + message);
        }

        fmt.Println("usage is :");
        flag.PrintDefaults();
    }

}

