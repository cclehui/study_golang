package main

import (
        "fmt"
        "strings"
        //"os"
        //"bufio"
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

        var new_string = make([]byte, 5);

        var e interface{};

        for {
            n := 0;
            if n, e = string_reader.Read(read_in);e == nil {
                fmt.Println(read_in, e);
                fmt.Println(string(read_in));
                new_string = append(new_string, read_in...);
            }

            if n < 1 {
                break;
            }

        }

        fmt.Println(string(new_string));


}


