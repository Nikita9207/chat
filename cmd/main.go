package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"
)

type Payload struct {
	UserName string `json:"user_name"`
	Message  string `json:"message"`
}

type ResponsePayload struct {
	Success bool `json:"success"`
}

var History = make([]Payload, 0, 0)

func main() {
	if _, err := os.Stat("History.txt"); err == nil {
		file, err := os.Open("History.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		decoder := json.NewDecoder(file)
		err = decoder.Decode(&History)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("История загружена!")
	} else if os.IsNotExist(err) {
		fmt.Println("Файл History.txt еще не создан!")
	} else {
		log.Fatal(err)
	}

	http.HandleFunc("/getlist", getlist)
	http.HandleFunc("/scripts/script.js", jFile)
	http.HandleFunc("/index", index)
	http.HandleFunc("/sendmessage", sendmessage)
	log.Println("Запущен сервер: http://127.0.0.1:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("/Users/nikita/chat/web/index.html")
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

func sendmessage(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var p Payload
	err := decoder.Decode(&p)
	if err != nil {
		panic(err)
	}
	History = append(History, p)
	log.Println(p.UserName)
	log.Println(p.Message)

	file, err := os.OpenFile("History.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(History)
	if err != nil {
		log.Fatal()
	}

	data := ResponsePayload{Success: true}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}

func getlist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(History)
}

func jFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	http.ServeFile(w, r, "/Users/nikita/chat/scripts/script.js")
}
