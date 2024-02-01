package main

import (
	"fmt"

	"github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
)

func main() {
	key, err := ecdsa.NewPrivateKeyFromHex("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ecdsa.WriteKey("incsq.ecdsa.key", key, "")
	if err != nil {
		fmt.Println(err)
		return
	}

}
