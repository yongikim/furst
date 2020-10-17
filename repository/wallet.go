package repository

import (
	"furst/model"
)

type WalletRepository struct {}

func (WalletRepository) SetWallet(wallet *model.Wallet) error {
  result := db.Create(&wallet)
  if result.Error != nil {
    return result.Error
  }
  return nil
}

func (WalletRepository) GetWalletList() []model.Wallet {
  wallets := make([]model.Wallet, 0)
  result := db.Limit(10).Find(&wallets)
  if result.Error != nil {
    panic(result.Error)
  }
  return wallets
}

func (WalletRepository) UpdateWallet(newWallet *model.Wallet) error {
  result := db.Save(&newWallet)
  if result.Error != nil {
    return result.Error
  }
  return nil
}

func (WalletRepository) DeleteWallet(id int) error {
  result := db.Delete(&model.Wallet{}, id)
  if result.Error != nil {
    return result.Error
  }
  return nil
}

func (WalletRepository) FindByName(name string) (model.Wallet, bool) {
  wallet := model.Wallet{}
  result := db.Where("name = ?", name).First(&wallet)
  if result.Error != nil {
    return wallet, false
  }
  return wallet, true
}
