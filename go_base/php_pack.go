package main

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

// 对应php 代码： base64_encode(pack( "a*H2", $fmID, "80"))
func main() {
	fmID := "1122331212"
	//testStr := "80"
	bytesBuffer := bytes.NewBufferString(fmID)
	appendStr, _ := hex.DecodeString("80")
	bytesBuffer.Write(appendStr)

	fmt.Println(base64.StdEncoding.EncodeToString([]byte(fmID)))
	fmt.Println(base64.StdEncoding.EncodeToString(bytesBuffer.Bytes()))

}
