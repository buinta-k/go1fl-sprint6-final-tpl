package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HandleOne(w http.ResponseWriter, req *http.Request) {
	data, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, "Файл не найден", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset-utf-8")

	w.Write(data)
}

func HandleTwo(w http.ResponseWriter, req *http.Request) {
	file, header, err := req.FormFile("file")
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Ошибка чтения файла", http.StatusInternalServerError)
		return
	}
	inputString := string(fileBytes)

	resultString, err := service.Converter(inputString)
	if err != nil {
		http.Error(w, "Ошибка конвертации: "+err.Error(), http.StatusInternalServerError)
		return
	}

	ext := filepath.Ext(header.Filename)
	currentTime := time.Now().UTC().String()

	newFileName := currentTime + ext

	err = os.WriteFile(newFileName, []byte(resultString), 0644)
	if err != nil {
		http.Error(w, "Не удалось сохранить файл на сервер", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "text/plain; charset-utf-8")

	w.Write([]byte(resultString))
}
