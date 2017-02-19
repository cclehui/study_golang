package main

import (
        "fmt"
        "strings"
        "os"
        "bufio"
        "strconv"
        //"io/ioutil"
        "./model/util"
        "sort"
)

func main() {

    file_name := "export_bill_summary.text";

    file_in,e := os.Open(file_name);
    if e != nil {
            fmt.Println(e);
            os.Exit(1);
    }

    buffer_reader := bufio.NewReader(file_in);


    curline := util.FileReadline(buffer_reader);
    fmt.Println(curline);

    fmt.Println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx");

    var buffer = make([]byte, 10);
    var raw_data sort.IntSlice;

    for {
            buffer, _, e = buffer_reader.ReadLine();

            if e != nil {
                    break;
            }

            temp_array := strings.Split(string(buffer), "\t");

            fmt.Println(temp_array[4]);

            cur_int ,e := strconv.Atoi(temp_array[4]);
            if e != nil {
                    continue;
            }

            raw_data = append(raw_data, cur_int);
            continue;

            fmt.Println(string(buffer));
            fmt.Printf("%d\t%d\t%s\n", len(buffer), cap(buffer), string(buffer));
    }

    //reflect_util := new(util.ReflectUtil);
    //reflect_util.ShowObject(new_data);

    fmt.Println(raw_data);
    sort.Sort(raw_data);
    fmt.Println(raw_data);
    //sort.Sort(sort.Reverse(raw_data));
    sort.Sort(Reverse{raw_data})

    fmt.Println(raw_data);

    //按行读取文件内容
    //文件内容排序

}

type Reverse struct {
    sort.Interface
}

func (reverse Reverse) Less(i, j int) bool {
    return reverse.Interface.Less(j, i);
}


