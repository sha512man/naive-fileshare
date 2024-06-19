package main

import (
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

func main() {
	http.HandleFunc("/file/{name}", func(res http.ResponseWriter, req *http.Request) {
		token := req.URL.Query().Get("token")
		fileName := req.PathValue("name")
		filePath := path.Join("storage", fileName)

		f, err := os.Open(filePath)

		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		h := sha512.New()

		if _, err := io.Copy(h, f); err != nil {
			log.Fatal(err)
		}

		expectedHash := fmt.Sprintf("%x", h.Sum(nil))

		if token == expectedHash {
			http.ServeFile(res, req, filePath)
			return
		}

		res.WriteHeader(http.StatusUnauthorized)
		res.Header().Set("Content-Type", "application/json")

		data := make(map[string]string)
		data["message"] = "Unauthorized"

		jsonData, err := json.Marshal(data)

		if err != nil {
			log.Fatal(err)
		}

		res.Write(jsonData)

	})

	log.Println("Listening on :8070...")
	err := http.ListenAndServe(":8070", nil)

	if err != nil {
		log.Fatal(err)
	}
}
