package single_byte_xor_cipher

import (
	"testing"
)

func TestDecryptSingleByteXOR(t *testing.T) {
	encryptedHex := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	decrypteds := DecryptSingleByteXOR(encryptedHex)
	expectedMostLikelyDecrypted := "Cooking MC's like a pound of bacon"
	if decrypteds[0] != expectedMostLikelyDecrypted {
		t.Errorf("Expected %s, got %s", expectedMostLikelyDecrypted, decrypteds[0])
	}
}
