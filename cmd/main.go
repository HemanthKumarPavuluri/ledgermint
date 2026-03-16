package main

import (
	"fmt"

	"github.com/HemanthKumarPavuluri/ledgermint/internal/blockchain"
	"github.com/HemanthKumarPavuluri/ledgermint/internal/wallet"
)

func main() {
	// Create a new wallet
	alice := wallet.NewWallet()
	bob := wallet.NewWallet()

	// Create a transaction from Alice to Bob
	tx := &blockchain.Transaction{
		Sender:    []byte(alice.Address),
		Recipient: []byte(bob.Address),
		Amount:    10,
	}
	// Alice signs the transaction
	tx.Sign(alice.PrivateKey)
	fmt.Println("Transaction Signed by Alice")

	// Network Verifies the Transaction
	isValid := tx.Verify(alice.PublicKey)
	if isValid {
		fmt.Println("Transaction Verified Successfully")
	} else {
		fmt.Println("Transaction Verification Failed")
	}

	// Create a new block with the transaction
	block := blockchain.NewBlock(string(tx.CreateID()), []byte{})
	fmt.Printf("New Block Created: %x\n", block.Hash)
}
