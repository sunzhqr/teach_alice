package main

import "flag"

// неэкспортированная переменная flagRunAddr содержит адрес и порт для запуска сервера
var flagRunAddr string

func parseFlags() {
	flag.StringVar(&flagRunAddr, "addr", ":8080", "address and port to run the server")
	flag.Parse()
}