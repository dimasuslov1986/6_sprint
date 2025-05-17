package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

// 1 Создайте структуру сервера с полями для логгера (log.Logger) и http-сервера (http.Server).
type Server struct {
	Logger *log.Logger
	Serv   *http.Server
}

// 2 Создайте функцию, в которой нужно создать http-роутер. Функция принимает log.Logger и возвращает экземпляр структуры вашего сервера.
func New(logger *log.Logger) *Server {
	mux := http.NewServeMux()

	// 3 Зарегистрируйте ваши хендлеры в http-роутере.
	mux.HandleFunc("/", handlers.MainHandle)
	mux.HandleFunc("/upload", handlers.UploadHandle)

	// 4 Создайте экземпляр структуры http.Server. Для настойки вашего сервера используйте следующие поля:
	// Addr — используйте порт 8080.
	// Handler — передайте ваш http-роутер.
	// ErrorLog — передайте ваш логгер.
	// ReadTimeout — таймаут для чтения. 5 секунд.
	// WriteTimeout — таймаут для записи. 10 секунд.
	// IdleTimeout — таймаут ожидания следующего запроса. 15 секунд.
	serv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	// 5 Верните ссылку на ваш сервер.
	return &Server{
		Logger: logger,
		Serv:   serv,
	}
}
