package lib

import "errors"

type CipherBuilder struct {
	key []byte
}

var emptyCipherKeyError = errors.New("key cant be empty")
var illegalCipherKeyError = errors.New("key is illegal")

func NewCipherBuilder(content string) (*CipherBuilder, error) {
	if content == "" {
		return nil, emptyCipherKeyError
	}
	if len(content)%16 != 0 {
		return nil, illegalCipherKeyError
	}
	return &CipherBuilder{key: []byte(content)}, nil
}
