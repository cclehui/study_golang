package main

import(
    //"fmt"
    "log"
    "strings"
    //"errors"
    //"bufio"
    "./tcp_server"
    //"net"
    //"os"
)

func main() {
    log.Println("start main func");

    //fmt.Println(stringReverse("localhost:8888"));

    tcpServer := tcp_server.New("localhost:8888");
    tcpServer.OnNewClient(onNewClient);
    tcpServer.OnNewMessage(onNewMessage);
    tcpServer.Listen();
}

// handle new message
func onNewMessage(client *tcp_server.Client, message string) {
    message = strings.TrimRight(message, "\r\n");
    if (strings.Count(message, "") - 1) < 1 {
        client.Send("input data is empty ,please try again\n");
        return;
    }

    response := stringReverse(message);
    response = "response: " + response + "\n";
    client.Send(response);

    return;
}

//new connection
func onNewClient(client *tcp_server.Client) {
    welcomeMsg := "welcome\n";
    welcomeMsg += "useage :help, quit\n";

    err := client.Send(welcomeMsg);
    if err != nil {
        client.Close();
        log.Println(err)
        return;
    }
    /*
    writer := bufio.NewWriter(client.conn);
    n, err := writer.WriteString(welcomeMsg);
    writer.Flush();
    */

    return;
}

//字符串反转
func stringReverse(data string) string {
    dataRune := []rune(data);
    for i := 0; i < int(len(dataRune) / 2); i++ {
        dataRune[i], dataRune[len(dataRune) - i - 1] = dataRune[len(dataRune) - i - 1] , dataRune[i];
    }

    return string(dataRune);
}
