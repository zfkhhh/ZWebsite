package dao

import "gorm.io/gorm"

type Account struct {
	gorm.Model

	Uid             string `json:"uid" gorm:"not null;unique;<-:create"`
	AccountName     string `json:"account_name" gorm:"not null;unique"`
	AccountPassword string `json:"account_password" gorm:"not null;unique"`
	Email           string `json:"email" gorm:"not null;unique"`
}

func (account Account) TableName() string {
	return "account"
}

func GetAccount(accountName, accountPassword string) (account *Account, err error) {

	maps := map[string]string{
		"account_name":     accountName,
		"account_password": accountPassword,
	}
	err = DB.Where(maps).Take(account).Error
	if err != nil {
		return
	}
	return
}

func CreateAccount(account *Account) (err error) {
	if err = DB.Create(account).Error; err != nil {
		return
	}
	return
}

func UpdateAccount(account *Account) (err error) {
	maps := map[string]interface{}{
		"account_name":     account.AccountName,
		"account_password": account.AccountPassword,
	}
	if err = DB.Model(&Account{}).Where("uid = ?", account.Uid).Updates(maps).Error; err != nil {
		return
	}
	return
}
