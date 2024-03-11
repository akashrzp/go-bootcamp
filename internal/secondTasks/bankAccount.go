package secondTasks

import (
	"fmt"
	"sync"
)

type BankAccount struct {
	Balance int
	mu      sync.Mutex
}

func NewBankAccount(initialBalance int) *BankAccount {
	return &BankAccount{
		Balance: initialBalance,
	}
}

func (acc *BankAccount) Deposit(amount int, wg *sync.WaitGroup) {
	defer wg.Done()
	acc.mu.Lock()
	defer acc.mu.Unlock()
	acc.Balance += amount
	fmt.Printf("Deposited Rs.%d\n", amount)
}

func (acc *BankAccount) Withdraw(amount int, wg *sync.WaitGroup) {
	defer wg.Done()
	acc.mu.Lock()
	defer acc.mu.Unlock()
	if acc.Balance-amount < 0 {
		fmt.Println("Insufficient balance for withdrawal.")
		return
	}
	acc.Balance -= amount
	fmt.Printf("Withdrawn Rs.%d\n", amount)
}
