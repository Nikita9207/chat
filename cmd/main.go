package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"text/template"
	"time"
)

type Payload struct {
	Time     string `json:"time"`
	UserName string `json:"user_name"`
	Message  string `json:"message"`
}

type ResponsePayload struct {
	Success bool `json:"success"`
}

type User struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var History = make([]Payload, 0)
var Users = make([]User, 0)

func main() {
	if _, err := os.Stat("Registration.txt"); err == nil {
		file, err := os.Open("Registration.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		decoder := json.NewDecoder(file)
		err = decoder.Decode(&Users)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("История регистраций загружена")
	} else if os.IsNotExist(err) {
		fmt.Println("Файл Registration.txt еще не создан!")
	} else {
		log.Fatal(err)
	}

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

	fileServer := http.FileServer(http.Dir("/Users/nikita/chat/scripts/"))
	http.Handle("/scripts/", http.StripPrefix("/scripts", fileServer))

	http.HandleFunc("/", login)
	http.HandleFunc("/auth", auth)
	http.HandleFunc("/registration", registration)
	http.HandleFunc("/registrationData", registrationData)
	http.HandleFunc("/getList", getList)
	http.HandleFunc("/index", index)
	http.HandleFunc("/sendMessage", sendMessage)
	log.Println("Запущен сервер: http://127.0.0.1:80")
	log.Fatal(http.ListenAndServe(":80", nil))
}

func login(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("/Users/nikita/chat/web/login.html")
	if err != nil {
		dir, err1 := os.Getwd()
		if err1 != nil {
			log.Fatal(err1)
		}
		fmt.Println(dir)
		_, err = io.WriteString(w, fmt.Sprintf("%v", err))
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, nil); err != nil {
		_, err = io.WriteString(w, fmt.Sprintf("%v", err))
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	_, err = io.WriteString(w, tpl.String())
	if err != nil {
		log.Fatal(err)
	}
}

func registration(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("/Users/nikita/chat/web/registration.html")
	if err != nil {
		dir, err1 := os.Getwd()
		if err1 != nil {
			log.Fatal(err1)
		}
		fmt.Println(dir)
		_, err = io.WriteString(w, fmt.Sprintf("%v", err))
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, nil); err != nil {
		_, err = io.WriteString(w, fmt.Sprintf("%v", err))
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	_, err = io.WriteString(w, tpl.String())
	if err != nil {
		log.Fatal(err)
	}
}

/*
==============================================================
0. Создаем файл в main если его нет
0.1. Если файл есть, читаем и декодируем в код
1. Создаем глобальную переменную var Users = make([]User, 0) +
2. Создаем экземпляр структуры User struct +
3. Сохраняем в глобальную переменную +
4. Перезаписываем User через jsonEncode +
==============================================================
1.Во время авторизации добавить проверку наличия регистрации
==============================================================
1. Добавить шифрование пароля
2. Хэширование пароля (кодирование пароля)
==============================================================
*/

func registrationData(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
		return
	}

	name := r.Form.Get("name")
	surname := r.Form.Get("surname")
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	u := User{
		Name:     name,
		Surname:  surname,
		Username: username,
		Password: password,
	}

	Users = append(Users, u)
	log.Println(u)

	file, err := os.OpenFile("Registration.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(Users)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Данные успешно записаны в файл.")
}

func auth(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
		return
	}

	username := r.Form.Get("username")
	password := r.Form.Get("password")

	expiration := time.Now()
	expiration = expiration.AddDate(1, 0, 0)
	cookie := http.Cookie{Name: "username", Value: url.QueryEscape(username), Expires: expiration, Path: "/"}
	http.SetCookie(w, &cookie)
	fmt.Fprintf(w, "username: %s, password: %s", username, password)
	return
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("/Users/nikita/chat/web/index.html")
	if err != nil {
		dir, err1 := os.Getwd()
		if err1 != nil {
			log.Fatal(err1)
		}
		fmt.Println(dir)
		_, err = io.WriteString(w, fmt.Sprintf("%v", err))
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, nil); err != nil {
		_, err = io.WriteString(w, fmt.Sprintf("%v", err))
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	_, err = io.WriteString(w, tpl.String())
	if err != nil {
		log.Fatal(err)
	}
}

func sendMessage(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var p Payload
	err := decoder.Decode(&p)
	if err != nil {
		panic(err)
	}
	History = append(History, p)
	log.Println(p)

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
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Fatal(err)
	}
}

func getList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(History)
	if err != nil {
		log.Fatal(err)
	}
}
