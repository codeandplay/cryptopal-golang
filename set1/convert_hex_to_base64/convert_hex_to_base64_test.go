package convert_hex_to_base64

import (
	"bytes"
	"log"
	"testing"
)

func TestHex2Bytes(t *testing.T) {
	hexStr := "4927"
	bArr, err := Hex2Bytes(hexStr)

	if err != nil {
		log.Fatal(err)
	}

	expected := []byte{
		73,
		39,
	}

	if !bytes.Equal(bArr, expected) {
		t.Errorf("Expect %v, got %v", expected, bArr)
	}
}

func TestBytes2Hex(t *testing.T) {
	bArr := []byte{
		73,
		39,
	}

	hexStr := Bytes2Hex(bArr)

	expected := "4927"

	if hexStr != expected {
		t.Errorf("Expect %v, got %v", expected, hexStr)
	}
}

func TestHex2Base64(t *testing.T) {
	hexStr := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	b64, err := Hex2Base64(hexStr)

	if err != nil {
		log.Fatal(err)
	}

	expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	if b64 != expected {
		t.Errorf("Expect %s, got %s", expected, b64)
	}
}
