package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"://github.com"
)

func HandleOne(w http.ResponseWriter, req *http.Request) {
	data, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, "Файл не найден", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(data)
}

func HandleTwo(w http.ResponseWriter, req *http.Request) {
	var inputString string
	var filename string

	file, header, err := req.FormFile("file")
	if err == nil {
		defer file.Close()
		fileBytes, readErr := io.ReadAll(file)
		if readErr == nil {
			inputString = string(fileBytes)
			filename = header.Filename
		}
	}

	if inputString == "" {
		bodyBytes, err := io.ReadAll(req.Body)
		if err == nil && len(bodyBytes) > 0 {
			inputString = string(bodyBytes)
		}
	}

	if inputString == "" {
		http.Error(w, "Данные запроса пусты", http.StatusBadRequest)
		return
	}

	resultString, err := service.AutoConvert(inputString)
	if err != nil {
		http.Error(w, "Ошибка конвертации: "+err.Error(), http.StatusInternalServerError)
		return
	}

	ext := filepath.Ext(filename)
	if ext == "" {
		ext = ".txt"
	}
	currentTime := time.Now().UTC().Format("20060102150405")
	newFileName := currentTime + ext

	err = os.WriteFile(newFileName, []byte(resultString), 0644)
	if err != nil {
		http.Error(w, "Не удалось найти файл", http.StatusInternalServerError)
		return
	}

	// 6. Ответ
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resultString))
}

