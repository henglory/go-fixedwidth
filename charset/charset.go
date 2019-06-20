package charset

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
)

func decodeUtf8(decodeTable [][]byte, ebc []byte) []byte {
	var uni []byte
	for _, v := range ebc {
		subArr := decodeTable[v]
		for _, b := range subArr {
			uni = append(uni, b)
		}
	}
	return uni
}

func encodeUtf8(encodeTable map[string]byte, utf []byte) []byte {
	var ebc []byte
	for _, ch := range string(utf) {
		ebc = append(ebc, encodeTable[fmt.Sprintf("%c", ch)])
	}
	return ebc
}

func DecodeUTF8(codePage string, b []byte) []byte {
	switch codePage {
	case "cp838":
		return decodeUtf8(ebc838table, b)
	case "hexstring":
		return []byte(hex.EncodeToString(b))
	default:
		return defaultDecodeUtf8(codePage, b)
	}
}

func EncodeUTF8(codePage string, b []byte) []byte {
	switch codePage {
	case "cp838":
		return encodeUtf8(utf8Toebc838Map, b)
	case "hexstring":
		ret, err := hex.DecodeString(string(b))
		if err != nil {
			return b
		}
		return ret
	default:
		return defaultEncodeUtf8(codePage, b)
	}
}

func defaultDecodeUtf8(codePage string, b []byte) []byte {
	newCodeReader := bytes.NewBuffer(b)
	reader, err := charset.NewReaderLabel(codePage, newCodeReader)
	if err != nil {
		return b
	}

	nb, err := ioutil.ReadAll(reader)
	if err != nil {
		return b
	}
	return nb
}

func defaultEncodeUtf8(codePage string, b []byte) []byte {
	e, _ := charset.Lookup(codePage)
	reader := transform.NewReader(bytes.NewReader(b), e.NewEncoder())
	nb, err := ioutil.ReadAll(reader)
	if err != nil {
		return b
	}
	return nb
}
