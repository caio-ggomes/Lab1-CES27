package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println("Erro: ", err)
		os.Exit(0)
	}
}

func main() {
	Address, err := net.ResolveUDPAddr("udp", ":10001")
	CheckError(err)
	Connection, err := net.ListenUDP("udp", Address)
	CheckError(err)
	defer Connection.Close()
	fmt.Printf("server listening %s\n", Connection.LocalAddr().String())

	for {
		//Loop infinito para receber mensagem e escrever todo
		//conteúdo (processo que enviou, relógio recebido e texto)
		//na tela
		//FALTA FAZER
		message := make([]byte, 20)
		r_length, remote, err := Connection.ReadFromUDP(message[:])
		CheckError(err)
		data := strings.TrimSpace(string(message[:r_length]))
		fmt.Printf("received: %s from %s\n", data, remote)
	}
}
