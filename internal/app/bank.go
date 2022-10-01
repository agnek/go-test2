package app

import "errors"

type Bank struct {
	balances map[string]float64
}

func NewBank() *Bank {
	return &Bank{balances: map[string]float64{}}
}

func (b *Bank) CreateUser(id string, balance float64) {
	b.balances[id] = balance
}

func (b *Bank) Transfer(from string, to string, amount float64) error {
	if b.balances[from] < amount {
		return errors.New("not enought money")
	}

	b.balances[from] -= amount
	b.balances[to] += amount

	return nil
}

func (b *Bank) GetBalance(id string) float64 {
	return b.balances[id]
}