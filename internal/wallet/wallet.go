package wallet

import (
	"crypto/ecdsa"
	"log"
    "github.com/ethereum/go-ethereum/crypto"
)

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  []byte
	Address    string
}

func NewWallet() *Wallet {
	// Generate private key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	// Derive public key from private key
	publicKeyBytes := crypto.FromECDSAPub(&privateKey.PublicKey)

	// Generate Address (Last 20 bytes of Keccak256 hash of public key)
	address := crypto.PubkeyToAddress(privateKey.PublicKey).Hex()

	return &Wallet{privateKey, publicKeyBytes, address}

}
