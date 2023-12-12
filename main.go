package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
)

func main() {
	modulus := big.NewInt(0).SetUint64(18446744069414584321)
	var modulusSmall uint64 = 18446744069414584321
	numBits := numBits(modulus)
	numBitsSmall := numBitsSmall(modulusSmall)
	fmt.Println("modulus bitLength: ", numBits)
	fmt.Println("small modulus bitLength: ", numBitsSmall)
	share := getRandomShare()
	fmt.Println(hex.EncodeToString(share[:]))
}

/*func shareToSmallFields(share [512]byte) uint64 {

}

func shareToBigFields(share [512]byte) []*big.Int {

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
