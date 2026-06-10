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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(data)
}

func HandleTwo(w http.ResponseWriter, req *http.Request) {
	file, header, err := req.FormFile("file")
	if err != nil {
		http.Error(w, "Ошибка получения файла: "+err.Error(), http.StatusInternalServerError)
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
		http.Error(w, "Ошибка конвертации", http.StatusInternalServerError)
		return
	}


	// время через time.Now().UTC().String() + расширение через filepath.Ext()
	currentTime := time.Now().UTC().String()
	ext := filepath.Ext(header.Filename)
	newFileName := currentTime + ext

	err = os.WriteFile(newFileName, []byte(resultString), 0644)
	if err != nil {
		http.Error(w, "Не удалось сохранить файл на сервер", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resultString))
}
