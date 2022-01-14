package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func handleUpload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(1024 << 10)
	if err != nil {
		log.Fatalf("Erro ao converter Form: %+v", err)
	}

	fileFromForm, fileHeader, err := r.FormFile("form_file")
	if err != nil {
		log.Fatalf("Erro ao criar arquivo vindo do form: %+v", err)
	}

	defer fileFromForm.Close()

	tempFile, err := os.CreateTemp("../temp", "file_*.png")

	if err != nil {
		log.Fatalf("Erro ao criar arquivo temporário: %+v", tempFile)
	}

	defer tempFile.Close()

	fileContent, err := io.ReadAll(fileFromForm)
	if err != nil {
		log.Fatalf("Erro ao ler dados do arquivo vindo do form: %+v", err)
	}

	_, err = tempFile.Write(fileContent)

	if err != nil {
		log.Fatalf("Erro ao escrever arquivo temporário: %+v", err)
	}

	fmt.Fprintln(w, "Uploado feito com sucesso!")
	fmt.Fprintln(w, "Nome do arquivo: ", fileHeader.Filename)
	fmt.Fprintln(w, "Tamanho do arquivo: ", fileHeader.Size, " bytes")
	fmt.Fprintln(w, "Header do form: ", fileHeader.Header)

}

func main() {
	http.HandleFunc("/upload", handleUpload)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Erro %+v", err)
	}
}
