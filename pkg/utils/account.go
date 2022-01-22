package utils

import (
	"ZWebsite/pkg/constant"
	"github.com/pkg/errors"
	"regexp"
	"unicode"
)

func IsValidAccountName(name string) (bool, error) {
	if len(name) > constant.AccountNameMaxLength || len(name) < constant.AccountNameMinLength {
		return false, errors.New("用户名长度在1到20之间")
	}

	for _, letter := range name {
		if unicode.IsLetter(letter) || unicode.IsNumber(letter) || string(letter) == "-" || string(letter) == "_" {
			continue
		}
		return false, errors.New("用户名只能包含小写字母、数字、下划线、横线")
	}

	return true, nil
}

func IsValidAccountPassword(password string) (bool, error) {
	// 密码长度在8到20之间且需包含至少一个大写字符，一个小写字符和一个数字
	if len(password) < constant.AccountPasswordMinLenth || len(password) > constant.AccountPasswordMaxLenth {
		return false, errors.New("密码长度在8到20之间!")
	}
	num := `[0-9]{1}`
	a_z := `[a-z]{1}`
	A_Z := `[A-Z]{1}`
	if b, err := regexp.MatchString(num, password); !b || err != nil {
		return false, errors.New("密码必须包含数字!")
	}
	if b, err := regexp.MatchString(a_z, password); !b || err != nil {
		return false, errors.New("密码必须包含小写字母!")
	}
	if b, err := regexp.MatchString(A_Z, password); !b || err != nil {
		return false, errors.New("密码必须包含大写字母!")
	}

	return true, nil
}
