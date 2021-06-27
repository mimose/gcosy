package lib

import (
	"bytes"
	"encoding/json"
	"reflect"
	"unsafe"
)

func Encode(data interface{}) ([]byte, error) {
	var jsonBuffer bytes.Buffer
	jsonEncoder := json.NewEncoder(&jsonBuffer)
	err := jsonEncoder.Encode(&data)
	if err != nil {
		return nil, err
	}
	return jsonBuffer.Bytes(), err
	//var gobBuffer bytes.Buffer
	//gobEncoder := gob.NewEncoder(&gobBuffer)
	//err = gobEncoder.Encode(jsonBuffer.Bytes())
	//if err != nil {
	//	return nil, err
	//}
	//return gobBuffer.Bytes(), err
}

func StringToSliceByte(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}
