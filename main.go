package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

type User struct {
	name     string
	password string
}

func NewUser(n string, p string) *User {
	user := new(User)
	user.name = n
	user.password = p
	return user
}

func (u *User) Input() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Name: ")
	if scanner.Scan() {
		u.name = scanner.Text()
	}
	fmt.Print("Pass: ")
	if scanner.Scan() {
		u.password = scanner.Text()
	}
}

func (u *User) Valid() bool {
	if u.name == "" || u.password == "" {
		return false
	}
	return true
}

func main() {
	//var hostKey ssh.PublicKey
	user := NewUser("", "")
	user.Input()
	if !user.Valid() {
		log.Fatal("user invalid")
	}
	config := &ssh.ClientConfig{
		User: user.name,
		Auth: []ssh.AuthMethod{
			ssh.Password(user.password),
		},
		//HostKeyCallback: ssh.FixedHostKey(hostKey),
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", "172.16.1.1:22", config)
	if err != nil {
		log.Fatal("Failed: ", err)
	}
	defer client.Close()
	// session
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Faile to create session", err)
	}
	defer session.Close()
	// run single command
	var b bytes.Buffer
	session.Stdout = &b
	cli := "show ip interface brief"
	if err := session.Run(cli); err != nil {
		log.Fatal("failed to run: " + err.Error())
	}
	fmt.Println(b.String())
}
