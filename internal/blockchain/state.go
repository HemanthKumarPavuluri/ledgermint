package blockchain

type State struct {
	Balances map[string]int
}

// NewState initializes a new State with empty balances
func NewState() *State {
	return &State{Balances: make(map[string]int)}
}

// Mint creates coins out of thin air (Genesis only)
func (s *State) Mint(account string, amount int) {
	s.Balances[account] += amount
}

// ApplyTransaction validates and updates balancesfunc (s *State) UpdateBalance(account string, amount int) {
func (s *State) ApplyTransaction(tx *Transaction) bool {
	sender := string(tx.Sender)
	recipient := string(tx.Recipient)

	// Check if sender has enough balance
	if s.Balances[sender] < tx.Amount {
		return false // Insufficient funds
	}

	// Update balances
	s.Balances[sender] -= tx.Amount
	s.Balances[recipient] += tx.Amount
	return true
}

// GetBalance retrieves the balance of a given address
func (s *State) GetBalance(address string) int {
	return s.Balances[address]
}
