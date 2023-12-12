package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
)

func main() {
	// SNARKs will use big fields
	modulus := big.NewInt(0).SetUint64(18446744069414584321)
	// STARKs and FRI will use more performant uint64s
	var modulusSmall uint64 = 18446744069414584321
	numBits := numBits(modulus)
	numBitsSmall := numBitsSmall(modulusSmall)
	fmt.Println("modulus bitLength: ", numBits)
	fmt.Println("small modulus bitLength: ", numBitsSmall)
	share := getRandomShare()
	fmt.Println(hex.EncodeToString(share[:]))
	fmt.Println("converting to small fields:")
	fields := shareToSmallFields(share, modulusSmall)
	fmt.Println(fields)
}

func shareToSmallFields(share [512]byte, modulus uint64) [8]uint64 {
	modulusNumBits := numBitsSmall(modulus)
	if modulusNumBits > 63 {
		panic("modulusNumBits too large")
	}
	shareAsNum := new(big.Int).SetBytes(share[:])
	mask := new(big.Int).SetInt64(int64(math.Exp2(float64(modulusNumBits))) - 1)
	fmt.Println("Mask: ", hex.EncodeToString(mask.Bytes()))
	var results [8]uint64
	for i := 0; i < 8; i++ {
		masked := new(big.Int).And(shareAsNum, mask)
		if masked.BitLen() > modulusNumBits {
			fmt.Printf("masked bitlen: %v modulusNumBits %v\n", masked.BitLen(), modulusNumBits)
			panic("too big")
		}
		results[i] = masked.Uint64()
		shareAsNum = shareAsNum.Rsh(shareAsNum, uint(modulusNumBits))
	}
	return results
}

/*func shareToBigFields(share [512]byte) []*big.Int {

}*/

func getRandomShare() [512]byte {
	randomBytes := make([]byte, 512)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic("couldn't generate shares")
	}
	return [512]byte(randomBytes)

}

func numBitsSmall(modulus uint64) int {
	return int(math.Floor(math.Log2(float64(modulus))))
}

func numBits(modulus *big.Int) int {
	float, _ := modulus.Float64()
	bits := int(math.Floor(math.Log2(float)))
	return bits
}
