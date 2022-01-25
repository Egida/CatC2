package main

import (
	"github.com/gliderlabs/ssh"
	"golang.org/x/crypto/ssh/terminal"
	"math/rand"
	"strings"
	"time"
)

type Admin struct {
	conn ssh.Session
}

func NewAdmin(ssh ssh.Session) *Admin {
	return &Admin{ssh}
}

func (this *Admin) Handle() {
	this.conn.Write([]byte("\033[?1049h"))
	this.conn.Write([]byte("\xFF\xFB\x01\xFF\xFB\x03\xFF\xFC\x22"))

	defer func() {
		this.conn.Write([]byte("\033[?1049l"))
	}()

	var userInfo AccountInfo
	userInfo = database.GetAccountInfo(this.conn.User())

	this.conn.Write([]byte("\r\n\033[0m"))

	this.conn.Write([]byte("\033[2J\033[1;1H"))
	this.SendMessage("\033[97mLoading cnc...", true)
	time.Sleep(500 * time.Millisecond)

	go func() {
		i := 0
		for {
			time.Sleep(time.Second)
			this.SetTitle("yes")
			i++
		}
	}()
	this.conn.Write([]byte("\033[2J\033[1;1H"))
	for {
		term := terminal.NewTerminal(this.conn, userInfo.username+"> ")
		cmd, err := term.ReadLine()
		cmd = strings.ToLower(cmd)
		if err != nil {
			return
		}

		if cmd == "clear" || cmd == "cls" || cmd == "c" || cmd == "CLEAR" || cmd == "CLS" {
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			continue
		}
		if cmd == "help" || cmd == "HELP" || cmd == "?" {
			continue
		}

		if err != nil || cmd == "exit" || cmd == "quit" || cmd == "logout" {
			return
		}
	}
}

func (this *Admin) SendMessage(message string, newline bool) {
	if newline {
		this.conn.Write([]byte(message + "\r\n"))
	} else {
		this.conn.Write([]byte(message))
	}
}

func (this *Admin) ClearScreen() {
	this.conn.Write([]byte("\033[2J\033[1;1H"))
}

func (this *Admin) SetTitle(message string) {
	this.conn.Write([]byte("\033]0;" + message + "\007"))
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
