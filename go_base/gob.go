package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// gob 私有数据不会encode 和decode

type GobTestStruct struct {
	Age  int
	Name string

	money int

	bankInfo *GobBankInfo
}

type GobBankInfo struct {
	Code string
}

func main() {

	sourceData := &GobTestStruct{
		Age:   18,
		Name:  "cclehui",
		money: 10000,
		bankInfo: &GobBankInfo{
			Code: "10002123",
		},
	}

	fmt.Printf("sourceData, %#v\n", sourceData)

	encodedBytes, _ := gobEncode(sourceData)

	decodedData := GobTestStruct{}
	_ = gobDecode(encodedBytes, &decodedData)

	fmt.Printf("decodedData, %#v\n", decodedData)
}

func gobEncode(data interface{}) ([]byte, error) {
	var b bytes.Buffer
	encoder := gob.NewEncoder(&b)

	err := encoder.Encode(data)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func gobDecode(b []byte, ptr interface{}) error {
	decoder := gob.NewDecoder(bytes.NewReader(b))

	err := decoder.Decode(ptr)
	if err != nil {
		return err
	}

	return nil
}
