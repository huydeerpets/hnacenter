package tool

import (
	"bytes"
	"encoding/base64"
	"strings"
)

const (
	base64key = "f5s8w9w84fa16 " //加密用的常量
)

func EncodeB64(message string) string {
	base64Text := make([]byte, base64.StdEncoding.EncodedLen(len(message)))
	base64.StdEncoding.Encode(base64Text, []byte(message))
	base64Text = bytes.TrimRight(base64Text, "\x00")
	return string(base64Text)
}

func DecodeB64(message string) string {
	base64Text := make([]byte, base64.StdEncoding.DecodedLen(len(message)))
	base64.StdEncoding.Decode(base64Text, []byte(message))
	base64Text = bytes.TrimRight(base64Text, "\x00")
	return string(base64Text)
}

func Encodedata(data string) string {
	var dataresult string
	if len(data) < 1 {
		return dataresult
	}
	//skey := base64key
	skey := EncodeB64(base64key)
	data_string := EncodeB64(data)

	datalength := len(data_string)
	skeylength := len(skey)
	if skeylength > datalength {
		skeylength = datalength
	}
	for i := 0; i < skeylength; i++ {
		dataresult += string(data_string[i]) + string(skey[i])
	}

	if skeylength < datalength {
		dataresult = dataresult + data_string[skeylength:]
	}

	dataresult = strings.Replace(dataresult, "=", "O0O0O", -1) //改内容
	dataresult = strings.Replace(dataresult, "+", "o000o", -1) //改内容
	dataresult = strings.Replace(dataresult, "/", "oo00o", -1) //改内容
	return dataresult
}

func Decodedata(data string) string {
	var dataresult string
	if len(data) < 2 {
		return dataresult
	}
	//skey := base64key
	skey := EncodeB64(base64key)
	data = strings.Replace(data, "O0O0O", "=", -1) //改内容
	data = strings.Replace(data, "o000o", "+", -1) //改内容
	data = strings.Replace(data, "oo00o", "/", -1) //改内容

	datalength := len(data)
	skeylength := len(skey)
	if skeylength >= datalength/2 {
		skeylength = datalength/2 - 1
	}
	dataresult += data[:1]
	for i := 1; i < skeylength+1; i++ {
		if string(data[2*i-1]) == string(skey[i-1]) {
			dataresult += data[2*i : 2*i+1]
		}
	}

	if len(skey) == skeylength {
		dataresult = dataresult + data[2*skeylength+1:]
	}

	return DecodeB64(dataresult)
}
