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

// Для корневого эндпоинта / нужно реализовать хендлер, который возвращает HTML из файла index.html.
func MainHandle(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	data, err := os.ReadFile("index.html")
	if err != nil {
		log.Fatal(http.StatusInternalServerError)
	}
	w.Write(data)
}

func UploadHandle(w http.ResponseWriter, r *http.Request) {

	// Парсить html-форму из файла index.html.
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Ошибка парсинга формы", http.StatusInternalServerError)
		return
	}

	// Получить файл из формы (не забудьте его закрыть).
	file, _, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Ошибка при получении файла", http.StatusInternalServerError)
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
	nameNewFile := time.Now().UTC().Format("2006_01_02_15-04-05")
	// получаем текущую директорию
	curDir, err := os.Getwd()
	if err != nil {
		log.Fatal(http.StatusInternalServerError)
		return
	}
	newFile, err := os.Create(filepath.Join(curDir, nameNewFile))
	if err != nil {
		fmt.Println("Ошибка создания файла:", http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	// Записать в локальный файл результат конвертации строки. Для генерации имени файла вы можете использовать время с помощью time.Now().UTC().String(). Чтобы получить расширения файла, используйте filepath.Ext().
	_, err = newFile.Write([]byte(result))
	if err != nil {
		log.Fatal(http.StatusInternalServerError)
		return
	}

	// Вернуть результат конвертации строки.
	w.Write([]byte(result))
}
