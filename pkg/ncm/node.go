package ncm

import (
	"os"
	"fmt"
	"errors"
	"gopkg.in/yaml.v2"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/bcrypt"
)

type UserInfo struct {
	username string
	userpass string
}

func (u UserInfo) InputUserInfomation() error {
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

type resources struct {
	Node      []nodeResource `yaml:"node"`
	Profile   []profileValue `yaml:"profile"`
	Namespace []string       `yaml:"namespace"`
}

type nodeResource struct {
	Master []nodeValue `yaml:"master"`
	Worker []nodeValue `yaml:"worker"`
}

type nodeValue struct {
	Name        string `yaml:"name"`
	Ipv4Address string `yaml:"Ipv4Address"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	Profile     string `yaml:"profile"`
}

type profileValue struct {
	Name         string `yaml:"name"`
	Port         string `yaml:"port"`
	Identify_key string `yaml:"identify_key"`
}

func (r resources) displayCache() {
	fmt.Println(r)
}

type configPath struct {
	path string
}

func (c configPath) YamlLoadUserInfomation() resources {
	resources := resources{}
	b, _ := os.ReadFile(c.path)
	yaml.Unmarshal(b, &resources)
	return resources
}

func SetConfigPath() configPath {
	c := configPath{path: "./.ncm.yaml"}
	return c
}
