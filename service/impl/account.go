package impl

import (
	"ZWebsite/dao"
	"ZWebsite/pkg/logger"
	"ZWebsite/pkg/utils"
	"errors"
	"strings"
)



func (s *Service) Login(accountName, accountPassword string) (uid string,result bool, err error) {
	logger.For(s.ctx).Infof("check accountName and accountPassword ... ")
	// check accountName : 小于20字符,只能包含小写字符、数组、-、_
	if result, err = utils.IsValidAccountName(accountName); err != nil {
		return "",result, err
	}

	// check accountPassword
	if result, err = utils.IsValidAccountPassword(accountPassword); err != nil {
		return "",result, err
	}

	account, err := dao.GetAccount(accountName, accountPassword)
	if err != nil {
		result = false
		logger.For(s.ctx).Errorf("dao GetAccount err: [%v]",err)
		return "",result, err
	}

	// compare accountName , accountPassword
	if account != nil &&
		strings.Compare(accountName, account.AccountName) == 0 &&
		strings.Compare(accountPassword, accountPassword) == 0 {
		result = true
		logger.For(s.ctx).Infof("user [%v] login success ",accountName)
		return account.Uid,result,nil
	} else {
		result = false
		logger.For(s.ctx).Errorf("account name or password wrong")
		return "",result, errors.New("account name or password wrong")
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
