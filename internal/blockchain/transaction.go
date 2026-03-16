package blockchain

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"fmt"
	"log"
	"github.com/ethereum/go-ethereum/crypto"
)

type Transaction struct {
	Sender    []byte
	Recipient []byte
	Amount    int
	Signature []byte
}

// CreateID generates a hash of the transaction data
func (tx *Transaction) CreateID() []byte {
	data := fmt.Sprintf("%x%x%d", tx.Sender, tx.Recipient, tx.Amount)
	hash := sha256.Sum256([]byte(data))
	return hash[:]
}

// Sign uses the Private Key to sign the transaction ID
func (tx *Transaction) Sign(privateKey *ecdsa.PrivateKey) {
	hash := tx.CreateID()
	signature, err := crypto.Sign(hash, privateKey)
	if err != nil {
		log.Fatal("Failed to sign transaction:", err)
	}
	tx.Signature = signature
}

// Verify checks if the transaction signature is valid for the given public key
func (tx *Transaction) Verify(publicKey []byte) bool {
	hash := tx.CreateID()
	return crypto.VerifySignature(publicKey, hash, tx.Signature[:len(tx.Signature)-1])
}
