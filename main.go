package main

import (
	"bytes"
	"fmt"
	"log"

	"golang.org/x/crypto/ssh"
)

func main() {
	//var hostKey ssh.PublicKey
	config := &ssh.ClientConfig{
		User: "user",
		Auth: []ssh.AuthMethod{
			ssh.Password("pass"),
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
