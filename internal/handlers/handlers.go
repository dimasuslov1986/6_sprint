package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func MainHandle(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	res.WriteHeader(http.StatusOK)
	//res.Write(?)
}

func UploadHandle(res http.ResponseWriter, req *http.Request) {

	// Парсить html-форму из файла index.html.
	// req.ParseMultipartForm(10 << 20)
	if err := req.ParseMultipartForm(10 << 20); err != nil {
		http.Error(res, "Ошибка при разборе формы", http.StatusInternalServerError)
		return
	}

	// Получить файл из формы (не забудьте его закрыть).
	file, handler, err := req.FormFile("myFile")
	if err != nil {
		http.Error(res, "Ошибка при получении файла", http.StatusInternalServerError)
		return
	}
	// закрываем файл
	defer file.Close()

	// Прочитать данные из файла.
	body, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Ошибка чтения:", http.StatusInternalServerError)
		return
	}
	// Передать эти данные в функцию автоопределения из пакета service, которую вы создали, чтобы получить переконвертируемую строку.
	result, err := service.Convert(string(body))
	if err != nil {
		fmt.Println("Ошибка конвертации:", http.StatusInternalServerError)
		return
	}

	// Создать локальный файл. Эта операция обычно небезопасна и так делать не рекомендуется, но в рамках нашего задания хотелось бы более наглядного результата, поэтому мы решились на этот шаг, ради видимого результата. А вообще, обычно используют временные файлы.
	nameNewFile := time.Now().UTC().Format("2006-01-02 15:04:05")
	// получаем текущую директорию
	curDir, err := os.Getwd()
	if err != nil {
		log.Fatal(http.StatusInternalServerError)
	}
	ext := filepath.Ext(handler.Filename)
	newFile := filepath.Join(curDir, nameNewFile, ext)

	// Записать в локальный файл результат конвертации строки. Для генерации имени файла вы можете использовать время с помощью time.Now().UTC().String(). Чтобы получить расширения файла, используйте filepath.Ext().
	err = os.WriteFile(newFile, []byte(result), 0755)
	if err != nil {
		log.Fatal(http.StatusInternalServerError)
	}
	// Вернуть результат конвертации строки.
	res.Write([]byte(req.FormValue(result)))
}
