/*
=== Утилита telnet ===
Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port
go-telnet mysite.ru 8080
go-telnet --timeout=3s 1.1.1.1 123
Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).
При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/


package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"

	flag "github.com/spf13/pflag"
)

func main() {

	// Обработка аргументов командной строки.
	timeout := flag.DurationP("timeout", "t", 10*time.Second, "Timeout for connection")

	host := flag.StringP("host", "h", "", "Hostname or IP address")

	port := flag.StringP("port", "p", "", "Port number")

	flag.Parse()

	if *host == "" || *port == "" {
		fmt.Println("Error: host and port must be specified")
	} else {

		// Подключение к серверу с указанным таймаутом.
		conn, err := net.DialTimeout("tcp", *host+":"+*port, *timeout)

		if err != nil { // Ошибка подключения - завершение программы.
			fmt.Println(err) // Выводим ошибку соединения
			os.Exit(1)       // Завершаем программу
		} else { // Соединения установлено - чтения/запись

			reader := bufio.NewReader(os.Stdin) // Reader STDIN

			for { // Read/Write loop

				fmt.Print("-> ") // Print prompt

				text, _ := reader.ReadString('\n') // Read line from STDIN

				fmt.Fprintf(conn, text+"\n") // Write line to socket

				message, _ := bufio.NewReader(conn).ReadString('\n') // Read line from socket

				fmt.Print("<- " + message) // Print line from socket

			}

		}

	}
}
