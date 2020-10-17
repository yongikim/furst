package repository

import (
	"fmt"
	"furst/model"
	"os"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestWallet(t *testing.T) {
	userID := uint(1)
	repo1 := UserRepository{}
	user, found := repo1.FindByID(userID)
	if !found {
		username := os.Getenv("TEST_USER")
		password := os.Getenv("TEST_PW")
		hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		passwordEncrypt := string(hash)
		user = model.User{
			Name:     username,
			Password: passwordEncrypt,
		}
		repo1.SetUser(&user)
	}

	repo := WalletRepository{}
	name := "test wallet 3"
	wallet, found := repo.FindByName(name)
	if !found {
		wallet = model.Wallet{
			Name:   name,
			Amount: 0,
			UserID: uint(1),
		}
		err := repo.SetWallet(&wallet)
		if err != nil {
			t.Fatalf(err.Error())
		}
	}
	fmt.Println(wallet.ID)
}
