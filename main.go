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
	numBits := numBits(modulus)
	fmt.Println("modulus bitLength: ", numBits)
	share := getRandomShare()
	fmt.Println(hex.EncodeToString(share[:]))
}

func getRandomShare() [512]byte {
	randomBytes := make([]byte, 512)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic("couldn't generate shares")
	}
	return [512]byte(randomBytes)

}

func numBits(modulus *big.Int) int {
	float, _ := modulus.Float64()
	bits := int(math.Floor(math.Log2(float)))
	return bits
}
