package main;
import(
    "fmt"
    "math/rand"
    "strconv"
    "bytes"
    "encoding/json"
    //"os"
)

func randStr(len int) string {
    if len <= 0 {
        return "";
    }

    const base int = 97;
    result := bytes.Buffer{};

    for i := 0; i < len; i++ {
        tempInt := base + rand.Intn(26);
        result.WriteRune(rune(tempInt));
    }

    return result.String();
}

func main() {
    fmt.Println("xxxxxx");

    data := make(map[string]interface{});

    for i := 1;i <= 10; i++ {
        data[strconv.Itoa(i)] = randStr(i);
    }

    fmt.Println(data);

    //encoder := json.NewEncoder(os.Stdout);
    //encoder.Encode(data);
    
    //fmt.Println(string(jsonString) + "---->" + error.Error());
    //fmt.Println(data);
    json_data, _ := json.Marshal(data);
    fmt.Println(string(json_data));


}

