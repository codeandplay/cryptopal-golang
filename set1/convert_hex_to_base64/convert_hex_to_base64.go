package convert_hex_to_base64

import (
	"encoding/base64"
	"encoding/hex"
	"log"
)

// Hex2Bytes Convert HEX string to byte array
func Hex2Bytes(h string) ([]byte, error) {
	byteArr := make([]byte, hex.EncodedLen(len(h)))

	n, err := hex.Decode(byteArr, []byte(h))

	if err != nil {
		log.Fatal(err)
	}

	return byteArr[:n], nil
}

// Bytes2Hex Convert byte array to HEX string
func Bytes2Hex(bs []byte) string {
	return hex.EncodeToString(bs)
}

// Hex2Base64 convert hex encoding to base64 encoding
func Hex2Base64(h string) (string, error) {
	// convert hex string to byte array
	byteArr, _ := Hex2Bytes(h)

	b64 := base64.StdEncoding.EncodeToString(byteArr)

	return b64, nil
}

