package actions

import (
	"crypto/ecdsa"
	"crypto/rand"
	"log"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	sdkecdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/urfave/cli"
)

func GenerateOperatorKeyStores(ctx *cli.Context) error {

	// generate new dsca key file
	privateKey, err := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	log.Println("Public key is:", privateKey.Public())
	// Debug: Log the private key details
	log.Printf("Private key: %x\n", privateKey.D)
	log.Printf("Public key X: %x\n", privateKey.PublicKey.X)
	log.Printf("Public key Y: %x\n", privateKey.PublicKey.Y)

	// Validate the public key
	if !privateKey.PublicKey.IsOnCurve(privateKey.PublicKey.X, privateKey.PublicKey.Y) {
		log.Fatalf("Public key is not on the curve")
	}

	log.Println("Resolved PubKey", crypto.PubkeyToAddress(privateKey.PublicKey))

	err = sdkecdsa.WriteKey("tests/keys/test.ecdsa.key-3.json", privateKey, "")
	if err != nil {
		log.Fatalf("Failed to write ECDSA key: %v", err)
	}

	// generate random bls key
	blsKeyPair, err := bls.GenRandomBlsKeys()

	if err != nil {
		panic(err)
	}

	blsKeyPair.SaveToFile("tests/keys/test.bls.key-3.json", "")

	return nil
}
