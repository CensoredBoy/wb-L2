package main

import (
	"fmt"
	"log"
)

/*
Паттерн «Фасад» предоставляет унифицированный интерфейс вместо набора интерфейсов некоторой подсистемы.
Фасад определяет интерфейс более высокого уровня, который упрощает использование подсистемы.
*/


type User struct{
	name string
	password string
}

func NewUser(name, password string) *User{
	return &User{
		name: name,
		password: password,
	}
}

func (user User) Auth(name, password string) error{
	if user.name == name && user.password == password{
		return nil
	}
	return fmt.Errorf("incorrect credentials")
}

type Wallet struct{
	balance int
}

func NewWallet() *Wallet{	
	return &Wallet{
		balance: 0,
	}
}

func (wallet *Wallet) AddMoney(amount int) {
	wallet.balance += amount
}

type Account struct{
	wallet *Wallet
	user *User
}

func NewAccount(username, password string) *Account{
	return &Account{
		wallet: NewWallet(),
		user: NewUser(username, password),
	}
} 

func (account *Account) AddMoney(username, password string, amount int) error{
	if err := account.user.Auth(username, password); err != nil{
		return err
	}
	account.wallet.AddMoney(amount)
	return nil
}


func main() {
	account := NewAccount("admin", "admin")
	if err := account.AddMoney("admin", "admin", 12); err != nil{
		log.Fatalf("Error: %s", err)
	}
	fmt.Printf("Balance of %s is: %d", account.user.name, account.wallet.balance)
}
