package fixed_xor

import (
	"log"
	"testing"
)

func TestFixedXOr(t *testing.T) {
	hex1 := "1c0111001f010100061a024b53535009181c"
	hex2 := "686974207468652062756c6c277320657965"
	xorStr, err := xor(hex1, hex2)

	if err != nil {
		log.Fatal(err)
	}

	expected := "746865206b696420646f6e277420706c6179"

	if xorStr != expected {
		t.Errorf("Expected %s, got %s", expected, xorStr)
	}
}
