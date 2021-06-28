package lib

import (
	"fmt"
	"testing"
)

func TestCipherBuilder_AesEncrypt(t *testing.T) {
	builder, err := NewCipherBuilder("a123d456q789e74r")
	if err != nil {
		t.Errorf("new cipher builder error: %s\n", err)
	}
	encrypt, err := builder.AesEncrypt([]byte("{\"Key\":\"WCKWGFAWHO\",\"Name\":\"测试一下新增分组\",\"CreateTime\":\"2021-06-28 10:22:57\",\"UpdateTime"))
	if err != nil {
		t.Errorf("AesEncrypt error: %s\n", err)
	}
	fmt.Printf("after AesEncrypt: %s\n", encrypt)
	decrypt, err := builder.AesDecrypt(encrypt)
	if err != nil {
		t.Errorf("AesDecrypt error: %s\n", err)
	}
	fmt.Printf("after AesDecrypt: %s\n", string(decrypt))
}
