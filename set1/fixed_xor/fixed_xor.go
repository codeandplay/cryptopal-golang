package fixed_xor

import (
	"github.com/ray.vhatt/cryptopals/set1/convert_hex_to_base64"
	"errors"
)

func xor(hex1 string, hex2 string) (string, error) {
	hex1Bytes, _ := convert_hex_to_base64.Hex2Bytes(hex1)
	hex2Bytes, _ := convert_hex_to_base64.Hex2Bytes(hex2)

	if len(hex1Bytes) != len(hex2Bytes) {
		return "", errors.New("hex strings must be equal length")
	}

	var hex3Bytes []byte

	for n := range hex1Bytes {
		hex3Bytes = append(hex3Bytes, hex1Bytes[n] ^ hex2Bytes[n])
	}

	return convert_hex_to_base64.Bytes2Hex(hex3Bytes), nil
}

