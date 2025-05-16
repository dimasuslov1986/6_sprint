package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {

	// создать логгер
	// New(out io.Writer, prefix string, flag int) *Logger
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	// создать сервер с помощью вашей функции из пакета server
	serv := server.New(logger)
	// запустить его
	if err := serv.Serv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
