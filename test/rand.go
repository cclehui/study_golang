package main

import (
    "fmt"
    "time"
    "bytes"
    "math/rand"
    crand "crypto/rand"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    fmt.Println("随机整数：math/rand Intn \t", rand.Intn(100))

    fmt.Println("随机字符串: 自定义函数 \t", randStr(10))

    fmt.Println("随机字符串: /dev/urandom \t", randStr(10))
}

//从 crypto/rand  rand reader中获取
func randByRandomReader(n int) string {
    if n < 1 {
        return ""
    }

    var result = make([]byte, n);

    if _, err := crand.Read(result); err != nil {
       return "" 
    }

    return string(result)

}

//自定义随机字符串
func randStr(len int) string {
    if len <= 0 {
        return "";
    }

    const base int = 97;
    result := bytes.Buffer{};

    rand.Seed(time.Now().UnixNano())

    for i := 0; i < len; i++ {
        tempInt := base + rand.Intn(26);
        result.WriteRune(rune(tempInt));
    }

    return result.String();
}
