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


}


