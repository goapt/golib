package cryptoutil

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"testing"
)


func TestAesEncrypt(t *testing.T) {
	{
		key := []byte("0123456789abcdef")
		text := []byte("23523452345234523542345234523452345234523452345234523452345")
		cipherText, err := AesEncrypt(key, text)
		if err != nil {
			t.Errorf("%v", err)
		}

		fmt.Println(hex.EncodeToString(cipherText),base64.StdEncoding.EncodeToString(cipherText))

		clearText, err := AesDecrypt(key, cipherText)
		if err != nil {
			t.Errorf("%v", err)
		}

		if bytes.Equal(clearText, text) == false {
			t.Errorf("text:%v, clearText:%v", hex.EncodeToString(text), hex.EncodeToString(clearText))
		}
		//fmt.Println("clearText:", hex.EncodeToString(clearText), "cipherText:", hex.EncodeToString(cipherText))
	}
}