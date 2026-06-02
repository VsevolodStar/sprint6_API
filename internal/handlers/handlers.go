package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"sprint6_API/internal/service"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, "Unable to read index.html", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusInternalServerError)
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusInternalServerError)
		return
	}

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Unable to get file from form", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Unable to read file", http.StatusInternalServerError)
		return
	}

	convertedString, err := service.ConvertText(string(fileBytes))
	if err != nil {
		http.Error(w, "Unable to convert text", http.StatusInternalServerError)
	}

	fileName := time.Now().UTC().Format("2006-01-02 15-04-05") + filepath.Ext(handler.Filename)

	newFile, err := os.Create(fileName)
	if err != nil {
		http.Error(w, "Unable to create file", http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	_, err = newFile.WriteString(convertedString)
	if err != nil {
		http.Error(w, "Unable to write to file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, convertedString)

}
