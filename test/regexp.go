package main

import (
    "fmt"
    "regexp"
)

func main() {
    fmt.Println("aaaaa")
    data_str := "haaadddddddsssfdaaaaaas"
    //pattern_str := "(a+).*?(a+)"
    pattern_str := "(a+)(.+?)s"

    if ok, _ := regexp.MatchString(pattern_str, data_str); ok {
        fmt.Println("matched")
    }

    fmt.Println("xxxxxxxxxxx")

    pattern := regexp.MustCompile(pattern_str)
    fmt.Println(pattern.FindString(data_str))
    fmt.Println(pattern.FindAllString(data_str, 10))

    fmt.Println(pattern.FindStringSubmatch(data_str))
    fmt.Println(pattern.FindAllStringSubmatch(data_str, 2))

    fmt.Println(pattern.ReplaceAllString(data_str, "$1 xxxxxxxx"))
    fmt.Println(pattern.ReplaceAllString(data_str, "${1}yyyyyyyy"))
}
