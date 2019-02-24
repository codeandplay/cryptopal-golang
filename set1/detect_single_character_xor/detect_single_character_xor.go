package detect_single_character_xor

import (
	"bufio"
	"github.com/ray.vhatt/cryptopals/set1/single_byte_xor_cipher"
	"log"
	"os"
)

func DetectSingleCharXOR() []string {
	file, err := os.Open("test.file")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var results []string
	for scanner.Scan() {
		decrypteds := single_byte_xor_cipher.DecryptSingleByteXOR(scanner.Text())
		results = append(results, decrypteds[0])
	}

	return results
}