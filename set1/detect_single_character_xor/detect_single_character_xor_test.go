package detect_single_character_xor

import (
	"testing"
)

func TestDetectSingleCharXOR(t *testing.T) {
	decrypteds := DetectSingleCharXOR()

	unEncryptedFound := false

	for _, d := range decrypteds {
		if d == "Now that the party is jumping\n" {
			unEncryptedFound = true
		}
	}

	if !unEncryptedFound {
		t.Error("unencrypted text not found")
	}
}
