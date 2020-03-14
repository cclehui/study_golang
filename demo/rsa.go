package main

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

func main() {
	pemKey := []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDFRNPZ+2A/XUdIm9+VvVQQ2xWELa/TprZrAboi4R7bDDom2nUc
eIcrmI8fV+9iHfx0cVNg9YSANj7VbUOSHVzDVy5R7oo3T7ME6dS/t24I/gGvnWcz
nKHCsCek+T50siLlwfHfId2hQAKraU2gWeF2lrATprl3mC12wFdEKnz2CwIDAQAB
AoGAbXXEi+b1QBO1My/yv3bfx76ZUM+9CZcvD29U5ne+FFPTjK2ZYCPs9R7hA8Za
eTokVERxvJJfZHk1Il5PqSsLxgsHmMq/yXNvyG2Bya43sO4Vzlzx/+SvpSBuwrr9
gg150M6vJJeKsjIYFDX3/GK+TVSIuPxoL21Ho7nlBZKxNzECQQDmknu8uTCztAPF
7cpZi7nMGtg+UqdoWU2YFceHYPhqvCbhpWctfmToHJEpRpClpE8dXF1a51HmqdNE
F+N/X6J5AkEA2wYftxTomO6jVW6ntE7pCvsbBJpfyvGU0lIkyYeCRRREHR2GkkAJ
lp0pZGoi1wgNjgYA+tWlrCurhysAVrPbowJBAIamBJyxiT9oYMu1kfW5I0eOZbn/
isPlYurtzRfCCVBLkGk1rotixIrII/12uAIDcjAzQFFVxP5vLnEVgkVgFAECQQDV
962UFgEFJly6YVfEdjKEX7uNS6K5iDhzH3yAxLkm8x13tBh7V8QGN5LwXh+bImrb
jFH4ui8Xe7IeYov6J8sxAkB+LgFFbzU2UJ6FiGN/DT9Nia2dBMI28Z2GW686ZTpA
Av4fcEDNAL1Tu0gHf1VvmGb5+xNIdLA6MJF2alCMkPNk
-----END RSA PRIVATE KEY-----
`)

	data := []byte("iaaaaaaaaaaaaaaddd")
	hashMd5 := md5.Sum(data)
	hashed := hashMd5[:]

	block, _ := pem.Decode(pemKey)
	if block == nil {
		panic(errors.New("private key error"))
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("ParsePKCS1PrivateKey err", err)
		panic(err)
	}
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.MD5, hashed)
	fmt.Println("消息的签名信息： ", base64.StdEncoding.EncodeToString(signature))

}
