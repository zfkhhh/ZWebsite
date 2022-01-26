package admin

import (
	"ZWebsite/dao"
	"ZWebsite/pkg/utils"
	"errors"
	"strings"
)

func Login(accountName, accountPassword string) (result bool, err error) {
	// check accountName : 小于20字符,只能包含小写字符、数组、-、_
	if result, err = utils.IsValidAccountName(accountName); err != nil {
		return result, err
	}

	// check accountPassword
	if result, err = utils.IsValidAccountPassword(accountPassword); err != nil {
		return result, err
	}

	account, err := dao.GetAccount(accountName, accountPassword)
	if err != nil {
		result = false
		return
	}

	// compare accountName , accountPassword
	if account != nil &&
		strings.Compare(accountName, account.AccountName) == 0 &&
		strings.Compare(accountPassword, accountPassword) == 0 {
		result = true
		return
	} else {
		result = false
		return result, errors.New(" account name or password wrong")
	}
}

func GetAccount(accountName, accountPassword string) (*dao.Account, error) {
	// check accountName : 小于20字符,只能包含小写字符、数组、-、_
	if _, err := utils.IsValidAccountName(accountName); err != nil {
		return nil, err
	}

	// check accountPassword
	if _, err := utils.IsValidAccountPassword(accountPassword); err != nil {
		return nil, err
	}
	account, err := dao.GetAccount(accountName, accountPassword)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func CreateAccount(account *dao.Account) (err error) {

	// check accountName : 小于20字符,只能包含小写字符、数组、-、_
	if _, err := utils.IsValidAccountName(account.AccountName); err != nil {
		return err
	}

	// check accountPassword
	if _, err := utils.IsValidAccountPassword(account.AccountPassword); err != nil {
		return err
	}

	err = dao.CreateAccount(account)
	if err != nil {
		return err
	}
	return nil
}

func ResetPassword(accountName, oldAccountPassword, newAccountPassword string) (err error) {
	// check accountName : 小于20字符,只能包含小写字符、数组、-、_
	if _, err := utils.IsValidAccountName(accountName); err != nil {
		return err
	}

	// check old accountPassword
	if _, err := utils.IsValidAccountPassword(oldAccountPassword); err != nil {
		return err
	}

	// check new accountPassword
	if _, err := utils.IsValidAccountPassword(newAccountPassword); err != nil {
		return err
	}
	// check password, get account uid
	account, err := dao.GetAccount(accountName, oldAccountPassword)
	if err != nil {
		return
	}
	if account.Uid != "" {
		account.AccountPassword = newAccountPassword
		err := dao.UpdateAccount(account)
		if err != nil {
			return err
		}
	}
	return nil
}
