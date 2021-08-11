package ncm

import (
	"fmt"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserInfo struct {
	username string
	userpass string
}

func (u UserInfo) GetUserInfomation() error {
	fmt.Printf("Enter username: ")
	fmt.Scan(&u.username)

	fmt.Printf("Enter password: ")
	fmt.Scan(&u.userpass)

	var confirmPass string
	fmt.Printf("Enter Confirm password: ")
	fmt.Scan(&confirmPass)

	if confirmPass != u.userpass {
		return errors.New("aaaa")
	}

	nameHash, err := bcrypt.GenerateFromPassword([]byte(u.username),12)
	if err != nil {
		return err
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(u.userpass),12)
	if err != nil {
		return err
	}

	fmt.Printf("Hash Username: %s\n", nameHash)
	fmt.Printf("Hash Userpass: %s\n", passHash)

	return nil
}
