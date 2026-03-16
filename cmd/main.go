package main

import (
	"fmt"

	"github.com/HemanthKumarPavuluri/ledgermint/internal/blockchain"
	"github.com/HemanthKumarPavuluri/ledgermint/internal/wallet"
)

func main() {
	// Create a new wallet
	worldState := blockchain.NewState()
	alice := wallet.NewWallet()
	bob := wallet.NewWallet()

	// Mint some coins to Alice (Genesis)
	worldState.Mint(alice.Address, 100)
	fmt.Printf("Initial Alice Balance: %d\n", worldState.Balances[(alice.Address)])

	// Create a transaction from Alice to Bob
	tx := &blockchain.Transaction{
		Sender:    []byte(alice.Address),
		Recipient: []byte(bob.Address),
		Amount:    30,
	}
	tx.Sign(alice.PrivateKey)

	// Verify the transaction
	if !tx.Verify(alice.PublicKey) {
		fmt.Println("Transaction verification failed!")
		return
	}
	// Check if Alice has enough balance
	err := worldState.ApplyTransaction(tx)
	if !err {
		fmt.Println("Transaction failed!")
		return
	}
	fmt.Println("Transaction successful!")
	// Check Final Balances
	fmt.Printf("Alice Balance: %d | Bob Balance: %d\n",
		worldState.Balances[alice.Address],
		worldState.Balances[bob.Address])
}
