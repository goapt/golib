package cryptoutil

import (
	"bytes"
	"errors"
)

// PKCS5补位
func PKCSPadding(text []byte, blockSize int) []byte {
	padding := blockSize - len(text)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(text, padtext...)
}

// 去除PKCS5补位
func PKCSUnPadding(text []byte) ([]byte, error) {
	length := len(text)

	if length == 0 {
		return text, nil
	}

	padtext := int(text[length-1])
	if length < padtext {
		return nil, errors.New("unpadding length error")
	}
	return text[:(length - padtext)], nil
}

// 补位方法
func Padding(text []byte, padding int) ([]byte, error) {
	switch padding {
	case NOPADDING:
		if len(text)%8 != 0 {
			return nil, errors.New("text length invalid")
		}
	case PKCS5PADDING:
		return PKCSPadding(text, 8), nil
	default:
		return nil, errors.New("padding type error")
	}

	return text, nil
}

// 去除补位方法
func UnPadding(text []byte, padding int) ([]byte, error) {
	switch padding {
	case NOPADDING:
		if len(text)%8 != 0 {
			return nil, errors.New("text length invalid")
		}
	case PKCS5PADDING:
		return PKCSUnPadding(text)
	default:
		return nil, errors.New("padding type error")
	}
	return text, nil
}
