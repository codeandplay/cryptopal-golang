package single_byte_xor_cipher

import (
	"github.com/ray.vhatt/cryptopals/set1/convert_hex_to_base64"
	"io/ioutil"
	"log"
	"sort"
)

func getFrequency(bs []byte) map[byte] int {
	var charFreq = make(map[byte] int)

	for _, b := range bs {
		if v, exist := charFreq[b]; exist {
			charFreq[b] = v + 1
		} else {
			charFreq[b] = 1
		}
	}

	return charFreq
}

type Pair struct {
	Key byte
	Value int
}
type PairList []Pair

func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func getCharFreqRank(cf map[byte] int) PairList {
	pl := make(PairList, len(cf))

	i := 0

	for k, v := range cf {
		pl[i] = Pair{k, v}
		i ++
	}

	sort.Sort(sort.Reverse(pl))
	return pl
}

func decrypt(enBS []byte, ci byte) string {
	var deBS []byte
	for _, b := range enBS {
		deBS = append(deBS, b ^ ci)
	}
	return string(deBS)
}

func DecryptSingleByteXOR(enHex string) []string {
	// get sample text character frequency
	bs, err := ioutil.ReadFile("sample.txt")

	if err != nil {
		log.Fatal(err)
	}

	// get byte frequency of the sample message
	sampleCharFreq := getFrequency(bs)
	sampleCharRank := getCharFreqRank(sampleCharFreq)

	// get byte frequency of the encrypted message
	enBs, _ := convert_hex_to_base64.Hex2Bytes(enHex)
	enCharFreq := getFrequency(enBs)
	topEnCharRank := getCharFreqRank(enCharFreq)[0]

	var possibleDecrypted []string
	// using the first 5 top ranked characters in character frequency to compare with highest ranked byte in the
	// encrypted message. in other word, character xor encrypted-character. work out the cipher. decrypt the message
	// with the deduce cipher
	for i, sr := range sampleCharRank {
		if i >= 5 {
			break
		}

		// get xor cipher
		ci := sr.Key ^ topEnCharRank.Key
		//ciHex := convert_hex_to_base64.Bytes2Hex([]byte{ci})
		//fmt.Printf("Decrypted with cipher %s\n", ciHex)
		decrypted := decrypt(enBs, ci)
		//fmt.Printf("Decrypted message is %s\n", decrypted)
		//fmt.Print("=============================\n\n")
		possibleDecrypted = append(possibleDecrypted, decrypted)
	}

	return possibleDecrypted
}
