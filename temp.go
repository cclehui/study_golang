package main

import (
        "fmt"
        "strings"
        "os"
        "bufio"
        //"strconv"
        //"io/ioutil"
        "./model/util"
)

func main() {

        string_data := "a阿点xx";

        string_reader := strings.NewReader(string_data);

        fmt.Printf("len:%d\n", len([]rune(string_data)));
        fmt.Println(string_reader.Size());

        reflect_util := new(util.ReflectUtil);
        reflect_util.ShowObject(string_reader);

        var read_in = make([]byte, 5);

        if _, e := string_reader.Read(read_in);e == nil {
                fmt.Println(read_in);
                fmt.Println(string(read_in));
        }

        file_name := "export_bill_summary.text";

        file_in,e := os.Open(file_name);
        if e != nil {
            fmt.Println(e);
            os.Exit(1);
        }

        buffer_reader := bufio.NewReader(file_in);

        //fmt.Println(file_in);
        reflect_util.ShowObject(buffer_reader);

        var buffer = make([]byte, 1024);

        for {
            buffer, _, e = buffer_reader.ReadLine();

            if e != nil {
                break;
            }

            fmt.Println(string(buffer));
        }


        //按行读取文件内容
        //文件内容排序



}


