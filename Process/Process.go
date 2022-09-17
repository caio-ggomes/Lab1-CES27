package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

// Variáveis globais interessantes para o processo
var err string
var myPort string          //porta do meu servidor
var nServers int           //qtde de outros processo
var CliConn []*net.UDPConn //vetor com conexões para os servidores dos outros processos

var ServConn *net.UDPConn //conexão do meu servidor (onde recebo mensagens dos outros processos)

func CheckError(err error) {
	if err != nil {
		fmt.Println("Erro: ", err)
		os.Exit(0)
	}
}

func PrintError(err error) {
	if err != nil {
		fmt.Println("Erro: ", err)
	}
}

func doServerJob() {
	//Loop infinito mesmo
	for {
		//Ler (uma vez somente) da conexão UDP a mensagem
		//Escrever na tela a msg recebida (indicando o
		//endereço de quem enviou)
		//FALTA ALGO AQUI
	}
}
func doClientJob(otherProcess int, i int) {
	//Enviar uma mensagem (com valor i) para o servidor do processo
	//otherServer.
	//FALTA ALGO AQUI
	_, err := CliConn[otherProcess].Write(buf)
	//FALTA ALGO AQUI
}

// ESSA FUNÇÃO ESTÁ PRONTA
func initConnections() {
	myPort = os.Args[1]
	nServers = len(os.Args) - 2
	/*Esse 2 tira o nome (no caso Process) e tira a primeira porta (que
	é a minha). As demais portas são dos outros processos*/
	CliConn = make([]*net.UDPConn, nServers)

	/*Outros códigos para deixar ok a conexão do meu servidor (onde re-
	cebo msgs). O processo já deve ficar habilitado a receber msgs.*/

	ServerAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1"+myPort)
	CheckError(err)
	ServConn, err = net.ListenUDP("udp", ServerAddr)
	CheckError(err)

	/*Outros códigos para deixar ok a minha conexão com cada servidor
	dos outros processos. Colocar tais conexões no vetor CliConn.*/
	for servidores := 0; servidores < nServers; servidores++ {
		ServerAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1"+os.Args[2+servidores])
		CheckError(err)
		Conn, err := net.DialUDP("udp", nil, ServerAddr)
		CliConn[servidores] = Conn
		CheckError(err)
	}
}

// ESSA FUNÇÃO ESTÁ PRONTA
func main() {
	initConnections()
	//O fechamento de conexões deve ficar aqui, assim só fecha
	//conexão quando a main morrer
	defer ServConn.Close()
	for i := 0; i < nServers; i++ {
		defer CliConn[i].Close()
	}

	/*Todo Process fará a mesma coisa: ficar ouvindo mensagens e man-
	dar infinitos i’s para os outros processos*/

	go doServerJob()
	i := 0
	for {
		for j := 0; j < nServers; j++ {
			go doClientJob(j, i)
		}
		// Espera um pouco
		time.Sleep(time.Second * 1)
		i++
	}
}
