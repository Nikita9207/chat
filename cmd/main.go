package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"
)

func main() {
	http.HandleFunc("/assets/script.js", jFile)
	http.HandleFunc("/index", index)
	log.Println("Запущен сервер: http://127.0.0.1:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("/Users/nikita/chat/assets/index.html")
	if err != nil {
		dir, err1 := os.Getwd()
		if err1 != nil {
			log.Fatal(err1)
		}
		fmt.Println(dir)
		io.WriteString(w, fmt.Sprintf("%v", err))
		return
	}
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, nil); err != nil {
		io.WriteString(w, fmt.Sprintf("%v", err))
		return
	}
	io.WriteString(w, tpl.String())
}

func jFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/Users/nikita/chat/assets/script.js")

}
