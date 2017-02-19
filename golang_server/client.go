package main;

import (
    "fmt"
    "net"
    "log"
    "io"
    "bufio"
)

func main() {

    tcp_client, err := net.Dial("tcp", "localhost:8888");
    if err != nil {
        log.Println("bad connection");
        return;
    }

    reader := bufio.NewReader(tcp_client);
    writer := bufio.NewWriter(tcp_client);

    ch := make(chan string);

    // read from tcp_client by go routine
    go func () {
        for {
            readData, err := reader.ReadString('\n');
            //fmt.Println(readData);
            if err != nil || err == io.EOF {
                close(ch);
                break;
            }
            //log.Print(readData);
            ch <- readData;
        }
    }();

    go sendRequest(writer);

    for {
        select {
            case readData, ok := <-ch:
                if !ok {
                    log.Println("read from channel error: ", ok);
                    break;
                }
                log.Print(readData);

        }
    }

}

func sendRequest(writer *bufio.Writer) {
    for {
            //fmt.Println("read from stdin...");

            var input string;
            for {
                    fmt.Print(">>>");
                    fmt.Scanf("%s", &input);
                    if input != "" {
                            break;
                    }
            }

            //n, err := writer.WriteString(input + "\n");
            writer.WriteString(input + "\n");
            writer.Flush();
            //log.Println("request send to server: ", input);
            //log.Println("input length : ", n , " error: ", err);
    }
    return;
}
