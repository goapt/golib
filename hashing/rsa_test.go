package hashing

import (
	"crypto"
	"fmt"
	"testing"
)

//openssl
//genrsa -out rsa_private_key.pem 2048
//rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem

var privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEAwC6IUIE1SV2ssBh64VP149Ta/HO1L/EJf0eKdJGrfl1+MFT/
16WLYcrX8SG+ttcBgJ0kI4nNPbTp/V/salFmsluzrPpbQrz9jO/XzNZO8fbwuEcn
D7WPdUh0qs2l9o4aYjihvCfR2k0zdLi9N5qdXswuWXjlNLkY9qdSAY50MYa0xS73
V7hVa4qlardRUFFp+sJR8outn0G8K5H5zxEdgvzt7NBuT0IAe+S/mRrVwv6dQ7sN
5gmIRZD03ul1W/ifcEbBYwPVQe9wgxoKFhz5PIY3FLv9+W9OmNIGG4j98X412jks
9/02GnK+kJYzJUCDdE4eeCbIm6i1zyCd/M9gTQIDAQABAoIBABV03IlIc7+WMtKS
WR3RNvHR8QUgkgkhOzM3tJChNulr3MveoZXdCLvJXuSwJM/bH3LRmJhTKVQLX9iH
Hikcn3+oV14nsYq4+QIEQS7Aep3vOR8J6qWJWtP4W/Y58Z6ebPmsYvpT480gs28D
tTSBEbLkzwP0SHrjc5MOSUydMAHbW09jEjcfJvKBWWSBiVvW5rt0CVPycwB9SSv9
iaH/040YgW5nQYuiULzGYJIL0p8K978szLsfwrVemjeClcnI39I77JRhnWd4oAld
5qrJ9EHhhr+ZJMuTFUm5CDr1qgjoKSJ3bnSGr5M1Ycnbd/pE5Rzp53dSBvAh3VRb
Y7Ghcs0CgYEA/uALowSEUDeb9GfC5m90K6Wx1B04agxCWgJlykK08B0cUkzDGN2W
ooBxRv9X955FfYtqKfVMGh7vLKar7SPuvm6QGFyISbeGlJT1UYpiH2DygF+rfKV+
fzSA0vMHGQwH/KQ6CHrn+3HfDTRrSLgTKBQLVYGCXYc4Hiq229MKabMCgYEAwQeo
J24PdNvBNj9GXKAhEDT9SpPxO2ykSgKXkvZstPmyDoaXrJm8Agh/SN0P/aLviTxf
csKTK/xkShZqMzj1xnKgIpgyXBhiyeCSFs/K+nitIJLZ13nNBfGDelJ15M5u+s6g
DPGep5On/EJErATQu6HKfIJLDWn9bxWkFHymDf8CgYEA10nsynCASzoN0+7ppLg0
SsIVafScuIdObLVlwemC6OfOCn6otZJHMVCJXs1FQvZAAIII1RRMiivjH7ZRt2gI
vHar5MgUnyAU3+DLL2tS4uqDLIijVBB+v3hb8NQ0BKYzAOa/1nXrCmvvWzoR/UTv
eYUk5b8VnWcgseqmtxyWyR0CgYEAqf3HPWo0nHIHFnHk0h+G3pH72litIesMR80+
lQMFOt+GyjoHis4cfyHijlV4BqMeFhqf6B2opBzyaTiSMCficfByS+UCvI4ROb3W
idZW5/usY7pPs+4k+y303p2OC4EsxR2AX8XNNcDYOFRXy7G17PePrdTEqbyEnvZM
+GnJhxsCgYEAz/smuM0+xtqr7yojEkQWmFuspeH4WXsW2iU+gCdjdhyQSW1z1M/3
cTMy2jvdU8WslmqMjOCATwhidDmAB8ih0RixOVvJA4Kk0RSaTvpa2EhWfYtgrVS6
8EP/Z7cPabasL7F+3Q/S7AOKHNy705RViER6d+wLA7EC3ibZo0S7VVQ=
-----END RSA PRIVATE KEY-----`

var publicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAwC6IUIE1SV2ssBh64VP1
49Ta/HO1L/EJf0eKdJGrfl1+MFT/16WLYcrX8SG+ttcBgJ0kI4nNPbTp/V/salFm
sluzrPpbQrz9jO/XzNZO8fbwuEcnD7WPdUh0qs2l9o4aYjihvCfR2k0zdLi9N5qd
XswuWXjlNLkY9qdSAY50MYa0xS73V7hVa4qlardRUFFp+sJR8outn0G8K5H5zxEd
gvzt7NBuT0IAe+S/mRrVwv6dQ7sN5gmIRZD03ul1W/ifcEbBYwPVQe9wgxoKFhz5
PIY3FLv9+W9OmNIGG4j98X412jks9/02GnK+kJYzJUCDdE4eeCbIm6i1zyCd/M9g
TQIDAQAB
-----END PUBLIC KEY-----`

func TestRSASign(t *testing.T) {
	data := "hello"
	str, err := RSASign([]byte(data), []byte(privateKey), crypto.SHA1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(str)
	str, err = RSASign([]byte(data), []byte(privateKey), crypto.SHA256)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(str)
}

func TestRSAVerify(t *testing.T) {
	data := "hello"
	str2, err := RSASign([]byte(data), []byte(privateKey), crypto.SHA256)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(str2)
	if err := RSAVerify([]byte(data), str2, []byte(publicKey), crypto.SHA256); err != nil {
		t.Error(err)
	}
}

func TestRSAEncrypt(t *testing.T) {
	data := "hello"
	str, err := RSAEncrypt([]byte(data), []byte(publicKey))
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(str))
}

func TestRSADecrypt(t *testing.T) {
	data := "hello"
	str, err := RSAEncrypt([]byte(data), []byte(publicKey))
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(str))
	ori, err := RSADecrypt(str, []byte(privateKey))
	if err != nil {
		t.Error(err)
	}

	if data != string(ori) {
		t.Error(fmt.Sprintf("decrypt error must get %s,but get %s", data, string(ori)))
	}

	fmt.Println(string(ori))
}
